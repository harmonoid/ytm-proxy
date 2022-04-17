# HarmonoidProxy
Simple proxy for Harmonoid written in Go - proxies requests for YouTube Music and YouTube.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/harmonoid/HarmonoidProxy)

# How does it work?
It's quite simple, actually. It forwards request body from your request (only for POST) and all headers to https://music.youtube.com.

Our proxy server also decodes gzipped response, so that you don't have to worry about anything.

# Example request
```http request
POST /music/youtubei/v1/browse?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Type: application/json
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36

{"context": {"client": {"clientName": "WEB", "clientVersion": "2.20210224.06.00", "newVisitorCookie": true}, "user": {"lockedSafetyMode": false}}, "params": "EgVhYm91dA%3D%3D", "browseId": "UC_aEa8K-EOJ3D6gOs7HcyNg"}
```

This request gets forwarded to following request:
```http request
POST https://music.youtube.com/youtubei/v1/browse?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Type: application/json
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36

{"context": {"client": {"clientName": "WEB", "clientVersion": "2.20210224.06.00", "newVisitorCookie": true}, "user": {"lockedSafetyMode": false}}, "params": "EgVhYm91dA%3D%3D", "browseId": "UC_aEa8K-EOJ3D6gOs7HcyNg"}
```
