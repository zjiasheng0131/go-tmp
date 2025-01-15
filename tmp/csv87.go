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
    // 存在CWE-78漏洞：直接执行用户输入的命令
    cmd := exec.Command("sh", "-c", userInput)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func main() {
    executeCommand()
}

