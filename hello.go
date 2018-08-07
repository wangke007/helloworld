package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func main() {
	user := make(map[string]interface{})
	user["name"] = "jiaojiao"
	user["email"] = "123@qq.com"

	bytesData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error)
	}
	reader := bytes.NewReader(bytesData)

	url := "http://localhost:43302/users"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error)
	}

	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error)
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
