import subprocess
import re

def parse_js(js_url):
    try:
        js = subprocess.check_output(
            ["curl", "-s", js_url],
            universal_newlines=True
        )

        endpoints = set(re.findall(r'["\'](/api/[^"\']+)["\']', js))

        with open("output/js_endpoints.txt", "a") as f:
            for ep in endpoints:
                f.write(ep + "\n")

    except:
        pass
