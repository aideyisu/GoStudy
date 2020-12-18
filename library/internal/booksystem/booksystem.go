package booksystem

import "fmt"

type OneBook struct{
  Auther string 
  CB bool
  Borrower string
}


func AddBook(LibrarySystem map[string]OneBook,BookName string, Auther string) (map[string]OneBook, error) {
  _, ok := LibrarySystem["BookName"]
  if ok == true{ // 冲突检测
    fmt.Printf("这本书已经存在")
    return LibrarySystem, nil // 这里不应该是nil,确认下错误报错格式再来写
  }

  var NewBook OneBook
  NewBook.Auther = Auther
  NewBook.CB = true

  LibrarySystem[BookName] = NewBook
  return LibrarySystem, nil
}

func BorrowBook(LibrarySystem map[string]OneBook, BookName string, Borrower string)(map[string]OneBook, error) {

  _, ok := LibrarySystem[BookName]
  if ok == false{
    fmt.Printf("此书不存在")
    return LibrarySystem, nil // 这里不应该是nil,确认下错误报错格式再来写
  }
  if LibrarySystem[BookName].CB == false {
    fmt.Printf("此书已被借阅")
    return LibrarySystem, nil // 这里不应该是nil,确认下错误报错格式再来写
  }
  BBook := LibrarySystem[BookName]
  BBook.CB = false
  BBook.Borrower = Borrower
  
  LibrarySystem[BookName] = BBook
  // fmt.Printf("借阅书籍 %v-%v-%v \n", BookName, Borrower, TimeLength)
  return LibrarySystem, nil
}

func Init() map[string]OneBook{
  NewLibrary := make(map[string]OneBook)
  fmt.Println("系统启动完成")
  return NewLibrary
}

