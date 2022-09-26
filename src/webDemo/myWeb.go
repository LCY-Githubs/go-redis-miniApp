package webDemo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Start() {
	fmt.Println("Starting..")
	http.HandleFunc("/hello", MyWebApp)
	http.HandleFunc("/app1", MyWebApp1)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

//writer 用于更新发送到浏览器的相应的值
func MyWebApp(writer http.ResponseWriter, request *http.Request) {
	message := NewSuccessRes("hello")
	b, _ := json.Marshal(message)
	_, err := writer.Write([]byte(string(b)))
	fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("/hello")
}

//writer 用于更新发送到浏览器的相应的值
func MyWebApp1(writer http.ResponseWriter, request *http.Request) {
	message := NewSuccessRes([]string{"1", "2", "3"})
	b, _ := json.Marshal(message)
	_, err := writer.Write([]byte(string(b)))
	fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("/app1")
}
