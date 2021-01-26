package gohelpers

import (
	"fmt"
	"log"
	"bytes"
	"reflect"
	"net/http"
	"crypto/rand"
	"encoding/json"
	"encoding/base64"
	"encoding/binary"
)

type Response struct {
	Code	int			`json:"code"`
	Status	string		`json:"status"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

type JSONResponse struct {
	Status	string		`json:"status"`
	Message string		`json:"message"`
	Data	interface{}	`json:"data"`
}

func (response *Response) Success(code int, message string, data interface{}) {
	response.Code = code
	response.Status = "success"
	response.Message = message
	response.Data = data
}

func (response *Response) Error(code int, message string, data interface{}) {
	response.Code = code
	response.Status = "error"
	response.Message = message
	response.Data = data
}

func (response *JSONResponse) Success(message string, data interface{}) {
	response.Status = "success"
	response.Message = message
	response.Data = data
}

func (response *JSONResponse) Error(message string, data interface{}) {
	response.Status = "error"
	response.Message = message
	response.Data = data
}

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

func HandleResponse(response http.ResponseWriter, code int, message string, data interface{}) {
	var responseStruct = new(Response)

	if code == 200 || code == 201 || code == 202 {
		responseStruct.Success(code, message, data)
	} else {
		ErrorMessage(message, data)

		if data == nil {
			responseStruct.Error(code, message, nil)
		} else if fmt.Sprintf("%v", reflect.TypeOf(data).Kind()) == "ptr" {
			responseStruct.Error(code, message, fmt.Sprintf("%v", data))
		} else {
			responseStruct.Error(code, message, data)
		}
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write([]byte(JSONEncode(responseStruct)))
}

func HandleJSONResponse(status string, message string, data interface{}) string {
	var responseStruct = new(JSONResponse)

	if status == "success" {
		responseStruct.Success(message, data)
	} else {
		ErrorMessage(message, data)

		if data == nil {
			responseStruct.Error(message, nil)
		} else if fmt.Sprintf("%v", reflect.TypeOf(data).Kind()) == "ptr" {
			responseStruct.Error(message, fmt.Sprintf("%v", data))
		} else {
			responseStruct.Error(message, data)
		}
	}

	log.Println("Closing")
	log.Println()

	return JSONEncode(responseStruct)
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
