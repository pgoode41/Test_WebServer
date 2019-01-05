package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<!doctype html>
    <html lang="en">
    <head>
    	<meta charset="utf-8">
    	<meta name="viewport" content="width=device-width, initial-scale=1">
    	<title>Test WebServer</title>
    	<style>
    		body {
    			font-family: -apple-system, BlinkMacSystemFont, sans-serif;
    		}
    	</style>
    </head>
    <body>
    	<h3>Welcome to Test WebServer</h3>
    	<h1 style="background-color:rgb(0, 255, 255);">Test WebServer</h1>
    	<img class="right" src="https://657cea1304d5d92ee105-33ee89321dddef28209b83f19f06774f.ssl.cf1.rackcdn.com/gophercloud-edf0a107430a35b63fae80ea5d465fe648e194637e78f52455482f49c543769d.png">
    	</body>
    	<h1>Test WebServer</h1>
			</html>
    `)
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
