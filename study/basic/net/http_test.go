package net

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGet1(t *testing.T) {
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

func TestGet2(t *testing.T) {
	apiUrl := "http://127.0.0.1:9090/get"
	// URL param
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("send request failed,%+v .\n", err)
		return
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read respbody failed,error:%+v", err)
		return
	}
	fmt.Printf("result:%s .\n", string(all))
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
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func TestHttpClient(t *testing.T) {
	//FIXME 注意线上一定不能使用默认的
	client := &http.Client{
		//RequestTimeout
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			//TLS最大握手超时时间
			TLSHandshakeTimeout: time.Second,
			//最大响应时间
			ResponseHeaderTimeout: time.Second,
			MaxIdleConns:          10,
		},
	}
	_ = client
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

func (h handler) getStatusCode2(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()
	//需要注意的点：
	//如果你没有读取Respose.Body的内容，那么默认的 http transport 会直接关闭连接
	//如果你读取了Body的内容，下次连接可以直接复用
	//在高并发的场景下，建议你使用长连接，可以调用 io.Copy(io.Discard, resp.Body) 读取Body的内容。
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.StatusCode, nil
}

type handler struct {
	client http.Client
	url    string
}
