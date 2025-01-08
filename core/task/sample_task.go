package task

import (
	"AdServerCollector/logger"
	cron "github.com/robfig/cron/v3"
	"sync"
)

type TaskRun struct {
	cron *cron.Cron
	ch   chan string
}

func NewTaskRun() *TaskRun {
	return &TaskRun{
		cron: cron.New(),
		ch:   make(chan string),
	}
}

func (t *TaskRun) RunTask(wg *sync.WaitGroup) {
	defer wg.Done()
	// 启动 Cron 调度器
	t.cron.Start()
	// 运行一段时间以观察任务执行情况
	select {
	case <-t.ch:
		logger.Info("TaskRun stopped")
		t.cron.Stop()
		return
	}
}

func (t *TaskRun) AddTask(spec string, cmd func()) {
	_, err := t.cron.AddFunc(spec, cmd)
	if err != nil {
		logger.Error("Error adding cron function:", err)
		return
	}
}

func (t *TaskRun) Stop() {
	close(t.ch)
}
