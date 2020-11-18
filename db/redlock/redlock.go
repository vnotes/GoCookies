package redlock

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	defaultRetryCount = 3
	unlockScript      = `
		if redis.call("get",KEYS[1]) == ARGV[1] then
			return redis.call("del",KEYS[1])
		else
			return 0
		end
`
	defaultRetryDelay = 5
	clockDriftFactor  = 0.01
)

type dlm struct {
	retryCount int
	retryDelay int
	quorum     int
	rdb        []*redis.Client
}

func New(p ...*redis.Client) *dlm {
	l := &dlm{
		retryCount: defaultRetryCount,
		retryDelay: defaultRetryDelay,
		rdb:        make([]*redis.Client, 0),
	}

	l.rdb = append(l.rdb, p...)
	l.quorum = (len(l.rdb) >> 1) + 1
	return l
}

func (l *dlm) SetRetryDelay(v int) {
	l.retryDelay = v
}

func (l *dlm) SetRetryCount(v int) {
	l.retryCount = v
}

func (l *dlm) Lock(resource string, ttl int64) *RDBLocker {
	var (
		val   = uniqueLockID()
		drift = float64(ttl)*clockDriftFactor + 2
	)
	for l.retryCount >= 0 {
		var (
			startTime = time.Now()
			cnt       int
		)
		for i := 0; i < len(l.rdb); i++ {
			succeed := lockInstance(l.rdb[i], resource, val, ttl)
			if succeed {
				cnt += 1
			}
		}
		validityTime := ttl - time.Since(startTime).Milliseconds() - int64(drift)
		if cnt >= l.quorum && validityTime > 0 {
			return &RDBLocker{
				dlm:      l,
				Resource: resource,
				Val:      val,
				Validity: validityTime,
			}
		}
		l.retryCount -= 1
		time.Sleep(time.Duration(l.retryDelay) * time.Millisecond)
	}
	for i := 0; i < len(l.rdb); i++ {
		_ = unlockInstance(l.rdb[i], resource, val)
	}
	return nil
}

func lockInstance(c *redis.Client, resource, val string, ttl int64) bool {
	v, e := c.SetNX(c.Context(), resource, val, time.Duration(ttl)*time.Millisecond).Result()
	if e != nil {
		return false
	}
	return v
}

func unlockInstance(c *redis.Client, resource, val string) error {
	_, err := c.Eval(c.Context(), unlockScript, []string{resource}, val).Result()
	if err != nil {
		return err
	}
	return nil
}

func uniqueLockID() string {
	b := make([]byte, 20)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

type RDBLocker struct {
	*dlm
	Resource string
	Val      string
	Validity int64
}

func (r *RDBLocker) Unlock() error {
	var err error
	for i := 0; i < len(r.rdb); i++ {
		err = unlockInstance(r.rdb[i], r.Resource, r.Val)
		if err != nil {
			return err
		}
	}
	return nil
}
