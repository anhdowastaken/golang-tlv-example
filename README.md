# A Golang TLV example

https://levelup.gitconnected.com/binary-encoding-of-variable-length-options-with-golang-4481ff59e767

## Server

```text
$ go run main.go
2021/09/09 10:33:51 Listening tcp://127.0.0.1:8081...
```

## Client

```text
$ echo -e '\x00\x08\x00\x0a\x68\x65\x6c\x6c\x6f\x2c\x20\x67\x6f\x21' | nc -c 127.0.0.1 8081
```

Server's output:
```text
2021/09/09 10:33:54 type: 8
2021/09/09 10:33:54 payload: hello, go!
```
