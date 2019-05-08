package main

import (
	"fmt"
	"github.com/DayDayYiDay/atreus-backend/common/auth"
	"github.com/DayDayYiDay/atreus-backend/common/metadata"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"os"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello
const uploadPath = "./tmp"

type TarFile struct {
	fileHash string
}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.POST("/upload").To(UploadFileHandler))
	ws.Route(ws.POST("/signin").To(auth.Signin))
	//ws.Route(ws.POST("/welcome").To(Welcome))
	//ws.Route(ws.POST("/refresh").To(UploadFileHandler))
	restful.Add(ws)

	//http.HandleFunc("/upload", config.UploadFileHandler())
	//
	//fs := http.FileServer(http.Dir(uploadPath))
	//http.Handle("/files/", http.StripPrefix("/files", fs))
	//
	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	// See func authHandler for an example auth handler that produces a token

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UploadFileHandler(req *restful.Request, resp *restful.Response) {

	result := make([]interface{}, 0)
	err := req.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Fatalln("parse error: ", err)
		return
	}
	file, handler, err := req.Request.FormFile("file")
	if err != nil {
		log.Fatalln(err)
		resp.WriteEntity(metadata.NewSuccessResp(result))
	}
	defer file.Close()
	dst, err := os.Create("/tmp/"+ handler.Filename)
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	//io.WriteString(resp, "world")
	err = resp.WriteEntity(metadata.NewSuccessResp(result))
	//err = resp.WriteEntity("success")
	if err != nil {
		log.Fatalln(err)
	}
}
