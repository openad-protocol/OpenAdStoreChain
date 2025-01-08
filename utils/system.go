package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// MakeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// SIGINT or SIGTERM received.
func MakeShutdownCh() chan struct{} {
	resultCh := make(chan struct{})

	shutdownCh := make(chan os.Signal, 4)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdownCh
		close(resultCh)
	}()
	return resultCh
}

// MakeSighupCh returns a channel that can be used for SIGHUP
// reloading. This channel will send a message for every
// SIGHUP received.
func MakeSighupCh() chan struct{} {
	resultCh := make(chan struct{})

	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, syscall.SIGHUP)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()
	return resultCh
}

func SignalHandler() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan) // 监听所有信号。注意，Notify接收可变参数，可以指定监听信号。
	log.Println("signal process handler")
	go func() {
		for {
			sig := <-signalChan // 监听到信号
			log.Printf("got signal to exit [signal = %v]", sig)
		}
	}()
}
