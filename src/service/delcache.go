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

    // Why is 0700, come here:http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
    err = os.MkdirAll(dirName, 0700)
    checkNil(err)
}