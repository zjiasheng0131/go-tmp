package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

func executeCommand() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command to execute: ")
    userInput, _ := reader.ReadString('\n')
    // 存在CWE-253漏洞：未检查命令执行是否成功
    cmd := exec.Command("sh", "-c", userInput)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run() // 未检查错误返回
}

func main() {
    executeCommand()

