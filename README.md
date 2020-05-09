# force-ssl-heroku 

Golang middleware that redirects unencrypted HTTP requests to HTTPS on Heroku instances.

Heroku does SSL termination at its load balancer. However, the app can tell if the original request was made with HTTP by inspecting headers inserted by Heroku. We can use this to redirect to the HTTPS Heroku url.

## Installation

```sh
go get github.com/jonahgeorge/force-ssl-heroku
```

## Usage

```go
package main

import (
	"net/http"

        heroku "github.com/jonahgeorge/force-ssl-heroku"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", helloWorldHandler)

	http.ListenAndServe(":8080", heroku.ForceSsl(r))
}
```

## Caveat

It works because Heroku exposes your app through a reverse proxy which is used for load-balancing and other things.  This reverse proxy does SSL termination and forwards to your app which __should only accept connections from localhost__.  The middleware detects this situation by inspecting headers inserted by Heroku's reverse proxy;  __since headers can be spoofed, you should not use this middleware anywhere that's not behind such a proxy__!
