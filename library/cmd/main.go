package main

import (
  "fmt"

  "greetings"
  bs "booksystem"
)

func main(){
  LibraryName := "荒谷遗精图书馆"
  UserName := "国服增幅王"
  BookName := "增幅的秘密"

  bs.Init()

  message2 := greetings.Hello(UserName)
  fmt.Println(message2)

  message := greetings.ComeLibrary(LibraryName)
  fmt.Println(message)

  bs.AddBook(BookName)
  bs.BorrowBook(BookName, UserName, 3)
}

