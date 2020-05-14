package main

import (
	"fmt"
	"io"
	"net"
)

//tcp server 服务端代码

func main() {
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr
	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")
	//循环接收客户端的连接，创建一个协程具体去处理连接
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

//具体处理连接过程方法
func tcpPipe(conn *net.TCPConn) {
	//tcp连接的地址
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	//获取一个连接的reader读取流
	//reader := bufio.NewReader(conn)

	//接收并返回消息
	for {
		reader := make([]byte, 2048)
		_, err := conn.Read(reader)
		if err != nil || err == io.EOF {
			break
		}
		// //buf数据 -> 去两端空格的string
		// inStr := strings.TrimSpace(string(reader[0:message]))
		// //去除 string 内部空格
		// cInputs := strings.Split(inStr, " ")
		//fmt.Println("客户端传输->" + cInputs[0])
		fmt.Println("客户端传输->" + string(reader))
		conn.Write([]byte(string(reader)))
	}
}
