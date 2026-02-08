package main

import (
	"fmt"
	"os"
	"curlhub/crawler"
	"curlhub/extractor"
)

func banner() {
	fmt.Println(`
 ██████╗ ██╗   ██╗██████╗ ██╗     ██╗  ██╗██╗   ██╗██████╗
██╔════╝ ██║   ██║██╔══██╗██║     ██║  ██║██║   ██║██╔══██╗
██║  ███╗██║   ██║██████╔╝██║     ███████║██║   ██║██████╔╝
██║   ██║██║   ██║██╔══██╗██║     ██╔══██║██║   ██║██╔══██╗
╚██████╔╝╚██████╔╝██║  ██║███████╗██║  ██║╚██████╔╝██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝

        CurlHub (Go Edition)
`)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curlhub https://example.com")
		return
	}

	target := os.Args[1]
	banner()

	html := crawler.Fetch(target)
	extractor.ExtractAll(html)

	fmt.Println("[+] Recon completed. Check /output")
}
