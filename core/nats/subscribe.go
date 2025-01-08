package nats

import (
	"AdServerCollector/logger"
	"github.com/nats-io/nats.go"
	"time"
)

type AdNatsConnection struct {
	conn       *nats.Conn
	name       string
	js         nats.JetStreamContext
	streamInfo *nats.StreamInfo
	useStream  bool
}

func NewNatsConnect(url string, useStream bool) (*AdNatsConnection, error) {
	conn, err := nats.Connect(url,
		nats.MaxReconnects(10), // 最大重连次数
		//nats.DontRandomize(),  // 不随机连接
		nats.ReconnectWait(5*time.Second), // 重连等待时间
	)
	if err != nil {
		logger.Error("nats connection error:", err)
		return nil, err
	}
	sub := &AdNatsConnection{
		name:      "adserver",
		conn:      conn,
		useStream: useStream,
	}
	if useStream {
		js, err := conn.JetStream(nats.PublishAsyncMaxPending(100))
		if err != nil {
			logger.Error("nats jet stream conn:", err)
			return nil, err
		}
		sub.js = js
	}
	return sub, nil
}

func (s *AdNatsConnection) IsConnected() bool {
	return s.conn.IsConnected()
}

func (s *AdNatsConnection) Subscribe(subject, durableName string, handler func(msg *nats.Msg)) (_subject *nats.Subscription, err error) {
	if s.useStream {
		_sub, err := s.js.QueueSubscribe(subject, durableName, handler, nats.Durable(durableName))
		if err != nil {
			logger.Errorf("Error subscribing to subject: %s", err.Error())
			return nil, err
		}
		logger.Infof("Subscribed to subject %s with durable name %s", subject, durableName)
		return _sub, nil
	} else {
		if _subject, err = s.conn.QueueSubscribe(subject, durableName, handler); err != nil {
			logger.Errorf("Error subscribing to subject: %s", err.Error())
			return nil, err
		}
	}
	logger.Infof("Subscribed to subject %s", subject)
	return
}

func (s *AdNatsConnection) Close() {
	s.conn.Close()
}

func (s *AdNatsConnection) AddStream(streamName string, subjects []string) error {
	stream, err := s.js.StreamInfo(streamName)
	if err != nil && stream != nil {
		return err
	}
	if stream == nil {
		s.streamInfo, err = s.js.AddStream(&nats.StreamConfig{
			Name:      streamName,
			Subjects:  subjects,
			Retention: nats.WorkQueuePolicy, // 使用工作队列策略，确保每条消息只能被消费一次
		})
		if err != nil {
			return err
		}
	}
	return nil
}
