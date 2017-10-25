package httpsvr

import "fmt"
import "net/http"
import "time"
import "log"
import "bytes"
import "io/ioutil"

type HttpServer struct {
	count int //http访问次数
}

//处理http访问
func (h *HttpServer) ServeHTTP(w http.ResponseWriter , r *http.Request){
	if r != nil {
		h.onRequest(r)
	}

	if nil != r {
		h.onResponse(w)
	}
}

//启动httpServer
func ( *HttpServer )StartHttpServer() {
	s := &http.Server{
		Addr: ":8080",
		Handler: &HttpServer{},
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

//处理请求数据
func (*HttpServer)onRequest(r *http.Request) {
	fmt.Println("-- on request --")
	fmt.Println("Method: " + r.Method)
	fmt.Println("URL: " + r.URL.Path)
	fmt.Println("Host: " + r.Host)
	body,err := ioutil.ReadAll(r.Body)
	if nil == err {
		fmt.Println("Body: " + bytes.NewBuffer(body).String())
	}
}

//处理返回
func (h *HttpServer)onResponse(w http.ResponseWriter) {
	h.count ++
	s := "<html><body>";
	s += fmt.Sprintf("No.%d request the server ",h.count)
	s += "<br><h1>11111</h1>测试"
	s += "</body></html>";
	w.Write(bytes.NewBufferString(s).Bytes())
}