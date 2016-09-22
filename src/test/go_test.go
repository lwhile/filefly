package test

import (
    "testing"
    "../service"
)

func TestQrImg(t *testing.T) {
    url := "0.0.0.0:80"
    service.CreateQrImg(url)
}
