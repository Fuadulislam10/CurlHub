import subprocess

def crawl(url):
    try:
        result = subprocess.check_output(
            ["curl", "-L", "-s", url],
            universal_newlines=True
        )
        return result
    except Exception as e:
        print("Error fetching target:", e)
        return ""
