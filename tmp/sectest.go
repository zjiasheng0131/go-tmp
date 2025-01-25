package main

import (
    "crypto/md5"
    "fmt"
    "io"
    "os"
)
const facebookAccessToken = "EAABwzLixnjYBAEXAMPLETOKEN1234567890" // 明文存储的 Facebook Access Token

const awsAccessKeyID = "AKIAEXAMPLEACCESSKEY" // 明文存储的 AWS Access Key ID
const awsSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" // 明文存储的 AWS Secret Access Key

func generateMD5Hash(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // 存在CWE-328漏洞：使用弱哈希函数MD5
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        fmt.Println("Error hashing file:", err)
        return
    }

    fmt.Printf("MD5 hash of %s: %x\n", filePath, hash.Sum(nil))
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <file_path>")
        return
    }
        fmt.Println("Using Facebook Access Token:", facebookAccessToken) 
    fmt.Println("Using AWS Access Key ID:", awsAccessKeyID) 
    generateMD5Hash(os.Args[1])
}
