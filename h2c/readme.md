http2已经合入标准库的http包，可以直接使用。如果需要使用无tls认证的http2，可以使用h2c。使用方法参照文章和文章里的评论 [Go http2 和 h2c](https://colobu.com/2018/09/06/Go-http2-%E5%92%8C-h2c/)

h2c协议的部分可以被http1理解，因而被http1的服务器所响应，即跳过tls。在http1的处理函数里，由h2c接管解析数据，从而实现http2协议。


**ps**：wireshark中抓去本地的http2, 需要在protocal的http2指定端口号。对于chrome的http2抓包, 要处理证书问题, 参考 [使用 Wireshark 调试 HTTP/2 流量 | JerryQu 的小站](https://imququ.com/post/http2-traffic-in-wireshark.html)