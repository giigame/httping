package main

import (
  "fmt"
  "net"
  "net/http"
  "net/url"
  "os"
  "time"
)

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("Usage: httping url")
		os.Exit(1)
  }

	urlStr := os.Args[1]
	if urlStr[:7] != "http://" {
		urlStr = "http://" + urlStr
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Cannot resolve: " + urlStr)
		os.Exit(1)
		return
	}

	fmt.Printf("PING %s (%s):\n", url.Host, urlStr)
  ping(url, httpHead)
  //ping(url, tcpConnect)
}

func ping(url *url.URL, fn func(*url.URL) error) {
  for i := 0;; i++ {
    bnsec := time.Now().UnixNano()
		err := fn(url)
    nsec := time.Now().UnixNano()
    if err != nil {
			fmt.Println("Error:", err)
      return
    } else {
			time := nsec - bnsec
      fmt.Printf("connected to %s, seq=%d time=%.2f ms\n", url.Host, i, float32(time)/1e6)
    }
    time.Sleep(1e9)
  }
}

func httpHead(url *url.URL) error {
	_, err := http.Head(url.String())
	return err
}

func tcpConnect(url *url.URL) error {
	c, err := net.Dial("tcp", url.Host + ":80")

	// connection closed immediately
  c.Close()

  return err
}

