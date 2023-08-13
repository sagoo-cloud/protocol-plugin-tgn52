package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/extend"
	"github.com/sagoo-cloud/sagooiot/extend/model"
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
	t.Log(manager)
}

// 测试获取插件信息
func TestProtocolInfo(t *testing.T) {
	p, err := extend.GetProtocolPlugin().GetProtocolPlugin("tgn52")
	if err != nil {
		return
	}
	t.Log(p.Info())
}

type TestData struct {
	Name  string
	Value string
}

// 测试协议的编码方法
func TestProtocolEncode(t *testing.T) {
	p, err := extend.GetProtocolPlugin().GetProtocolPlugin("tgn52")
	if err != nil {
		t.Fatal(err.Error())
	}
	td := new(TestData)
	td.Name = "aaaa"
	td.Value = "bbbbb"
	res := p.Encode(td)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res)
}

// 测试自定义协议解析
func TestProtocol(t *testing.T) {
	data := gconv.Bytes("NB1;1234567;1;2;+25.5;00;030;+21;+22")
	p, err := extend.GetProtocolPlugin().GetProtocolPlugin("tgn52")
	if err != nil {
		return
	}
	var dr model.DataReq
	dr.Data = data
	res := p.Decode(dr)
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
