---
title: "Test Post"
date: 2021-01-23T12:14:08Z
draft: false
---
# Introduction
This is going to be used for me to try out a few things

To run the server:

`hugo server -D --disableFastRender --bind 0.0.0.0 --baseURL http://192.168.xx.xx:1313`


Here is a little calculator test app I stole from [here](https://tutorialedge.net/golang/go-webassembly-tutorial/) [calculator](/calculator.html) 

To build this:


`
cd static

GOARCH=wasm GOOS=js go build -o calculator.wasm calculator/main.go
GOARCH=wasm GOOS=js go build -o puzzle.wasm puzzle/main.go
`

I am working on a room allocator tool - [Here](/room_allocator.html)

And - [Here](/puzzle.html) is a puzzle solvig website I'm developing. Uses JS for the UI, with GO WASM for the calculations.


Look! You can use colons to say where to align, including centring it.


|Title                   |Comment|
|:------:|------|
|Marvel      |Silly|
|DC| Dark, Broooding, silly!|
|Here|
|Bob|Tim|susan|
