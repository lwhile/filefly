package service

import (
    "os"
    "path/filepath"
    "log"
)

func deletefile(path string, f os.FileInfo, err error) error {
    err1 := os.Remove(path)
    if err != nil {
        log.Fatal("清除缓存失败", err1)
        return err1
    }
    return nil
}

func DeleteCache(dirName string) {
    err := filepath.Walk(dirName, deletefile)
    checkNil(err)
}