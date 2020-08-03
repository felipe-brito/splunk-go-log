package splunk

import (
	"errors"
	"github.com/felipe-brito/splunk-go-log/pkg/hec"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/tcp"
)

type HecConnectionImpl struct {
	hecConnection model.HecConnection
}

type TcpConnectionImpl struct {
	tcpConnection model.TcpConnection
}

func (h *HecConnectionImpl) Configuration(settings interface{}) error {

	var err error = nil

	hecSettings, success := settings.(model.HecSettings)
	if success {
		h.setValidConfiguration(true)
		h.hecConnection.HecSettings = hecSettings
	} else {
		err = errors.New("the informed configuration is not a valid configuration. To configure the Http Event Collector use the HecSettings configuration structure")
	}
	return err
}

func (h *HecConnectionImpl) Initialize() {
	go func(hecConnection model.HecConnection) {
		hec.InitializeHecService(h.hecConnection)
	}(h.hecConnection)
}

func (h *HecConnectionImpl) setValidConfiguration(validConfiguration bool) {
	h.hecConnection.ActiveConfiguration = validConfiguration
}

func (t *TcpConnectionImpl) Configuration(settings interface{}) error {

	var err error = nil

	tcpSettings, success := settings.(model.TcpSettings)
	if success {
		t.setValidConfiguration(true)
		t.tcpConnection.TcpSettings = tcpSettings
	} else {
		err = errors.New("the informed configuration is not a valid configuration. To configure the Http Event Collector use the TcpSettings configuration structure")

	}

	return err
}

func (t *TcpConnectionImpl) Initialize() {
	go func(tcpConnection model.TcpConnection) {
		tcp.InitializeTcpService(t.tcpConnection)
	}(t.tcpConnection)
}

func (t *TcpConnectionImpl) setValidConfiguration(validConfiguration bool) {
	t.tcpConnection.ActiveConfiguration = validConfiguration
}
