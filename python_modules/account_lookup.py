#!/bin/python

print("tonto")

import requests
from requests.structures import CaseInsensitiveDict

print("tonto 2")

headers = CaseInsensitiveDict()
headers["Connection"] = "keep-alive"
headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"
headers["Upgrade-Insecure-Requests"] = "1"
headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
headers["Accept-Language"] = "en-US,en;q=0.9"
headers["Accept-Encoding"] = "gzip, deflate"

# Please solve twitter as it does give a 200 even if account exists -> more advanced web scraping
services = ["github.com", "instagram.com", "facebook.com", "twitter.com", "aminoapps.com", "google.com"]

def make_request(url) -> bool:
    try:
        response = requests.get(url, headers=headers)
        
        print(f"\n\n\n\n\n\n\n\n\nTEXT RESPONSE -> << {response.text} >>\n\n\n\n\n\n\n\n\n")

        if response.status_code >= 200 and response.status_code < 300:
            return True
        else:
            return False
    except:
        return None

def get_results(username) -> list: 
    url_list = []
    for service in services:
        url = "https://" + service + "/" + username
        url_list.append(url)
    for url in url_list:
        exists = make_request(url)
        if exists:
            print(f"[*] ~ Found match -> {url}")

def interface():
    while True:
        username = input("Enter a username to search >> ")
        if username == "exit":
            exit(0)
        else:
            get_results(username)

if __name__ == "__main__":
    try:
        interface()
    except Exception as error:
        print(f"Core error happened -> {error}")
        exit(0)

