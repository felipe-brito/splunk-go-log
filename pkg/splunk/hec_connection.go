package splunk

import (
	"errors"
	"github.com/felipe-brito/splunk-go-log/pkg/hec"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/util"
)

type HECConnectionImpl struct {
	hecConnection model.HECConnection
}

func (h *HECConnectionImpl) Configuration(settings interface{}) error {

	util.Log().Info("configuring HEC type connection in host:", util.HostName())
	hecSettings, success := settings.(model.HECSettings)
	if !success {
		return errors.New("the informed configuration is not a valid configuration. To configure the HTTP Event Collector use the HECSettings configuration structure")
	}

	util.Log().Info("splunk server in url:", hecSettings.Url)

	h.setActiveConfiguration(true)
	h.hecConnection.HECSettings = hecSettings

	return nil
}

func (h *HECConnectionImpl) Initialize() error {

	util.Log().Info("initializing the connection to the splunk server")
	if !h.hecConnection.ActiveConfiguration {
		return errors.New("the service could not be started because no valid configuration was entered. Enter a service configuration and try again")
	}

	go func(hecConnection model.HECConnection) {
		hec.InitializeHECService(h.hecConnection)
	}(h.hecConnection)

	return nil
}

func (h *HECConnectionImpl) setActiveConfiguration(activeConfiguration bool) {
	h.hecConnection.ActiveConfiguration = activeConfiguration
}
