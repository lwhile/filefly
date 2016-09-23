package test

import (
	"../service"
	"testing"
    "fmt"
)

func TestQrImg(t *testing.T) {
	url := "0.0.0.0:80"
	service.CreateQrImg(url)
}

func TestGetInterIp(t *testing.T) {
    fmt.Println(service.GetInterIp())
}

func TestDeleteFile(t *testing.T) {
    service.DeleteCache("../upload")
}