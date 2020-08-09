package cryptoutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

type RSARepository interface {
	GetRSAPrivateKey() (*rsa.PrivateKey, error)
	GetRSAPublicKey() (*rsa.PublicKey, error)
	ParseRSAPrivateKeyToPem() []byte
	ParseRSAPublicKeyToPem() []byte
}

type randomRSAKey struct {
	random io.Reader
	bits   int

	rsa *rsa.PrivateKey
}

func NewRandomRSAKey(random io.Reader, bits int) RSARepository {
	return &randomRSAKey{
		random: random,
		bits:   bits,
	}
}

func (r *randomRSAKey) GetRSAPrivateKey() (*rsa.PrivateKey, error) {
	if r.rsa != nil {
		return r.rsa, nil
	}
	privateKey, err := rsa.GenerateKey(r.random, r.bits)
	if err != nil {
		return nil, err
	}
	r.rsa = privateKey
	return privateKey, err
}

func (r *randomRSAKey) GetRSAPublicKey() (*rsa.PublicKey, error) {
	if r.rsa == nil {
		_, err := r.GetRSAPrivateKey()
		if err != nil {
			return nil, err
		}
	}
	return &r.rsa.PublicKey, nil
}

func (r *randomRSAKey) ParseRSAPrivateKeyToPem() []byte {
	privateByte := x509.MarshalPKCS1PrivateKey(r.rsa)
	privatePem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateByte,
	})
	return privatePem
}

func (r *randomRSAKey) ParseRSAPublicKeyToPem() []byte {
	publicByte := x509.MarshalPKCS1PublicKey(&r.rsa.PublicKey)
	publicPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicByte,
	})
	return publicPem
}

type pemRSAKey struct {
	privatePem string
	publicPem  string

	rsa *rsa.PrivateKey
}

func NewRSAFromPem(pem string) RSARepository {
	return &pemRSAKey{
		privatePem: pem,
	}
}

func (p *pemRSAKey) GetRSAPrivateKey() (*rsa.PrivateKey, error) {
	if p.rsa != nil {
		return p.rsa, nil
	}
	block, _ := pem.Decode([]byte(p.privatePem))
	if block == nil {
		return nil, fmt.Errorf("decode pem error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	p.rsa = privateKey
	return privateKey, err
}

func (p *pemRSAKey) GetRSAPublicKey() (*rsa.PublicKey, error) {
	if p.rsa == nil {
		_, err := p.GetRSAPrivateKey()
		if err != nil {
			return nil, err
		}
	}
	return &p.rsa.PublicKey, nil
}

func (p *pemRSAKey) ParseRSAPrivateKeyToPem() []byte {
	return []byte(p.privatePem)
}

func (p *pemRSAKey) ParseRSAPublicKeyToPem() []byte {
	if p.publicPem != "" {
		return []byte(p.publicPem)
	}
	publicByte := x509.MarshalPKCS1PublicKey(&p.rsa.PublicKey)
	publicPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicByte,
	})
	p.publicPem = string(publicPem)
	return publicPem
}

func ParseRSAPublicKeyFromPem(publicPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicPem))
	if block == nil {
		return nil, fmt.Errorf("decode pem error")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func VerifyMessage(msg []byte) error {
	shaBO := sha256.New()
	_, _ = shaBO.Write(msg)
	data := shaBO.Sum(nil)
	rsaBO := NewRandomRSAKey(rand.Reader, 2014)
	privateKey, _ := rsaBO.GetRSAPrivateKey()
	signature, _ := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, data, nil)
	publicKey, _ := rsaBO.GetRSAPublicKey()
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, data, signature, nil)
	return err
}
