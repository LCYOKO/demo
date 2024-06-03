package net

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("status not ok, resp:%v", resp)
		return
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read respBody failed")
		return
	}
	fmt.Println(string(all))
}

func TestPost(t *testing.T) {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func TestHttpClient(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig:    &tls.Config{RootCAs: pool},
	//	DisableCompression: true,
	//}
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://example.com")
}

type myHandler struct{}

func (m *myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusAlreadyReported)
}

func testHttpServer() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
