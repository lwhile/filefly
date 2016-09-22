package service

import (
    "fmt"
    //"net"
    "net/http"
    // "math/rand"
    // "strconv"
    "os"
    "log"
    "io"
    "github.com/toqueteos/webbrowser"
)

var PORT []string

func init() {
    PORT = []string{
        "65534",
    }
}
func index(w http.ResponseWriter, r *http.Request) {
    fp, err := os.Open("./template/index.html")
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
        r.ParseMultipartForm(32 << 20)
        file, handle, err := r.FormFile("uploadfile")
        if err != nil {
            log.Fatal("Get UploadFile Error")
        }
        defer file.Close()
        f, err := os.OpenFile("./static/" + handle.Filename, os.O_WRONLY | os.O_CREATE, 0666)
        if err != nil {
            log.Fatal("Save UploadFile Error")
        }
        defer f.Close()
        io.Copy(f, file)
        fmt.Fprint(w, "Upload Success.")
    }
}

func OpenBrowser(url string) {
    webbrowser.Open(url)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
    http.StripPrefix("/file", http.FileServer(http.Dir("./static/"))).ServeHTTP(w, r)
}

func download(w http.ResponseWriter, r *http.Request) {
    http.StripPrefix("/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
}
func StartService() {
    var url string = "0.0.0.0:" + PORT[0]
    http.HandleFunc("/index", index)
    http.HandleFunc("/upload", uploadfile)
    http.HandleFunc("/file", StaticServer)
    http.HandleFunc("/", download)
    fmt.Println("Starting service and opening browser.")
    go OpenBrowser(url + "/index")
    CreateQrImg(url)
    err1 := http.ListenAndServe(url, nil)

    if err1 != nil {
        fmt.Println("Start service error.")
        fmt.Println(err1)
    }
}


