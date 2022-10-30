package webDemo

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestIo(t *testing.T) {
	open, err := os.Open("test.txt")
	if err != nil {
		log.Print(err)
	}
	defer open.Close()
	//建立缓冲区
	buffer := make([]byte, 1024)
	//存放文件的所有内容
	var bytes []byte
	for {
		n, err1 := open.Read(buffer)
		if err1 == io.EOF {
			break
		}
		bytes = append(bytes, buffer[:n]...)
	}
	fmt.Println(string(bytes))
	//println(bytes)
}
