package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(conn)

	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s", clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") // Windows 平台下用 "\r\n"，Linux平台下使用 "\n"

	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--%s--", input)
		// fmt.Printf("trimmedInput:--%s--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		wr.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n"))
		wr.Flush()
		var str string = ""

		var S []byte = []byte(str)
		line, err := rd.Read(S)
		if err != nil {
			log.Printf("err ! %v", err)
			return
		}
		fmt.Println(line, S)
		// _, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
