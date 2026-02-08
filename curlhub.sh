#!/bin/bash

while getopts "l:" opt; do
  case $opt in
    l) LIST=$OPTARG ;;
    *) echo "Usage: $0 -l domains.txt"; exit 1 ;;
  esac
done

if [ -z "$LIST" ]; then
  echo "[-] Domain list required"
  exit 1
fi

OUT=output
mkdir -p $OUT
> $OUT/urls.txt
> $OUT/parameters.txt
> $OUT/paths.txt
> $OUT/js_endpoints.txt
> $OUT/alive.txt

echo "[+] Running CurlHub v2"

cat $LIST | while read domain; do
  echo "[*] Processing $domain"

  # Live URLs
  curl -Ls https://$domain |
    grep -oP '(?<=href=")[^"]+' >> $OUT/urls.txt

  # Wayback + Common Crawl
  gau $domain >> $OUT/urls.txt
done

# Parameters
grep -oP '\?[a-zA-Z0-9_-]+=' $OUT/urls.txt |
  sed 's/[?=]//g' >> $OUT/parameters.txt

# Paths
grep -oP '/(api|v1|v2|admin|user)[^" ]*' $OUT/urls.txt >> $OUT/paths.txt

# Threaded JS parsing
echo "[+] Parsing JS files (threaded)"
grep '\.js' $OUT/urls.txt | sort -u |
  xargs -n 1 -P 10 curl -s |
  grep -oP '["'\''](/api/[^"'\'']+)["'\'']' |
  tr -d '"' | tr -d "'" >> $OUT/js_endpoints.txt

# Alive check
echo "[+] Checking alive endpoints"
cat $OUT/paths.txt $OUT/js_endpoints.txt | sort -u |
while read url; do
  code=$(curl -o /dev/null -s -w "%{http_code}" "$url")
  if [[ "$code" == "200" || "$code" == "302" ]]; then
    echo "$url [$code]" >> $OUT/alive.txt
  fi
done

# JSON output
echo "[+] Generating JSON output"
jq -n \
  --argfile urls <(jq -R . $OUT/urls.txt | jq -s .) \
  --argfile params <(jq -R . $OUT/parameters.txt | jq -s .) \
  --argfile paths <(jq -R . $OUT/paths.txt | jq -s .) \
  --argfile js <(jq -R . $OUT/js_endpoints.txt | jq -s .) \
  --argfile alive <(jq -R . $OUT/alive.txt | jq -s .) \
  '{
    urls: $urls,
    parameters: $params,
    paths: $paths,
    js_endpoints: $js,
    alive: $alive
  }' > $OUT/results.json

for f in $OUT/*.txt; do sort -u $f -o $f; done

echo "[✔] CurlHub v2 finished"
