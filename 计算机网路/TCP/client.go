package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	// var tcpAddr *net.TCPAddr
	// tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	var wg sync.WaitGroup
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " : Client connected!")

	go onMessageReceived(conn, wg)
	go onMessageRead(conn, wg)

	wg.Add(2)
	wg.Wait()
}
func onMessageRead(c net.Conn, wg sync.WaitGroup) {
	for {
		//缓存conn中的数据
		buf := make([]byte, 2048)
		//服务器端返回的数据写入空buf
		cnt, err := c.Read(buf)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}

		//回显服务器端回传的信息
		fmt.Print("服务器端回复" + string(buf[0:cnt]))
	}
	wg.Done()
}

func onMessageReceived(c net.Conn, wg sync.WaitGroup) {

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		//去除输入两端空格
		input := strings.TrimSpace(msg)
		//客户端请求数据写入 conn，并传输
		c.Write([]byte(input))

	}
	wg.Done()
}
