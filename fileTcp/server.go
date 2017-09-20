package main
import (
	"os"
	"log"
	"net"
	"strconv"
	"io"
)
func  handler( conn net.Conn )  {
	defer conn.Close()
	remoteAddr:=conn.RemoteAddr().String()
	log.Println("远程IP地址:",remoteAddr)
	// 获取头文件信息。例如文件名
	p:=make([]byte,1024)
	// p 默认实际长度1024
	// p2:=make([]byte,0,1024)
	// p2-->实际长度0
	n,err:=conn.Read(p)
	if err!=nil{
		log.Printf("读取头文件失败(%s):%v",remoteAddr,err)
		return
	}else if n==0{
		log.Printf("空头文件(%s)",remoteAddr)
		return 
	}
	fileName:=string(p[:n])
	log.Printf("文件:(%s)",fileName)
	// 服务端接收到，返回给客户端信息
	conn.Write([]byte("ok"))
	// 创建文件目录
	os.MkdirAll("receive",os.ModePerm)
	// 网络流拷贝到本地文件
	f,err:=os.Create("receive/"+fileName)
	if err!=nil{
		log.Println("服务端无法创建文章")
	}
	defer f.Close()
	_,err=io.Copy(f,conn)
	for {
		buffer:=make([]byte,1024*200)
		// 一次传送200KB，Read先创建空间
		_,err:=conn.Read(buffer)
		if err!=nil&&err!=io.EOF{
			log.Printf("读取失败（%s）:%v",remoteAddr,err)
		}else if err==io.EOF{
			break
		}
	}
	if err!=nil{
		log.Printf("文件接收失败（%s）：%v",remoteAddr,err)
			return
	}
	log.Printf("文件接收成功(%s):%s",remoteAddr,fileName)

}
func  runServer(port int )  {
	//启动监听
	l,err:=net.Listen("tcp",":"+strconv.Itoa(port))
	if err!=nil{
		log.Fatalf("监听失败")
	}
	log.Println("服务端已启动")
	// 循环接收请求
	for {
		conn,err:=l.Accept()
		if err!=nil{
			// 类型断言，是否为网络错误err.(net.Error)
			 if ne,ok:=err.(net.Error);!ok||ne.Temporary(){
				log.Printf("接收请求失败")
				 }	
				 continue //放弃当前请求，继续监听
			}
			log.Println("请求接收成功")
			// _=conn
			go handler(conn)
	}
	
}