package service

import (
    "fmt"
    "net/http"
    "math/rand"
    "strconv"
    "os"
    "log"
    "io"
)

func index(w http.ResponseWriter, r *http.Request) {
    fp, err := os.Open("./template/upload.html")
    if err != nil {
        log.Fatal("Open upload.html Fail!")
    }
    buf := make([]byte, 1024)
    fp.Read(buf)
    ret := string(buf)
    fmt.Fprint(w, ret)
    defer fp.Close()
}

func uploadfile(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseMultipartForm(32<<20)
        file, handle, err := r.FormFile("uploadfile")
        if err != nil {
            log.Fatal("Get UploadFile Error")
        }
        defer file.Close()
        f, err := os.OpenFile("./static/" + handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            log.Fatal("Save UploadFile Error")
        }
        defer f.Close()
        io.Copy(f, file)
        fmt.Fprint(w, "Upload Success.")
    }
}

func StartService() {
    var port int
    var url string
    http.HandleFunc("/", index)
    http.HandleFunc("/upload", uploadfile)
    for {
        port = rand.Intn(16000) + 49152
        url = "0.0.0.0:" + strconv.Itoa(port)
        fmt.Println(url)
        err := http.ListenAndServe(url, nil)
        if err == nil {
            break
        }
    }
}


