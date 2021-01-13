package splunk

const (
	HEC = ServiceType("HEC")
	TCP = ServiceType("TCP")
)

// 	ServiceType -
type ServiceType string

// 	Connection :
// 	Configuration -
// 	Initialize -
// 	setActiveConfiguration -
type Connection interface {
	Configuration(settings interface{}) error
	Initialize() error
	setActiveConfiguration(activeConfiguration bool)
}

// 	GetInstance :
// 	serviceType -
func GetInstance(serviceType ServiceType) Connection {
	switch serviceType {
	case HEC:
		return &HECConnectionImpl{}
	case TCP:
		return &TCPConnectionImpl{}
	default:
		return &TCPConnectionImpl{}
	}
}
