package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") { // 判断url前缀
			url = "http://" + url
		}
		resp, err := http.Get(url) // 发起请求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) // 网页源代码
		c := resp.Status                    // 状态码
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf("%s", c)
	}
}
