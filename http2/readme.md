参照 https://colobu.com/2018/09/06/Go-http2-%E5%92%8C-h2c/ , client会有如下报错

2020/03/01 23:32:07 error making request: Get http://localhost:8972: unexpected EOF

据 https://www.jianshu.com/p/ff16b0308e7c 的说法，是http2的情况下没有TLS, 所以降级为1.1，然后应为client是2.0所以主动关闭了(有点牵强, client可以妥协用1.1吧)

抓包看到的, client发送了http2请求, 然后server发现事态不对, 400然后TCP FIN。client不管，继续发送http2的get请求，server怕场面尴尬，TCP RST结束聊天

按照上文的做法, 自己实现 ListenAndServe 的逻辑确实可以正常响应

---

抄官方文档肯定不会错啦

https://pkg.go.dev/golang.org/x/net/http2/h2c?tab=doc#example-NewHandler

---

另外client如果是1.1请求，2.0server也可以正常响应

---

todo 以上只是实践结果, 需要翻阅代码求证.

---

ps

wireshark中抓去本地的http2, 需要在protocal的http2指定端口号

对于chrome的http2抓包, 要处理证书问题, 参考 https://imququ.com/post/http2-traffic-in-wireshark.html
