package main

import (
  "fmt"
)

func main(){
  var LibraryName string 
  LibraryName = "伊苏"
  message := fmt.Sprintf("这里是%v图书管理系统", LibraryName)
  fmt.Println(message)
}

