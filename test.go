package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "www.budtit.com"
	if !strings.HasPrefix(url, "http") {
		url = url + "http://"
		fmt.Print(url)
	} else {
		fmt.Print(url)
	}
}
