#!/usr/bin/env python3

import sys
from core.crawler import crawl
from core.extractor import extract_all

def banner():
    print("""
   ██████╗ ██╗   ██╗██████╗ ██╗     ██╗  ██╗██╗   ██╗██████╗ 
  ██╔════╝ ██║   ██║██╔══██╗██║     ██║  ██║██║   ██║██╔══██╗
  ██║  ███╗██║   ██║██████╔╝██║     ███████║██║   ██║██████╔╝
  ██║   ██║██║   ██║██╔══██╗██║     ██╔══██║██║   ██║██╔══██╗
  ╚██████╔╝╚██████╔╝██║  ██║███████╗██║  ██║╚██████╔╝██████╔╝
   ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝

        CurlHub - Public Recon Learning Tool
    """)

def main():
    if len(sys.argv) != 2:
        print("Usage: python3 curlhub.py https://example.com")
        sys.exit(1)

    target = sys.argv[1]
    banner()

    html = crawl(target)
    extract_all(html, target)

    print("\n[+] Done! Results saved in /output folder")

if __name__ == "__main__":
    main()
