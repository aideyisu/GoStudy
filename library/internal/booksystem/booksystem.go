package booksystem


import "fmt"

func AddBook(BookName string){
    fmt.Printf("入库书籍%v \n",BookName)
}

func BorrowBook(BookName string, Borrower string, TimeLength int){
  fmt.Printf("借阅书籍 %v-%v-%v \n", BookName, Borrower, TimeLength)
}

func Init(){
  fmt.Println("系统启动完成")
}

