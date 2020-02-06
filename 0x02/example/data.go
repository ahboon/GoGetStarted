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