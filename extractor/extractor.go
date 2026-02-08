package extractor

import (
	"os"
	"regexp"
	"curlhub/jsparser"
)

func save(file string, data map[string]bool) {
	f, _ := os.Create("output/" + file)
	defer f.Close()
	for k := range data {
		f.WriteString(k + "\n")
	}
}

func ExtractAll(html string) {
	urls := make(map[string]bool)
	params := make(map[string]bool)
	paths := make(map[string]bool)
	scripts := make(map[string]bool)

	reURL := regexp.MustCompile(`href=["'](.*?)["']`)
	reJS := regexp.MustCompile(`src=["'](.*?\.js)["']`)
	reParam := regexp.MustCompile(`\?([a-zA-Z0-9_-]+)=`)
	rePath := regexp.MustCompile(`/(api|v1|v2|admin|user)[^"' ]*`)

	for _, m := range reURL.FindAllStringSubmatch(html, -1) {
		urls[m[1]] = true
	}
	for _, m := range reJS.FindAllStringSubmatch(html, -1) {
		scripts[m[1]] = true
	}
	for _, m := range reParam.FindAllStringSubmatch(html, -1) {
		params[m[1]] = true
	}
	for _, m := range rePath.FindAllStringSubmatch(html, -1) {
		paths[m[0]] = true
	}

	save("urls.txt", urls)
	save("parameters.txt", params)
	save("paths.txt", paths)

	for js := range scripts {
		jsparser.Parse(js)
	}
}
