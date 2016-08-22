package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	thread = 20
)

func usage() {
	fmt.Println(`
  ____             _        __  ______ ____
 | __ ) _ __ _   _| |_ ___  \ \/ / ___/ ___|
 |  _ \| '__| | | | __/ _ \  \  /\___ \___ \
 | |_) | |  | |_| | ||  __/  /  \ ___) |__) |
 |____/|_|   \__,_|\__\___| /_/\_\____/____/

 BruteXSS - Cross-Site Scripting BruteForcer for go

 Author: redsword - https://7jdg.com

 Note: Using incorrect payloads in the custom
 wordlist may give you false positives so its
 better to use the wordlist which is already
 provided for positive results.

 brutexss.go GET http://www.baidu.com/?index=1 wordlist.txt
 brutexss.go POST http://www.baidu.com/?index=1 wordlist.txt

`)

	fmt.Printf("Example: %s GET https://example.com/?index=1 wordlist.txt\r\n", os.Args[0])
	fmt.Printf("Example: %s POST https://example.com/?index=1 wordlist.txt\r\n", os.Args[0])
	//os.Exit(1)
	//fmt.Println(os.Args[1])
}

func readfile(dir string) ([]byte, error) {
	f, err := os.Open(dir)
	if err != nil {
		err.Error()
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) != 4 {
		usage()
		return
	}

}
