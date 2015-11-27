package main

// To build:
// 1) Set $GOOS and $GOARCH to the correct values
// 2) go build

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func main() {
	fmt.Println("This was cross-compiled on linux for the following OS and Architecture:")
	fmt.Printf("OS: %s\nArchitecture: %s\n\n", runtime.GOOS, runtime.GOARCH)

	url := "https://example.com"
	fmt.Printf("Lets contact %s...\n\n", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		fmt.Println("RESPONSE BODY:")
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}
