# 0x01 - Go To School

## What is REST?

Representational State Transfer - Software Architecture Style (https://en.wikipedia.org/wiki/Representational_state_transfer)
- Form of messaging style between client <--> server 
- Widely used today as a standard of exposing services to third-party consumer
- The four types of request: GET,POST,PUT,DELETE
    - GET -> Retrieve a resource (R of the CRUD)
    - POST -> Create a resource (C of the CRUD)
    - PUT -> Update a resource (U of the CRUD)
    - DELETE -> Delete a reosource (D of the CRUD)

Difference of the four?
    - GET variables exists in the URL as a form of key value variables or in-url variables
        - http://example.com/resource?key=value
        - http://example.com/resource/value
    - POST,PUT,DELETE variables exists in the request body. Typically in the REST form, JSON is used. However, your may define your way such as post request body or XML as variable representation in the request body. However, you may include inurl variable similar to that of what was shown in the GET request above for additional parameterized requests.

The term "request" and "response" will be used in this tutorial. Request refers to a client request of a service, not limited to a browser client. Response refers to a service response when a service request is received. The different types of requests stated will be used to show the examples of creating a GoLang REST web service.



## Go to the GoLand

### Variables
Variable types:
- string
- bool
- int / int64
- float / float64
- byte
- null (This is important because you will use it for error checking.)

You may create variables with shorthand or standard notation.
standard notation: 
```go
var x int 
x = 1
```
shorthand 
```go
x := 1
```

### Slice (Arrays)

### Structs

### Pointers

### For Loops

### Function

### go 'commands'

