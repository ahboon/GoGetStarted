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
Notice that your previous codes are not replying in JSON? 
Time to include structs!

Your data.go should look like this
```go
package main


type MyFirstStruct struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetString() *MyFirstStruct{
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


### Now lets give it inputs
Now lets try providing inputs. First, URL variables.
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
	} 
    output := GetString()
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```

Explaination:
```go
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
	} 
```
The if statement checks if the url variable "name" was given. So if the url was http://localhost:8000/helloworld, it would throw an error with the status code of Bad Request.
This is a way of error handling in Go. Because go does not have try catch, it uses nil if it does not exists.

Now lets read the variable.
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name == "" {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    output := GetString()
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```
Yet again, we are checking if the variable name is empty. Because http://localhost:8000/helloworld?name would pass the check for name, but we want to ensure that the value contains something. So we check if it is empty.


Now, now we need to pass it into the GetSting() function. Note that r.URL.Query()["name"] returns a slice of variables. So it is possible to have multiple "name" passed in and it reads as a slice. ``http://localhost:8000/helloworld?name=terry&name=bob -> Example of an array passed in``
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name[0] == "" {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    output := GetString(name[0])
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```
We also need to change data.go to handle the new variable.

Your data.go should look like this
```go
package main


type MyFirstStruct struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetString(name string) *MyFirstStruct{
    var tag *MyFirstStruct
    tag = &MyFirstStruct{}
    tag.Name = name
    tag.Age = 25
    return tag
}
```

Now lets pass in age as well.

Your logic.go
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name[0] == "" {
        		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    if r.URL.Query()["age"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    } 
    age,_ := r.URL.Query()["age"]
    if age[0] == "" {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    }
    output := GetString(name[0],age[0])
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}
```

Your data.go
```go
package main

import (
    "strconv"
)

type MyFirstStruct struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetString(name string,age string) *MyFirstStruct{
    intAge,_ := strconv.Atoi(age)
    var tag *MyFirstStruct
    tag = &MyFirstStruct{}
    tag.Name = name
    tag.Age = intAge
    return tag
}
```

Now lets try sending a JSON request!

Previously, we used URL variables and it can be done by doing a simple GET request. However, to send request bodies, we need to use POST request. (NOTE: It is still possible to read from the URL variables when sending a POST request. So you can sort of combine both of them. For the sake of this class, we will be using the previous code by adding on a POST request.)

We will be creating a new route for a POST request, and it will have a new function.

No change to main.go

Your logic.go should look like this:
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

type GoodBye struct {
    Name string `json:"name"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name[0] == "" {
        		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    if r.URL.Query()["age"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    } 
    age,_ := r.URL.Query()["age"]
    if age[0] == "" {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    }
    output := GetString(name[0],age[0])
    json.NewEncoder(w).Encode(output)
}

func ByeWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var inputVar GoodBye
    _ = json.NewDecoder(r.Body).Decode(&inputVar)
    output := ByeString(inputVar.Name)
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
    router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
    router.HandleFunc("/byeworld", ByeWorld).Methods("POST", "OPTIONS")
	return router
}
```

Your data.go should look like this
```go
package main

import (
    "strconv"
)

type MyFirstStruct struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetString(name string,age string) *MyFirstStruct{
    intAge,_ := strconv.Atoi(age)
    var tag *MyFirstStruct
    tag = &MyFirstStruct{}
    tag.Name = name
    tag.Age = intAge
    return tag
}

func ByeString(name string) *MyFirstStruct{
    var tag *MyFirstStruct
    tag = &MyFirstStruct{}
    tag.Name = name
    tag.Age = 999
    return tag
}
```

You can try some input validation by using nil as well:

logic.go
```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

type GoodBye struct {
    Name string `json:"name"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name[0] == "" {
        		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    if r.URL.Query()["age"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    } 
    age,_ := r.URL.Query()["age"]
    if age[0] == "" {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    }
    output := GetString(name[0],age[0])
    json.NewEncoder(w).Encode(output)
}

func ByeWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var inputVar GoodBye
    _ = json.NewDecoder(r.Body).Decode(&inputVar)
    if inputVar.Name == ""{
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "Service requires name in JSON"})
		return
    }
    output := ByeString(inputVar.Name)
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
    router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
    router.HandleFunc("/byeworld", ByeWorld).Methods("POST", "OPTIONS")
	return router
}
```