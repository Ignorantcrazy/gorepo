package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	h bool
	d string
	p int
)

func init() {
	flag.BoolVar(&h, "h", false, "帮助")

	flag.StringVar(&d, "d", "", "设置工作目录")
	flag.IntVar(&p, "p", 8000, "设置端口，默认为8000")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}

	dir := `./`
	if d != "" {
		if _, err := os.Stat(d); err == nil {
			dir = d
		}
	}

	port := 8000
	if p > 0 {
		port = p
	}

	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port %v", port)
	listenport := fmt.Sprintf("%s%d", ":", port)
	fmt.Println(listenport)
	if err := http.ListenAndServe(listenport, nil); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `miniweb version: v1.0
Usage: miniweb [-hdp] [-h 帮住] [-d 目录] [-p 端口]

Options:
`)
	flag.PrintDefaults()
}
