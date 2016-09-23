package service

import (
    "os"
    //"path/filepath"
    "log"
    "fmt"
)

func deletefile(path string, f os.FileInfo, err error) error {
    fmt.Println(path)
    if path == "./upload" {
        return nil
    }
    err1 := os.Remove(path)
    if err1 != nil {
        log.Fatal(err1)
        return err1
    }
    return err1
}

func DeleteCache(dirName string) {
    //err := filepath.Walk(dirName, deletefile)
    //checkNil(err)
    err := os.RemoveAll(dirName)
    checkNil(err)
    err = os.MkdirAll(dirName, 0700)
    checkNil(err)
}