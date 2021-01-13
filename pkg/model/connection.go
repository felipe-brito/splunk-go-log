package model

// 	HECConnection :
// 	HECSettings -
// 	ActiveConfiguration -
type HECConnection struct {
	HECSettings         HECSettings
	ActiveConfiguration bool
}

// 	TCPConnection :
// 	TCPSettings -
// 	ActiveConfiguration -
type TCPConnection struct {
	TCPSettings         TCPSettings
	ActiveConfiguration bool
}
