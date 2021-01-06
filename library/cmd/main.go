package main

import (
  "fmt"
  "log"

  "greetings"
  bs "booksystem"
)

func main(){
  LibraryName := "酥酥小黑共同守护的图书馆"
  UserName := "蜜桃芙"
  BookName := "大货车的秘密"
  
  log.SetPrefix("greeting: ")
  log.SetFlags(0)
  fmt.Println("%v - %v", UserName, LibraryName)
  BS := bs.Init()

  // message,err := greetings.Hello(UserName)
  message,err := greetings.Hello("TT")
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(message)

  // 人名切片
  names := []string{"Gladys", "Samantha", "Darrin"}
  messages,errs := greetings.Hellos(names)
  if errs != nil{
    log.Fatal(errs)
  }

  fmt.Println(messages)

  BS, ER := bs.AddBook(BS, BookName, "旭旭宝宝")
  if ER != nil {
    fmt.Println("入库失败请查看原因")
  }
  fmt.Println(BS)


  BS, ER = bs.BorrowBook(BS, BookName, "aideyisu")

  if ER != nil {
    fmt.Println("借阅失败请查看原因")
  }
  fmt.Println(BS)
}

