package gohelpers

import (
	"fmt"
	"log"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"encoding/base64"
	"encoding/binary"
)

func ErrorMessage(message string, err interface{}) {
	fmt.Println()
	log.Println("-------------------- Start Error Message --------------------")
	log.Println("Message => " + message + ".")

	if err != nil {
		log.Println("Error => ", err)
	}

	log.Println("-------------------- End Of Error Message --------------------")
	fmt.Println()
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
