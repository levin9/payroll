package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.POST("upload", func(context *gin.Context) {
		fmt.Println("开始上传图片······")
		file, header, err := context.Request.FormFile("file")
		if err != nil {
			fmt.Println("bad request")
			context.String(http.StatusBadRequest, "bad request")
			return
		}
		biztype := context.Request.PostForm.Get("biztype")
		if biztype == "" {
			biztype = "temp"
		}
		dir, _ := os.Getwd()
		//filePath := fmt.Sprintf(biztype, `\`, header.Filename)
		filePath := dir + `/files/` + biztype + `/` + time.Now().Format("2006/01/02/")
		//fmt.Println(filePath)
		os.MkdirAll(filePath, os.ModePerm)
		os.Mkdir(filePath, os.ModePerm)
		out, err_file := os.Create(filePath + header.Filename)
		if err_file != nil {
			fmt.Println("can't create")
			log.Fatal(err)
		}
		defer out.Close()
		_, err_writer := io.Copy(out, file)
		if err_writer != nil {
			fmt.Println("can't writer")
			context.String(http.StatusBadRequest, "fail")
			log.Fatal(err)
		}
		context.String(http.StatusOK, "success file", filePath)

	})

}
