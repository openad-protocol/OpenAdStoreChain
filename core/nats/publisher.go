package nats

import (
	"AdServerCollector/logger"
)

func (s *AdNatsConnection) Publish(subject string, msg []byte) {
	if s.useStream {
		ack, err := s.js.Publish(subject, msg)
		if err != nil {
			logger.Error()
		}
		logger.Infof("ack info %v", ack)
	} else {
		err := s.conn.Publish(subject, msg)
		if err != nil {
			logger.Errorf("Error publishing message: %v", err)
		}
	}
	logger.Infof("Published message on subject %s: %s\n", subject, string(msg))
}
