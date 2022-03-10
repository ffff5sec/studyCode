package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "http://www.baidu.com"

	if strings.Contains(url, "www") {
		fmt.Println("cunzai")
	} else {
		fmt.Println("bucunzai")
	}
}
