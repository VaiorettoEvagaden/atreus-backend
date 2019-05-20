package main

import (
	"fmt"
	"github.com/DayDayYiDay/atreus-backend/common/auth"
	"github.com/DayDayYiDay/atreus-backend/common/metadata"
	"github.com/DayDayYiDay/atreus-backend/common/workDir"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"os"
)

type TarFile struct {
	fileHash string
}

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.POST("/upload").To(UploadFileHandler))
	ws.Route(ws.POST("/signin").To(auth.Signin))
	restful.Add(ws)
	log.Print("Server started on localhost:9444, use /upload for uploading files ")
	log.Fatal(http.ListenAndServe(":9444", nil))
}

func UploadFileHandler(req *restful.Request, resp *restful.Response) {
	WorkPath := "/tmp/"
	args := os.Args[1:]
	if len(args) > 0 {
		for _, s := range args {
			WorkPath = s
		}
	}

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
	filename := WorkPath+ handler.Filename
	dst, err := os.Create(filename)
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		log.Fatalln(err)
	}

	// decompress
	var r io.Reader
	r, _ = os.Open(filename)
	err = workDir.Untar(WorkPath, r)
	if err != nil {
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

