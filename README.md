```md
# CurlHub 🚀  
**Public Web Recon & Learning Toolkit**
CurlHub is a learning-focused web reconnaissance tool designed to extract **publicly accessible** URLs, parameters, paths, JavaScript endpoints, and historical URLs using Wayback/Common Crawl data.
It is built for **education, practice, and experience** in web reconnaissance and bug bounty methodology.

> ⚠️ CurlHub does NOT bypass authentication, access private/local files, or perform any illegal activity.


## ✨ Features

- ✅ Subdomain list input (`-l domains.txt`)
- 🔗 Extract URLs & direct links
- 🔑 Extract query parameters
- 📌 Discover paths & API endpoints
- 📜 JavaScript endpoint extraction (threaded)
- 🕰️ Wayback + Common Crawl URL discovery (via `gau`)
- ❤️ Alive endpoint checker (HTTP status based)
- 📦 JSON output for easy parsing
- 🐚 Bash version (lightweight, Termux-friendly)
- 🐹 Go version (fast, extendable)


## 📁 Project Structure

----

CurlHub/
├── curlhub.sh
├── domains.txt
├── output/
│   ├── urls.txt
│   ├── parameters.txt
│   ├── paths.txt
│   ├── js_endpoints.txt
│   ├── alive.txt
│   └── results.json
└── README.md

````

## ⚙️ Requirements

### Mandatory
- `git`
- `curl`
- `grep`
- `gau`
- `jq`

### Optional (Go version)
- `golang`

## 🔧 Installation

### Kali / Ubuntu
```bash
sudo apt update
sudo apt install git curl grep jq -y
````

### Termux (Android)

```bash
pkg update && pkg upgrade -y
pkg install git curl grep jq golang -y
```

## 📦 Install gau (Wayback Support)

```bash
go install github.com/lc/gau/v2/cmd/gau@latest
export PATH=$PATH:$HOME/go/bin
```

Verify:

```bash
gau --version
```

## 📝 Prepare Domain List

Create a file named `domains.txt`:

```
example.com
api.example.com
test.example.com
```

Only include **in-scope domains**.

## ▶ Usage (Bash Version)

```bash
chmod +x curlhub.sh
./curlhub.sh -l domains.txt
```

## 📤 Output Files Explained

| File               | Description                         |
| ------------------ | ----------------------------------- |
| `urls.txt`         | All collected URLs (live + Wayback) |
| `parameters.txt`   | Extracted query parameters          |
| `paths.txt`        | Discovered paths & endpoints        |
| `js_endpoints.txt` | API endpoints found in JS files     |
| `alive.txt`        | Alive endpoints (HTTP 200/302)      |
| `results.json`     | Combined JSON output                |


## 📦 JSON Output Example

```json
{
  "urls": ["https://example.com/api"],
  "parameters": ["id", "user"],
  "paths": ["/api/login"],
  "js_endpoints": ["/api/data"],
  "alive": ["/api/login [200]"]
}
```

## 🐹 Go Version (Optional)

### Build & Run

```bash
go mod init curlhub
go build -o curlhub
./curlhub -l domains.txt
```

## 🧠 How CurlHub Works

1. Fetches live HTML using `curl`
2. Extracts URLs, parameters, and paths
3. Collects historical URLs via Wayback & Common Crawl (`gau`)
4. Parses JavaScript files using threading
5. Checks which endpoints are alive
6. Saves everything in structured output files

## ⚠️ Legal & Ethical Notice

* ✅ Use only on domains you own or have permission to test
* ❌ Do NOT use on unauthorized targets
* ❌ No authentication bypass
* ❌ No private or local file access

CurlHub is intended **strictly for education and ethical security research**.

## 🚀 Future Enhancements

* Subdomain enumeration integration
* HTTP status grouping
* Screenshot alive endpoints
* SQLite / database storage
* Full Go rewrite (no shell dependencies)

## 🧑‍💻 Author

Built for learning, recon practice, and security education.
Ideal for students, beginners, and bug bounty learners.

## ⭐ Support

If you find CurlHub useful:

* ⭐ Star the repository
* 🛠️ Fork and improve
* 📢 Share with learners

Happy recon! 🔍🚀
Just say the word, **mrrobots487** 🤖💙
```
