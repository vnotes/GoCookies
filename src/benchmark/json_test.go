package benchmark

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"testing"
)

type Activity struct {
	ActivityID  int    `json:"activity_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var activityData = &Activity{17528, "活动-Japan", "Japan"}
var rawActivityData = []byte(`{"activity_id": 17528, "title": "活动-Japan", "description": "Japan"}`)

func BenchmarkStdMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(activityData)
		if err != nil {
			b.Error(err)
		}
	}
}
func BenchmarkStdUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := &Activity{}
		err := json.Unmarshal(rawActivityData, &result)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJsonIteratorMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := jsoniter.Marshal(activityData)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJsonIteratorUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := &Activity{}
		err := jsoniter.Unmarshal(rawActivityData, &result)
		if err != nil {
			b.Error(err)
		}
	}
}
