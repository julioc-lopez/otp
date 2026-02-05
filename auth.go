package otp

import (
	"crypto/hmac"
	"encoding/binary"
	"fmt"
	"hash"
	"log"
	"time"
)

// GetInterval returns the unix epoch divided by period and the number of seconds remaining till expiration.
func GetInterval(period int64) (int64, int64) {
	t := time.Now().Unix()
	iv := t / period
	remain := period - (t - (iv * period))
	return iv, remain
}

// GetCode returns a one-time password.
// The secret32 parameter is a Base32-encoded HMAC key.
// The iv parameter is the initialization value.
// The h parameter is a hash function to use in the HMAC.
// The digits parameter is the length of returned code.
//
// Example:
//
//	code, err := GetCode("MFRGGZDFMZTWQ2LK", 1, sha1.New, 6)
func GetCode(secret32 string, iv int64, h func() hash.Hash, digits int) (string, error) {
	key, err := decodeBase32(secret32)
	if err != nil {
		return "", err
	}

	var ivBuffer [8]byte

	n, err := binary.Encode(ivBuffer[:], binary.BigEndian, iv)

	assertNoError(err)
	assert(n == len(ivBuffer))

	mac := hmac.New(h, key)

	_, err = mac.Write(ivBuffer[:])
	assertNoError(err)

	digest := mac.Sum(nil)
	offset := digest[len(digest)-1] & 0xF

	assert(len(digest) >= int(offset+4))
	code := (binary.BigEndian.Uint32(digest[offset:offset+4]) & 0x7FFFFFFF)

	stringCode := fmt.Sprintf("%0*v", digits, code)
	if cl := len(stringCode); cl > digits {
		stringCode = stringCode[cl-digits:]
	}

	assert(len(stringCode) == digits)

	return stringCode, nil
}

func assert(condition bool, msg ...any) {
	if condition {
		return
	}

	if len(msg) > 0 {
		log.Fatalln(msg...)
	}

	log.Fatal("condition failed")
}

func assertNoError(err error) {
	assert(err == nil, "unexepcted error:", err)
}
