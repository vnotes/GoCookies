package cryptoutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomRSAKey(t *testing.T) {
	message := []byte("RSA DEMO...")
	rsaBO := NewRandomRSAKey(rand.Reader, 2048)
	privateKey, err := rsaBO.GetRSAPrivateKey()
	assert.Nil(t, err)
	publicKey, err := rsaBO.GetRSAPublicKey()
	assert.Nil(t, err)
	encryptionMsg, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	assert.Nil(t, err)
	decryptionMsg, err := privateKey.Decrypt(nil, encryptionMsg, &rsa.OAEPOptions{Hash: crypto.SHA256})
	assert.Nil(t, err)

	assert.Equal(t, message, decryptionMsg)
	t.Log(string(rsaBO.ParseRSAPrivateKeyToPem()))
	t.Log(string(rsaBO.ParseRSAPublicKeyToPem()))
}

func TestRSAKeyPem(t *testing.T) {
	message := []byte("RSA DEMO...")
	privatePem := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpwIBAAKCAQEAnge9i581KXsaP+N9KWiEzCsMvRotgmlGxVjgSV54+7dSXyTv\nzp1erxXRMZjxUky0VxqULIMJ+DySXTSndC3t9k5AtqzrxOK/eJFJqtcnM5fzR27R\nApQh3xyyOI3r4VdqHLBkUEQRtZMLgdf0ssHE7qqOWUvZbrFnAzSwrS1JjtceL4OT\nuOJnt+047slQqIBauVTWJB22JtixY9xgEWeyRPTl4M28C9xdgZ1zWqV24N8+krAE\n9cs44LCqJGOMF0WC2UXfeokfnBK6cXP/9qZQtk0rNTGT0jbkRoAc02y9rqum3Sc4\nPtKnsSTPjIobjkL4tyfgf0iVOSfFa0sFFILluwIDAQABAoIBAHGrV6eqaO7hYUw7\nMbgrOfoxrQQIZFu9pe4ls46SqatFfbZ6NxQDFiwSIVrvjAWKrt1IfPRXfUBVMWJh\nkWF7+AKG3v2G/D+uadMrjjWYdNYjohXrm+Oi8YoudfNAAj8gRW9FYaJqk+JrVoea\n8qPxxNl3TNTmPgvlBm0Fu5ThiVJZcrk9xSCGdWdGQgBIbCj69NY5hKg9XQsiRN+l\nc8aaqKttGUZNStLu8x5lqPRjTN5Sx1aD5W+A+DCSTba9+l+FbBvMu6RroqVz009P\n4AYopkBOYFgxwAeGwWNtadQQbYTP+klCNXJ73gkD5DsS/+3UXTimhEIEP1XkM6m6\nquf0dYECgYkA7V+yjVjqeqUg6HT3dafJi3LKr+zwKaH/+NUNbjrqr3tQMH4rLvp/\noU/SiwQVQY9uSV4KW+fNs5Tf7kWvCunnN80ovN9pDDkKf0G+o2OSG1G4FMYazuxS\n0BFYyfuGqI6xK8ioGcLX/F0WDkFMMZW1eZNj2IHiv1KAXIuWelh2ilBatYGrUlUq\ncQJ5AKpuNcSF8AapYY8VL5ulg6nXy5Y7bDMHYNw0GI9G5+CBlCxXXfMGuh6ci1+a\n3jl+DV0S9hrBPX1SF3QnhlIylpfn873+wUeIl6593Fo33L6A8jBbUtazxXrRwHjA\nbdpNJIzlSjRCRl5CzWpdp0+l8AL2SeDJKLvw6wKBiF9ZQmqenclYDSjy2vfqxv15\nxcr2/N1sUlrMkdGGXwDQIrzn4UbEnoHYg3UN1c/44k8cNEMIkMsi8PRQD2jt3c+/\nXC7J+vNK8ll9uir9cxIAOFY34UrfCMDFRwoTO9r3PlmdB1EZwBKB/bsKJaYBQd0t\nBD1Spxc894y1EWPhpvQWJOmlxYCT7zECdx/0iiBO/LJDEAfD8Sk933H5BDMm8SKg\nP4kKO5Chuthdc2rc9sCagks7DubeIsyk1dydsYdStDTLL5qXzpea5KGW3BzDp3nC\nIq6U6rv/vWP6yW5HsUCQdHaS3YPGkNJdpHzGgSNLSeZioCFRbg2BD39+rlt3XCvh\nAoGIT4R2BvsLu8qtT+2iD7qECdkLRqfmD1ZqG/sJtcr0IAXvIMFrZvojy8P/nsLo\nXlLfuGre+j9ID0DvZmwskZzb91PsQBHoSQxxQSoa2+bM8mJTHBZKGamxvmhOGfm5\nQ7EhYmZWRXe8roihKp5CP5lkUdoME9x0Z7SgGQ3ZMOAmHd7D99XNaFPKbg==\n-----END RSA PRIVATE KEY-----\n"
	rsaBO := NewRSAFromPem(privatePem)
	privateKey, err := rsaBO.GetRSAPrivateKey()
	assert.Nil(t, err)
	publicKey, err := rsaBO.GetRSAPublicKey()
	assert.Nil(t, err)
	encryptionMsg, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	assert.Nil(t, err)
	decryptionMsg, err := privateKey.Decrypt(nil, encryptionMsg, &rsa.OAEPOptions{Hash: crypto.SHA256})
	assert.Nil(t, err)

	assert.Equal(t, message, decryptionMsg)
	t.Log(string(rsaBO.ParseRSAPrivateKeyToPem()))
	t.Log(string(rsaBO.ParseRSAPublicKeyToPem()))
}

