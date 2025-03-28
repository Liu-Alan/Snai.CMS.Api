package app

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"strings"
	"time"
)

func OtpSecret() string {
	randomStr := randStr(32)
	return strings.ToUpper(base32.StdEncoding.EncodeToString(randomStr))
}

func randStr(strSize int) []byte {
	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes = make([]byte, strSize)
	_, _ = rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return bytes
}

// 忽略时间误差，前后30秒时间
func OtpVerifyCode(secret string, code int32) bool {
	// 当前Otp Code
	if getCode(secret, 0) == code {
		return true
	}
	/*
		// 前30秒Otp Code
		if getCode(secret, -30) == code {
			return true
		}

		// 后30秒Otp Code
		if getCode(secret, 30) == code {
			return true
		}
	*/
	return false
}

// 获取Otp Code
func getCode(secret string, offset int64) int32 {
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix() + offset
	code := int32(oneTimePassword(key, toBytes(epochSeconds/30)))
	return code
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

// 获取动态码二维码内容
func OtpQrcode(issuer, user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s", issuer, user, secret)
}
