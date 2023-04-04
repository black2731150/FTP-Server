package main

import (
	"bytes"
	"flag"
	"fmt"
	"fpt/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	filename := flag.String("f", "", "上传文件的名字")
	Text := flag.String("m", " ", "上传文件的描述信息")
	Branch := flag.String("b", "", "上传文件所属分支")
	URL := flag.String("url", "http://121.5.145.152:1234", "上传文件的网址")
	flag.String("h", "Exampol: ./upload -url {http://127.0.0.1:8080} -f {filename} -m {文件备注} -b{文件所属分支}", "用法说明")
	flag.Parse()

	ok, err := common.PathExists(*filename)
	if !ok {
		fmt.Println("File Path error: ", err)
		os.Exit(-1)
	}

	_, err = postFile(*filename, *URL, *Text, *Branch)
	if err != nil {
		fmt.Println("Upload faild: ", err)
	} else {
		fmt.Println("Upload success!")
		common.HashMD5(*filename)
	}
}

func postFile(filename string, URL string, text string, branch string) (*http.Response, error) {

	url := URL + "/upload"

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	_, err := body_writer.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error: writing to buffer")
		return nil, err
	}

	//统一转换成绝对路径处理
	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("Get abs filepath error: ", err)
	}
	// fmt.Println("Filepath: ", absPath)

	fh, err := os.Open(absPath)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}

	boundary := body_writer.Boundary()
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filename)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, request_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "multipart/form-data;boundary="+boundary)

	query := req.URL.Query()
	query.Add("text", text)
	query.Add("branch", branch)
	req.URL.RawQuery = query.Encode()

	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)
}
