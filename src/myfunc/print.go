package myfunc

import (
	"bytes"
	"fmt"
	"os/exec"
)

func MyPrint(s string) {
	fmt.Println(s)
}

func ExecuteCommand() {
	c := exec.Command("ls", "-l")

	c.Stdout = &bytes.Buffer{}
	c.Stderr = &bytes.Buffer{}
	err := c.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println(c.Stderr.(*bytes.Buffer).String())
		return
	}
	fmt.Println(c.Stdout.(*bytes.Buffer).String())
}

func ExecuteCommand2() {
	cmd := exec.Command("echo 0")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

}
