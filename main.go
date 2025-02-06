package main

import (
	"AdServerCollector/conf"
	"AdServerCollector/logger"
	"AdServerCollector/service/message"
	"AdServerCollector/service/task_service"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {

	// 读配置
	//shutdownCh := utils.MakeShutdownCh()
	natsHost := conf.Config.Nats.Host
	natsPort := conf.Config.Nats.Port
	natsUrl := "nats://" + natsHost + ":" + strconv.Itoa(natsPort)
	csName := conf.Config.Nats.ConsumerName
	useStream := conf.Config.Nats.UseJetStream

	// 初始化服务
	q := message.NewNatsQueen(natsUrl, csName, useStream)
	for _, v := range conf.Config.Nats.Subjects {
		bAddResult := q.AddSubscription(v)
		if bAddResult {
			logger.Infof("Add subscription %s success", v)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 启动服务
	go q.Start(&wg)
	wg.Add(1)
	task := task_service.NewTaskService()
	task.SaveHashTask()
	go task.StartService(&wg)
	//// 设置信号监控
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-shutdownCh
		logger.Infof("Received signal: %s", sig)
		// 执行清理操作
		q.Stop()
		task.StopService()
	}()

	wg.Wait()
}
