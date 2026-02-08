#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: ./curlhub.sh https://example.com"
  exit 1
fi

TARGET=$1
DOMAIN=$(echo $TARGET | sed 's~https\?://~~' | sed 's~/.*~~')
OUT=output

mkdir -p $OUT
> $OUT/urls.txt
> $OUT/parameters.txt
> $OUT/paths.txt
> $OUT/js_endpoints.txt
> $OUT/wayback_urls.txt

echo "[+] Fetching live HTML..."
HTML=$(curl -Ls $TARGET)

echo "[+] Extracting live URLs..."
echo "$HTML" | grep -oP '(?<=href=")[^"]+' >> $OUT/urls.txt

echo "[+] Extracting parameters..."
echo "$HTML" | grep -oP '\?[a-zA-Z0-9_-]+=' | sed 's/[?=]//g' >> $OUT/parameters.txt

echo "[+] Extracting paths..."
echo "$HTML" | grep -oP '/(api|v1|v2|admin|user)[^" ]*' >> $OUT/paths.txt

echo "[+] Fetching Wayback + CommonCrawl URLs using gau..."
gau $DOMAIN >> $OUT/wayback_urls.txt

echo "[+] Extracting JS endpoints from Wayback URLs..."
grep '\.js' $OUT/wayback_urls.txt | sort -u | while read js; do
  curl -s "$js" | grep -oP '["'\''](/api/[^"'\'']+)["'\'']' \
  | tr -d '"' | tr -d "'" >> $OUT/js_endpoints.txt
done

echo "[+] Cleaning & sorting output..."
for f in $OUT/*.txt; do
  sort -u $f -o $f
done

echo "[✔] Done! Results saved in /output"
