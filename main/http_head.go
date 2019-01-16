package main

import (
	"net/http"
	"fmt"
	"net"
	"time"
)

func main()  {
	url := []string{
		"http://www.baidu.com",
		"http://google.com",
		"http://taobao.com",
	}
	for _,v := range url{
		//设置请求超时
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err :%v\n", v, err)
			continue
		}
		fmt.Printf("head %s succ, status:%v\n", v, resp.Status)
	}
}
