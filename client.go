package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //当前客户端模式
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       -1,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	return client
}

// 客户端菜单选项
func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公屏聊天")
	fmt.Println("2.私聊")
	fmt.Println("3.更改用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围的数字<<<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		switch client.flag {
		case 1:
			//公屏聊天
			fmt.Println("公聊模式选择")
			break
		case 2:
			//私聊
			fmt.Println("私聊模式选择")
			break
		case 3:
			//修改名称
			fmt.Println("修改名称")
			break
		}
	}
}

// 用户接收和初始化ip and port
var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器默认ip地址为127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器默认端口为8888")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>链接服务器失败...")
		return
	}
	fmt.Println(">>>服务器链接成功...")
	//启动客户端的业务
	client.Run()
}
