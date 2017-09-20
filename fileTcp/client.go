package main
import(
	"net"
	"log"
	"strconv"
	"os"
	"io"
)
func  runClient( port int ,file string)  {
	// 连接服务器
  conn,err:= net.Dial("tcp",":"+strconv.Itoa(port))
  if err!=nil{
	  log.Printf("无法建立链接:%v",err)
	  return 
  }
//   函数返回前释放资源
  defer conn.Close()
  log.Println("连接建立成功")
	// 打开要传送的文件
	f,err:=os.Open(file)
if err!=nil{
	log.Printf("无法打开文件：%s",file)
	return
}
 defer f.Close()
//  写入头信息,并等待确认
	conn.Write([]byte(file))
	// 读取服务端返回的信息，服务端返回OK，2个
	p:=make([]byte,2)
	_,err=conn.Read(p)
	if err!=nil{
		log.Printf("无法获得服务端信息：%v",err)
		return 
	}else if string(p)!="ok" {
		log.Printf("无效的服务端相应：%s",string(p))
		return 
	}
	log.Println("头信息发送成功！")
	// 本地流拷贝到网络流
	_,err =io.Copy(conn,f)
	if err!=nil{
		log.Printf("发送文件失败(%s):%v",file,err)
			return 
	}
	log.Println("文件发送成功！")
}