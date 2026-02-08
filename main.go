package main

import (
	"fmt"
	"os"
	"curlhub/crawler"
	"curlhub/extractor"
	"bufio"
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
		

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-l" {
		fmt.Println("Usage: ./curlhub -l domains.txt")
		return
	}

	file, _ := os.Open(os.Args[2])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		RunWayback(domain)
	}
}

