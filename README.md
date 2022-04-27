# ytm-proxy
Simple proxy server to retrieve music data anonymously from YouTube & YouTube Music. Written in Go, is blazingly simple & fast.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/harmonoid/ytm-proxy)

# How does it work?
It's quite simple, actually. It forwards request body from your request (only for POST) and all headers to https://music.youtube.com or https://www.youtube.com, depending on what you select in the URL path (`/music/*` is for YouTube Music & `/youtube/*` is for YouTube). This proxy server also decompresses gzipped response, so that you don't have to worry about decoding/decompressing anything on client-side.

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
