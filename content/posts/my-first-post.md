---
title: "Test Post"
date: 2021-01-23T12:14:08Z
draft: false
---
# Introduction
This is going to be used for me to try out a few things

To run the server:

```bash
hugo server -D --disableFastRender --bind 0.0.0.0 --baseURL http://192.168.xx.xx:1313
```


Here is a little calculator test app I stole from [here](https://tutorialedge.net/golang/go-webassembly-tutorial/) [calculator](/calculator.html) 

To build this:


```bash
cd static
GOARCH=wasm GOOS=js go build -o calculator.wasm calculator/main.go
GOARCH=wasm GOOS=js go build -o puzzle.wasm puzzle/main.go
```

I am working on a room allocator tool - [Here](/room_allocator.html).Have you found yourself with n people and want them to meet everyone, but one meeting room is too many people. This could be the thing for you. 

And - [Here](/puzzle.html) is a puzzle solving website I'm developing. Uses JS for the UI, with GO WASM for the calculations.
