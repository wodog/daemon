# daemon

Run application as daemon

## Install

```
go get -u github.com/wodog/daemon/cmd/daemon
```

## Usage

#### Code
```go
package main

import (
	"log"
	"net/http"

	_ "github.com/wodog/daemon"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
```

```
go run main.go
```

#### CMD
```
daemon main.go
```
