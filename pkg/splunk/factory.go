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
// 	setValidConfiguration -
type Connection interface {
	Configuration(settings interface{}) error
	Initialize()
	setValidConfiguration(validConfiguration bool)
}

// 	GetInstance :
// 	serviceType -
func GetInstance(serviceType ServiceType) Connection {
	switch serviceType {
	case HEC:
		return &HecConnectionImpl{}
	case TCP:
		return &TcpConnectionImpl{}
	default:
		return &TcpConnectionImpl{}
	}
}
