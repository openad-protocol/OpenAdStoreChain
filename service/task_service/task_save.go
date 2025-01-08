package task_service

import (
	"AdServerCollector/core/attestation"
	"AdServerCollector/core/task"
	"sync"
)

type TaskService struct {
	Name      string
	TaskEntry *task.TaskRun
}

func NewTaskService() *TaskService {
	return &TaskService{
		Name:      "SaveHashTask",
		TaskEntry: task.NewTaskRun(),
	}
}

func (t *TaskService) StartService(wg *sync.WaitGroup) {
	t.TaskEntry.RunTask(wg)
}

func (t *TaskService) AddTask(spec string, cmd func()) {
	t.TaskEntry.AddTask(spec, cmd)
}

func (t *TaskService) StopService() {
	t.TaskEntry.Stop()
	t.closeTaskSaveAll()
}

func (t *TaskService) GetName() string {
	return t.Name
}

func (t *TaskService) SaveHashTask() {
	t.AddTask("*/5 * * * *", func() {
		// 每5分钟执行一次
		attestation.LogTree.SaveVersionStateDB(false)
		attestation.CallBackTree.SaveVersionStateDB(false)
		attestation.ClickTree.SaveVersionStateDB(false)
		attestation.GetAdTree.SaveVersionStateDB(false)
	})
}

func (t *TaskService) closeTaskSaveAll() {
	attestation.LogTree.SaveVersionStateDB(false)
	attestation.CallBackTree.SaveVersionStateDB(false)
	attestation.ClickTree.SaveVersionStateDB(false)
	attestation.GetAdTree.SaveVersionStateDB(false)
}
