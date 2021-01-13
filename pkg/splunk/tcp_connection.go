package splunk

import (
	"errors"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/tcp"
	"github.com/felipe-brito/splunk-go-log/pkg/util"
)

type TCPConnectionImpl struct {
	tcpConnection model.TCPConnection
}

func (t *TCPConnectionImpl) Configuration(settings interface{}) error {

	util.Log().Info("configuring TCP type connection in host:", util.HostName())
	tcpSettings, success := settings.(model.TCPSettings)
	if !success {
		return errors.New("the informed configuration is not a valid configuration. To configure the TCP Collector use the TCPSettings configuration structure")
	}

	util.Log().Info("splunk server in url:", tcpSettings.Url)

	t.setActiveConfiguration(true)
	t.tcpConnection.TCPSettings = tcpSettings

	return nil
}

func (t *TCPConnectionImpl) Initialize() error {

	util.Log().Info("initializing the connection to the splunk server")
	if !t.tcpConnection.ActiveConfiguration {
		return errors.New("the service could not be started because no valid configuration was entered. Enter a service configuration and try again")
	}
	go func(tcpConnection model.TCPConnection) {
		tcp.InitializeTcpService(t.tcpConnection)
	}(t.tcpConnection)

	return nil
}

func (t *TCPConnectionImpl) setActiveConfiguration(activeConfiguration bool) {
	t.tcpConnection.ActiveConfiguration = activeConfiguration
}
