package service

import (
    //"../github.com/boombuler/barcode"
    //"../github.com/boombuler/barcode/qr"
    "log"
    //"os"
    //"image/png"
    //"image/png"
    qrcode "github.com/skip2/go-qrcode"
)

const (
    ImageName = "qrimg.png"
)

func checkNil(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func CreateQrImg(url string) {
    url = "http://" + url
    err := qrcode.WriteFile(url, qrcode.Medium, 256, ImageName)
    checkNil(err)
}