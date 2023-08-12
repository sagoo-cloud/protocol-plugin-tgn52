package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/extend"
	"github.com/sagoo-cloud/sagooiot/extend/module"
	"net"
	"testing"
)

func TestManagerInit(t *testing.T) {
	manager := extend.NewManager("protocol", "protocol-*", "./built", &module.ProtocolPlugin{})
	defer manager.Dispose()
	err := manager.Init()
	if err != nil {
		t.Fatal(err.Error())
	}
	err = manager.Launch()
	for id, info := range manager.Plugins {
		t.Log(id)
		t.Log(info.Path)
		t.Log(info.Client)
	}

	p, err := manager.GetInterface("tgn52")
	if err != nil {
		t.Fatal(err.Error())
	}

	data := gconv.Bytes("NB1;1234567;1;2;+25.5;00;030;+21;+22")

	obj := p.(module.Protocol)
	res, err := obj.Decode(data, "")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res)

}

// 测试插件服务使用，需要先将要测试的插件进行编译
func TestProtocolPluginServer(t *testing.T) {
	extend.GetProtocolPlugin()
	NetData()
}

func NetData() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	//获取插件
	//pm := GetPlugin(ProtocolPluginName)

	for {
		buf := make([]byte, 512)
		l, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:l]))

		//获取协议插件解析后的数据 传入插件ID，及需要解析的数据
		data, err := extend.GetProtocolPlugin().GetProtocolUnpackData("tgn52", buf[:l])
		fmt.Println("============通过插件获取数据：", data)
	}
}
