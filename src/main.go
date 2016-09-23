package main

import (
    "./service"
)

func main() {
    service.StartService()
    defer service.DeleteCache("../upload")
}
