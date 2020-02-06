# 0x02 - Go Get Good

Time to create REST services!

## Introduction
It is important to remember that every API call comes with a REQUEST and a RESPONSE. It also also imporant to remember that there are REQUEST headers and RESPONSE headers. Though this tutorial is not about the basics of HTTP, it is still important to the know difference.

Typical HTTP Request structure:
- URL
- REQUEST TYPE 
- REQUEST HEADERS
- VARIABLES (In URL or Request Body)

Typical HTTP Response Structure:
- RESPONSE CODE
- RESPONSE HEADERS
- VARIABLES (In URL or Response Body)


Now lets get started with a sample code!
## Code Code Code

In your command line, run this command:
```
go get -u github.com/gorilla/mux
```

Reccomended way of having your API file structure:
main.go  --> A file to say what port to host your files on
logic.go --> A file to handle your presentation layer
data.go  --> A file to handle your data layer

*Note: Presentation Layer -> What to do when a certain logic is called
*Note: Data layer -> The layer which data is connected to.

main.go
```go
package main

import (
	"net/http"
)

func main() {

	http.ListenAndServe(":8000", ServeService())
}
```

logic.go
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    output := GetString()
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```

data.go
```go
package main

func GetString() string{
    return "Hello World"
}
```

### JSON reply
Time to include structs!

Your data.go should look like this
```go
package main


type MyFirstStruct struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetString() MyFirstStruct{
    var tag *MyFirstStruct
    tag = &MyFirstStruct{}
    tag.Name = "Michael Tan You Zhuang"
    tag.Age = 25
    return tag
}
```

We also have to tell the server to write the response as "application/json" by changing the content type header.
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    output := GetString()
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```