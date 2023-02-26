# Webassembly demo

Due to the fact that browsers do not support using `tcp` and `udp` directly, the client side should use `websocket` instead.

```go
js.Global().Get("WebSocket").New("ws://localhost:8080/echo")
```



## Quick Start

1. Build Webassembly

```
make build-wasm
```

2. Run TCP server

```
cd server
go run server
```

3. Host `index.html` with the HTTP server

```
python3 -m http.server 8000
```

4. Open `http://localhost:8080` in your browser.

5. Now you can see the client connect to the TCP server in terminal

```
Received message: GET /echo HTTP/1.1
Host: localhost:8080
Connection: Upgrade
Pragma: no-cache
Cache-Control: no-cache
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36
Upgrade: websocket
Origin: http://0.0.0.0:8000
Sec-WebSocket-Version: 13
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7
Sec-WebSocket-Key: KJPuTDiTFeddEzYXqus2Bg==
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
```