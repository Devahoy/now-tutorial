package main

import (
    "fmt"
    "net/http"
)

const Message = `
<link href="https://fonts.googleapis.com/css?family=Quicksand:700" rel="stylesheet">
<div style="display: flex;flex-direction: column; justify-content: center; align-items: center; height: 100vh; font-family: 'Quicksand', san-serif;">
    <h1>Dockerfile with <a href="https://now.sh">Now</a></h1>
    <p>See more : <a href="https://devahoy.com/posts/deploy-website-with-now/">มา Deploy Website แบบไม่เสียตังด้วย Now กันเถอะ</a></p>
</div>
`

func handler(w http.ResponseWriter, r *http.Request) {
	  w.Header().Set("Content-Type", "text/html;charset=UTF-8")
    fmt.Fprintf(w, Message)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}