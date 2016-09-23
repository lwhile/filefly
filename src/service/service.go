package service

import (
	"fmt"
	//"net"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/toqueteos/webbrowser"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type MyHandler struct{}

type home struct {
	Title string
}

const (
	TemplateDir = "./template"
	UploadDir   = "./upload/"
)

func checkNil(err error, msg string) {
	if err != nil {
		log.Fatal(err, msg)
	}
}

// 实现 MyHandler的ServerHttp方法
func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if funciton, ok := mux[r.URL.String()]; ok {
		funciton(w, r)
		return
	}

	// 匹配不到路由
	if ok, _ := regexp.MatchString("/css", r.URL.String()); ok {
		http.StripPrefix("/css", http.FileServer(http.Dir("./template/css"))).ServeHTTP(w, r)
	} else if ok, _ := regexp.MatchString("/images", r.URL.String()); ok {
		http.StripPrefix("/images", http.FileServer(http.Dir("./template/images"))).ServeHTTP(w, r)
	} else {
		http.StripPrefix("/", http.FileServer(http.Dir("./upload"))).ServeHTTP(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fp, err := os.Open("./template/index.html")
	if err != nil {
		log.Fatal("Open index.html Fail!")
	}
	buf := make([]byte, 2048)
	fp.Read(buf)
	ret := string(buf)
	fmt.Fprint(w, ret)
	defer fp.Close()
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("uploadfile")
		checkNil(err, "*")

		fp, err := os.OpenFile(UploadDir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0660)
		checkNil(err, "&")

		_, err1 := io.Copy(fp, file)
		checkNil(err1, "&")

		// 返回二维码图片
		CreateQrImg("http://192.168.1.185:65534/file")
		fpCode, err := os.Open("./qrimg.png")
		if err != nil {
			log.Fatal("Open File Error")
			log.Fatal(err)
		}
		buf := make([]byte, 32<<20)
		fpCode.Read(buf)
		ret := string(buf)
		fmt.Fprint(w, ret)

		// fmt.Println("Success")
		defer fpCode.Close()
		defer file.Close()
		defer fp.Close()
	}
}

func staticServer(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/file", http.FileServer(http.Dir("./upload"))).ServeHTTP(w, r)
}

func StartService() {
	server := http.Server{
		Addr:        "0.0.0.0:65534",
		Handler:     &MyHandler{},
		ReadTimeout: 10 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = index
	mux["/upload"] = upload
	mux["/file"] = staticServer
	go webbrowser.Open(server.Addr)
	server.ListenAndServe()
}

//
//var PORT []string
//
//func init() {
//    PORT = []string{
//        "65534",
//    }
//}
//func index(w http.ResponseWriter, r *http.Request) {
//    fp, err := os.Open("./template/index.html")
//    if err != nil {
//        log.Fatal("Open upload.html Fail!")
//    }
//    buf := make([]byte, 1024)
//    fp.Read(buf)
//    ret := string(buf)
//    fmt.Fprint(w, ret)
//    defer fp.Close()
//}
//
//func uploadfile(w http.ResponseWriter, r *http.Request) {
//    if r.Method == "POST" {
//        r.ParseMultipartForm(32 << 20)
//        file, handle, err := r.FormFile("uploadfile")
//        if err != nil {
//            log.Fatal("Get UploadFile Error")
//        }
//        defer file.Close()
//        f, err := os.OpenFile("./static/" + handle.Filename, os.O_WRONLY | os.O_CREATE, 0666)
//        if err != nil {
//            log.Fatal("Save UploadFile Error")
//            log.Fatal(err)
//        }
//        defer f.Close()
//        io.Copy(f, file)
//        log.Println("url:", "0.0.0.0:65534/" + handle.Filename)
//        CreateQrImg("http:0.0.0.0:65534/" + handle.Filename)
//        fp, err := os.Open("./qrimg.png")
//        if err != nil {
//            log.Fatal("Open File Error")
//            log.Fatal(err)
//        }
//        buf := make([]byte, 32 << 20)
//        fp.Read(buf)
//        ret := string(buf)
//        fmt.Fprint(w, ret)
//        defer fp.Close()
//    }
//}
//
//func OpenBrowser(url string) {
//    webbrowser.Open(url)
//}
//
//func StaticServer(w http.ResponseWriter, r *http.Request) {
//    http.StripPrefix("/file", http.FileServer(http.Dir("./static/"))).ServeHTTP(w, r)
//}
//
//func download(w http.ResponseWriter, r *http.Request) {
//    http.StripPrefix("/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
//}
//
//func TemplateServer(w http.ResponseWriter, r *http.Request) {
//    http.StripPrefix("/css", http.FileServer(http.Dir("./template/css"))).ServeHTTP(w, r)
//
//}
//
//func StartService() {
//    var url string = "0.0.0.0:" + PORT[0]
//    http.HandleFunc("/index", index)
//    http.HandleFunc("/upload", uploadfile)
//    http.HandleFunc("/file", StaticServer)
//    http.HandleFunc("/", download)
//    http.HandleFunc("/css", TemplateServer)
//    http.HandleFunc("/image", TemplateServer)
//    fmt.Println("Starting service and opening browser.")
//    go OpenBrowser(url + "/index")
//    err1 := http.ListenAndServe(url, nil)
//
//    if err1 != nil {
//        fmt.Println("Start service error.")
//        fmt.Println(err1)
//    }
//}
