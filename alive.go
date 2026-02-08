package main

import "net/http"

func IsAlive(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	return resp.StatusCode < 400
}
