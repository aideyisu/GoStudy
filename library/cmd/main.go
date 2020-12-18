package main

import (
  "fmt"

  "greetings"
  bs "booksystem"
)

func main(){
  LibraryName := "酥酥小黑共同守护的图书馆"
  UserName := "蜜桃芙"
  BookName := "大货车的秘密"

  BS := bs.Init()

  message2 := greetings.Hello(UserName)
  fmt.Println(message2)

  message := greetings.ComeLibrary(LibraryName)
  fmt.Println(message)

  BS, ER := bs.AddBook(BS, BookName, "AABB")
  if ER != nil {
    fmt.Println("入库失败请查看原因")
  }
  fmt.Println(BS)


  BS, ER = bs.BorrowBook(BS, BookName, "aideyisu")

  if ER != nil {
    fmt.Println("入库失败请查看原因")
  }
  fmt.Println(BS)
}

