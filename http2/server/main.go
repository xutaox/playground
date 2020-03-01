package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr: ":8972",
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("this is a http2 test sever")
			w.Write([]byte("this is a http2 test sever"))
		}), h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}

/*
	如果访问的是 http://127.0.0.1:8972，client会有如下报错
	2020/03/01 23:32:07 error making request: Get http://localhost:8972: read tcp [::1]:51376->[::1]:8972: read: connection reset by peer
	client tcp发送了一串二进制, 然后server就是400，然后就会 TCP RST, 改成 http://127.0.0.1:8972/ 即可 =。=

	如果直接用 http.ListenAndServe, client会有如下报错
	2020/03/01 23:32:07 error making request: Get http://localhost:8972: unexpected EOF

	据 https://www.jianshu.com/p/ff16b0308e7c 的说法，是http2的情况下没有TLS, 所以降级为1.1，然后应为client是2.0所以主动关闭了(有点牵强, client可以妥协用1.1吧)
	我抓包看到的，和文章描述的还不一样，文章里client发送了http2请求, 我看到的是client tcp发送了一串二进制, 然后server就是400，然后就会 TCP RST???
	但按照上文的做法, 自己实现 ListenAndServe 的逻辑确实可以正常响应

	抄官方文档肯定不会错啦
	https://pkg.go.dev/golang.org/x/net/http2/h2c?tab=doc#example-NewHandler

	另外client如果是1.1请求，2.0server也可以正常响应

	todo 以上只是实践结果, 需要翻阅代码求证.

	todo 关于400, 目测是本机的wireshark解析不到http2的请求, 需要更新下wireshark.

*/
