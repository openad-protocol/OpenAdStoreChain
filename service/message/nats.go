package message

import (
	adNats "AdServerCollector/core/nats"
	"AdServerCollector/logger"
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
)

type NatsQueen struct {
	url       string
	conn      *adNats.AdNatsConnection
	subList   []*nats.Subscription
	mp        *MProcess
	isRun     bool
	method    string
	csName    string
	useStream bool
}

func NewNatsQueen(url, csName string) IMessage {
	useStream := false
	conn, err := adNats.NewNatsConnect(url, useStream)
	if err != nil {
		panic(fmt.Sprintf("Connection to Nats Error %v", err))
	}
	// 绑定消息处理器
	mp := NewMProcess()
	return &NatsQueen{url: url,
		conn:      conn,
		isRun:     false,
		mp:        mp,
		csName:    csName,
		useStream: useStream,
	}
}

func (n *NatsQueen) Start(wg *sync.WaitGroup) bool {
	defer wg.Done()
	// 创建流
	if n.useStream {
		if err := n.conn.AddStream("ad_info", []string{"ad_info.*"}); err != nil {
			logger.Error("create stream error: %v\n", err)
			panic(fmt.Sprintf("create stream error: %v\n", err))
		}
	}
	n.isRun = true
	for {
		if !n.isRun {
			break
		}
		if !n.conn.IsConnected() {
			logger.Info("NATS 连接丢失，尝试重新连接...")
			conn, err := adNats.NewNatsConnect(n.url, n.useStream)
			if err != nil {
				logger.Errorf("重新连接失败: %v\n", err)
				continue
			}
			n.conn = conn
		}
		time.Sleep(5 * time.Second)
	}
	return true

}

func (n *NatsQueen) Stop() bool {
	n.isRun = false
	logger.Infof("close nats connection!")
	n.conn.Close()
	return true
}

func (n *NatsQueen) Publish(subject string, msg []byte) bool {
	//todo:实现发布消息
	return true
}

func (n *NatsQueen) AddSubscription(subject string) bool {
	switch subject {
	case "ad_info.tracerHash":
		_s, err := n.Subscribe(subject, n.mp.ProcessTracerHash) // 处理跟踪hash
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)
	case "ad_info.get_ad_miss":
		_s, err := n.Subscribe(subject, n.mp.ProcessAdMissing) // 处理丢失的广告信息
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)
	case "ad_info.get_ad":
		_s, err := n.Subscribe(subject, n.mp.ProcessAdMessage) // 处理广告信息
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)
	case "ad_info.clickinfo":
		_s, err := n.Subscribe(subject, n.mp.ProcessAdMessage) // 处理广告信息
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)
	case "ad_info.loginfo":
		_s, err := n.Subscribe(subject, n.mp.ProcessAdMessage) // 处理广告信息
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)
	case "ad_info.ad_in_call":
		_s, err := n.Subscribe(subject, n.mp.ProcessAdMessage) // 处理广告信息
		if err != nil {
			logger.Error("subject %s 订阅失败: %s\n", subject, err)
			return false
		}
		n.subList = append(n.subList, _s)

	default:
		logger.Error("未知的主题: %s\n", subject)
	}
	return true
}

func (n *NatsQueen) Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	_sub, err := n.conn.Subscribe(subject, n.csName, handler)
	if err != nil {
		logger.Errorf("Failed to subscribe to subject: %s,durableName:%s err: %s\n", subject, n.csName, err)
	}
	return _sub, err
}

func (n *NatsQueen) Unsubscribe(sub *nats.Subscription) bool {
	err := sub.Unsubscribe()
	if err != nil {
		logger.Infof("Failed to unsubscribe: %v\n", err)
		return false
	}
	logger.Debugf("Unsubscribed %s\n", sub.Subject)
	return true
}

func (n *NatsQueen) UnsubscribeAll() bool {
	for _, sub := range n.subList {
		if !n.Unsubscribe(sub) {
			logger.Errorf("Failed to unsubscribe all")
			return false
		}
		logger.Infof("Unsubscribed %s\n", sub.Subject)
	}
	return true
}
