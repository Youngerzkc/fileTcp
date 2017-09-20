package main
import(
	"fmt"
	"flag"
	"log"
)
var (
	// 命令行解析参数
	mode=flag.String("mode","","运行模式")
	port=flag.Int("port", 5050,"服务器监听端口")
	file=flag.String("file","","文件名称")
)
func main(){
	flag.Parse()//解析命令行
	// ./fileTcp  -mode="dev" -file="filename"
	fmt.Println(*mode,*port,*file)
	switch *mode {
	case "server":
		runServer(*port)
	case "client":
		runClient(*port,*file)
	default:

		log.Fatalf("未知的运行模式:%v",*mode)	
	}

}
