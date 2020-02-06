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