func TestParseRSAPublicKeyFromPem(t *testing.T) {
	pem1 := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpwIBAAKCAQEAnge9i581KXsaP+N9KWiEzCsMvRotgmlGxVjgSV54+7dSXyTv\nzp1erxXRMZjxUky0VxqULIMJ+DySXTSndC3t9k5AtqzrxOK/eJFJqtcnM5fzR27R\nApQh3xyyOI3r4VdqHLBkUEQRtZMLgdf0ssHE7qqOWUvZbrFnAzSwrS1JjtceL4OT\nuOJnt+047slQqIBauVTWJB22JtixY9xgEWeyRPTl4M28C9xdgZ1zWqV24N8+krAE\n9cs44LCqJGOMF0WC2UXfeokfnBK6cXP/9qZQtk0rNTGT0jbkRoAc02y9rqum3Sc4\nPtKnsSTPjIobjkL4tyfgf0iVOSfFa0sFFILluwIDAQABAoIBAHGrV6eqaO7hYUw7\nMbgrOfoxrQQIZFu9pe4ls46SqatFfbZ6NxQDFiwSIVrvjAWKrt1IfPRXfUBVMWJh\nkWF7+AKG3v2G/D+uadMrjjWYdNYjohXrm+Oi8YoudfNAAj8gRW9FYaJqk+JrVoea\n8qPxxNl3TNTmPgvlBm0Fu5ThiVJZcrk9xSCGdWdGQgBIbCj69NY5hKg9XQsiRN+l\nc8aaqKttGUZNStLu8x5lqPRjTN5Sx1aD5W+A+DCSTba9+l+FbBvMu6RroqVz009P\n4AYopkBOYFgxwAeGwWNtadQQbYTP+klCNXJ73gkD5DsS/+3UXTimhEIEP1XkM6m6\nquf0dYECgYkA7V+yjVjqeqUg6HT3dafJi3LKr+zwKaH/+NUNbjrqr3tQMH4rLvp/\noU/SiwQVQY9uSV4KW+fNs5Tf7kWvCunnN80ovN9pDDkKf0G+o2OSG1G4FMYazuxS\n0BFYyfuGqI6xK8ioGcLX/F0WDkFMMZW1eZNj2IHiv1KAXIuWelh2ilBatYGrUlUq\ncQJ5AKpuNcSF8AapYY8VL5ulg6nXy5Y7bDMHYNw0GI9G5+CBlCxXXfMGuh6ci1+a\n3jl+DV0S9hrBPX1SF3QnhlIylpfn873+wUeIl6593Fo33L6A8jBbUtazxXrRwHjA\nbdpNJIzlSjRCRl5CzWpdp0+l8AL2SeDJKLvw6wKBiF9ZQmqenclYDSjy2vfqxv15\nxcr2/N1sUlrMkdGGXwDQIrzn4UbEnoHYg3UN1c/44k8cNEMIkMsi8PRQD2jt3c+/\nXC7J+vNK8ll9uir9cxIAOFY34UrfCMDFRwoTO9r3PlmdB1EZwBKB/bsKJaYBQd0t\nBD1Spxc894y1EWPhpvQWJOmlxYCT7zECdx/0iiBO/LJDEAfD8Sk933H5BDMm8SKg\nP4kKO5Chuthdc2rc9sCagks7DubeIsyk1dydsYdStDTLL5qXzpea5KGW3BzDp3nC\nIq6U6rv/vWP6yW5HsUCQdHaS3YPGkNJdpHzGgSNLSeZioCFRbg2BD39+rlt3XCvh\nAoGIT4R2BvsLu8qtT+2iD7qECdkLRqfmD1ZqG/sJtcr0IAXvIMFrZvojy8P/nsLo\nXlLfuGre+j9ID0DvZmwskZzb91PsQBHoSQxxQSoa2+bM8mJTHBZKGamxvmhOGfm5\nQ7EhYmZWRXe8roihKp5CP5lkUdoME9x0Z7SgGQ3ZMOAmHd7D99XNaFPKbg==\n-----END RSA PRIVATE KEY-----\n"
	_, err1 := ParseRSAPublicKeyFromPem(pem1)
	assert.NotNil(t, err1)

	pem2 := "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAnge9i581KXsaP+N9KWiEzCsMvRotgmlGxVjgSV54+7dSXyTvzp1e\nrxXRMZjxUky0VxqULIMJ+DySXTSndC3t9k5AtqzrxOK/eJFJqtcnM5fzR27RApQh\n3xyyOI3r4VdqHLBkUEQRtZMLgdf0ssHE7qqOWUvZbrFnAzSwrS1JjtceL4OTuOJn\nt+047slQqIBauVTWJB22JtixY9xgEWeyRPTl4M28C9xdgZ1zWqV24N8+krAE9cs4\n4LCqJGOMF0WC2UXfeokfnBK6cXP/9qZQtk0rNTGT0jbkRoAc02y9rqum3Sc4PtKn\nsSTPjIobjkL4tyfgf0iVOSfFa0sFFILluwIDAQAB\n-----END RSA PUBLIC KEY-----\n"
	publicKey2, err2 := ParseRSAPublicKeyFromPem(pem2)
	assert.Nil(t, err2)
	t.Log(publicKey2)

}

func TestVerifyMessage(t *testing.T) {
	var msg = []byte("RSA VERIFY...")
	err := VerifyMessage(msg)
	assert.Nil(t, err)
}
