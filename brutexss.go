package main

import (
	"fmt"
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

 brutexss.go get http://www.baidu.com/?index=1
 brutexss.go post http://www.baidu.com/?index=1

`)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 2 {
		usage()
		return
	}

}
