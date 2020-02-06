# 0x03 - Go Crazy

Now we have created simple app,
Lets try to create the following use case.

Given a range of 'n' from 1 - 100, how can we create a service that returns:
{
    "1":"1",
    "2":"2",
    ..
    ..
    ..
    "100":"100"
}

```go
func ReturnCrazy() map[string]interface{} {
    output := make(map[string]interface{},0)
    for i := 1; i < 101; i++ {
        num := strconv.Itoa(i)
        output[num] = num
    }
    return output
}
```

Now can you use the above example and return the following?

Use the following slice:

```go
albums := []string{"Taylor Swift", "Fearless", "Speak Now", "Red", "1989", "Reputation", "Lover"}
sales := []string{"100","200","300","400","500","600","700"}
```

And the output should look like this:
{
    "Taylor Swift":"100",
    ..
    ..
    "Lover":"700"
}