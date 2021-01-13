package tcp

import (
	"errors"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/work"
)

var w *work.Work

func InitializeTcpService(tcpConnection model.TCPConnection) {
	w = work.New(tcpConnection.TCPSettings.WorkSettings.NumberOfWorkers)
	w.StartWorkers()
}

func Send(obj interface{}) error {
	if w == nil {
		return errors.New("the message receiving channel has not been opened. Check your implementation to make sure the connection is initialized. See the example section on github.com to understand how to initialize the service")
	}

	w.SendToWork(obj)
	return nil
}
