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
y := 2
```
Full Example: https://play.golang.org/p/_JfDN-J9qrd

### Slice (Arrays)
The slice type is an abstraction built on top of Go's array type, and so to understand slices we must first understand arrays.

An array type definition specifies a length and an element type. For example, the type [4]int represents an array of four integers. An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types). Arrays can be indexed in the usual way, so the expression s[n] accesses the nth element, starting from zero. (https://blog.golang.org/go-slices-usage-and-internals)

To initilize:
```go
var a [4]int
a[0] = 1
i := a[0]
// i == 1
```
Another way to make an array of a type:
```go
studentNameArray = make([]string,0) //Returns a slice of studentNameArray with type string with an initial length of zero but NO LIMIT TO CAPACITY
anotherStudentNameArray = make([]string,0,10) //Returns a slice of studentNameArray with type string with an initial length of zero and a capacity of 10 elements
// Docuemntation https://golang.org/pkg/builtin/#make
```
### For Loops
There are only FOR loops in Go. So what do we do?
There are many ways of using the for keyword. You can use it to do iterations, or conditional interations. Sounds familiar? You may be doing a O(n) loop or a O(nlogn) loop. But the key idea to remember here is that you are not restricted by the traditional "3 variable" loop. (initializer, condition, counter). Let's take a look at the different types of loops.

```go
// Your typical for loop with the 3 variables needed in the condition statements
sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum)
//----------------------------

// An emulation of while loop using for keyword
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n)
//----------------------------

// Endless Loop
sum := 0
for {
    sum++ // repeated forever
}
fmt.Println(sum)
//----------------------------

// This is similar to foreach loop, i will give index, while s will give value.
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
//----------------------------

// A loop with loop conditions
sum := 0
for i := 1; i < 5; i++ {
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
}
fmt.Println(sum)
```

### Structs

### Pointers

### Function

### go 'commands'

