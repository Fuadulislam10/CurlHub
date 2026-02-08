import re
from core.js_parser import parse_js

def extract_all(html, base):
    urls = set(re.findall(r'href=["\'](.*?)["\']', html))
    scripts = set(re.findall(r'src=["\'](.*?\.js)["\']', html))
    params = set(re.findall(r'\?(.*?)=', html))
    paths = set(re.findall(r'/(api|v1|v2|admin|user)[^"\']*', html))

    with open("output/urls.txt", "w") as f:
        f.write("\n".join(urls))

    with open("output/parameters.txt", "w") as f:
        f.write("\n".join(params))

    with open("output/paths.txt", "w") as f:
        f.write("\n".join(paths))

    for js in scripts:
        parse_js(js)
