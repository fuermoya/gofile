//package main
//
//import (
//	"fmt"
//	"gofile/server"
//	"net/http"
//	"time"
//)
//
//func main() {
//	var sr = server.Server{}
//	router := sr.Routers()
//	address := fmt.Sprintf(":%d", 8088)
//	var s = http.Server{
//		Addr:           address,
//		Handler:        router,
//		ReadTimeout:    20 * time.Second,
//		WriteTimeout:   20 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	// 保证文本顺序输出
//	time.Sleep(2 * time.Microsecond)
//	fmt.Printf(`
//	欢迎使用 gofile
//	访问地址:http://127.0.0.1%s
//	`, address)
//	err := s.ListenAndServe().Error()
//	fmt.Println(err)
//}

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"gofile/server"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

// 获取可执行文件的绝对路径
func ExecPath() string {
	file, e := os.Executable()
	if e != nil {
		log.Printf("Executable file path error : %s\n", e.Error())
	}
	path := filepath.Dir(file)
	return path
}

// 获取 service 对象
func getSrv() service.Service {
	File, err := os.Create(ExecPath() + "/http-server.log")
	if err != nil {
		File = os.Stdout
	}
	defer File.Close()

	log.SetOutput(File)

	s := &Services{
		Cfg: &service.Config{
			Name:        "GoFileService",
			DisplayName: "GoFileService",
			Description: "基于gin的file服务",
		}}
	serv, er := service.New(s, s.Cfg)
	if er != nil {
		log.Printf("Set logger error:%s\n", er.Error())
	}
	s.Log, er = serv.SystemLogger(nil)
	return serv
}

// 启动windows服务
func (srv *Services) Start(s service.Service) error {
	//if srv.Log != nil {
	srv.Log.Info("Start run http server")
	//}
	go srv.StarServer()
	return nil
}

// 停止windows服务
func (srv *Services) Stop(s service.Service) error {
	if srv.Log != nil {
		srv.Log.Info("Start stop http server")
	}
	log.Println("Server exiting")
	return srv.Srv.Shutdown(context.Background())
}

// 运行gin web服务
func (srv *Services) StarServer() {
	gin.DisableConsoleColor()
	var sr = server.Server{}
	router := sr.Routers()
	address := fmt.Sprintf("0.0.0.0:%d", 8088)
	f, _ := os.Create(ExecPath() + "/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	srv.Srv = &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	time.Sleep(2 * time.Microsecond)
	fmt.Printf(`
		欢迎使用 gofile
		访问地址:http://127.0.0.1%s
		`, address)
	err := srv.Srv.ListenAndServe().Error()
	srv.Log.Error(err)
}

func main() {
	s := getSrv()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := s.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
			fmt.Printf("服务已安装")
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
			fmt.Printf("服务已卸载")
		case "start":
			err := s.Start()
			if err != nil {
				log.Fatalf("Start service error:%s\n", err.Error())
			}
			fmt.Printf("服务已启动")
		case "stop":
			err := s.Stop()
			if err != nil {
				log.Fatalf("top service error:%s\n", err.Error())
			}
			fmt.Printf("服务已关闭")
		}
		return
	}
	err := s.Run()
	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}
