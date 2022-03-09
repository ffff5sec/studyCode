package main

// echo 输出其命令行参数
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		fmt.Printf("%d-->%s\n", i, os.Args[i])
		sep = " "
	}
	fmt.Println(s)
}
