package gohelpers

import (
	"io"
	"fmt"
	"time"
	"bytes"
	"errors"
	"runtime"
	"strings"
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"

	"github.com/MrAndreID/golog"
	"golang.org/x/crypto/bcrypt"
)

func ErrorMessage(message string, err interface{}) {
	golog.Error("Message : " + message + ".")

	if err != nil {
		golog.Error("Detail : " + fmt.Sprint(err))
	}
}

func JSONEncode(data interface{}) string {
	jsonResult, _ := json.Marshal(data)

	return string(jsonResult)
}

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)

	return b
}

func RandomByte(n int) string {
	return base64.URLEncoding.EncodeToString(Bytes(n))
}

func Random(randomType string, length int) string {
	var letterRunes []rune
	var bytesBuffer bytes.Buffer

	bytesBuffer.Grow(length)

	if randomType == "str" {
		letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	} else {
		letterRunes = []rune("0123456789")
	}

	letterLength := uint32(len(letterRunes))

	for i := 0; i < length; i++ {
		bytesBuffer.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%letterLength])
	}

	return bytesBuffer.String()
}

func GenerateKey(length int) string {
	return hex.EncodeToString(Bytes(length))
}

func Encrypt(key string, plainText string) (string, error) {
	bytesKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(bytesKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", aesGCM.Seal(nonce, nonce, []byte(plainText), nil)), nil
}

func Decrypt(key string, encryptedString string) (string, error) {
	bytesKey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(bytesKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func GetNewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for _, i := range maps {
		for j, k := range i {
			result[j] = k
		}
	}

	return result
}

func GenerateEncryptedKey(data []string, separator, key string) (string, error) {
	var plainText string

	for i, e := range data {
		if i == 0 {
			plainText = plainText + e
		} else {
			plainText = plainText + separator + e
		}
	}

	encryptedData, err := Encrypt(key, plainText)
	if err != nil {
		return "", err
	}

	return encryptedData, nil
}

func GenerateEncryptedKeyWithDatetime(data []string, separator, key string, date time.Time) (string, error) {
	var plainText string

	for i, e := range data {
		if i == 0 {
			plainText = plainText + e
		} else {
			plainText = plainText + separator + e
		}
	}

	plainText = plainText + separator + date.Format("2006-01-02 15:04:05")

	encryptedData, err := Encrypt(key, plainText)
	if err != nil {
		return "", err
	}

	return encryptedData, nil
}

func UngenerateEncryptedKey(data, separator, key string) ([]string, error) {
	plainText, err := Decrypt(key, data)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(plainText, separator), nil
}

func GenerateHashAndSalt(plainText string, customByte int, key string, cost int) (string, string, error) {
	if cost < 4 || cost > 31 {
		return "", "", errors.New("the cost must be between 4 and 31")
	}

	salt := RandomByte(customByte)

	encryptedSalt, _ := Encrypt(key, salt)

	hash, _ := bcrypt.GenerateFromPassword([]byte(plainText + salt), cost)

	encryptedHash, _ := Encrypt(key, string(hash))

	return encryptedHash, encryptedSalt, nil
}

func VerifyHashAndSalt(data, encryptedHash, encryptedSalt, key string) bool {
	hash, err := Decrypt(key, encryptedHash)
	if err != nil {
		return false
	}

	salt, err := Decrypt(key, encryptedSalt)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data + salt))
	if err != nil {
		return false
	}

	return true
}
