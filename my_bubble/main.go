package main

import (
	"context"
	"fmt"
	"log"
	"my_bubble/configs"
	"my_bubble/dao/mysql"
	"my_bubble/logger"
	"my_bubble/routers"
	"my_bubble/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 获取配置
	if err := settings.Init("settings/config.yaml"); err != nil {
		panic("Get config failed")
	}
	fmt.Printf("%#v\n", configs.Conf)
	fmt.Println("Init Settings successfully")

	// 初始化 logger
	if err := logger.Init(configs.Conf.LogConf, configs.Conf.AppConf.Mode); err != nil {
		fmt.Printf("logger, err: %v\n", err)
		zap.L().Error("Init Logger failed\n", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	fmt.Println("Init logger successfully")

	// 初始化MySql
	if err := mysql.Init(configs.Conf.MysqlConf); err != nil {
		zap.L().Error("Init Mysql failed\n", zap.Error(err))
		return
	}
	fmt.Println("Init MySql successfully")

	// 注册路由
	r := routers.Setup(configs.Conf.AppConf.Mode)
	fmt.Println("Setup routers successfully")

	// 优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	fmt.Println("Ready for Server...:", viper.GetInt("app.port"))

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
