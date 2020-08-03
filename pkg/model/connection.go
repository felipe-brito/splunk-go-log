package model

// 	HecConnection :
// 	hecSettings -
// 	isValidConfiguration -
type HecConnection struct {
	HecSettings         HecSettings
	ActiveConfiguration bool
}

// 	TcpConnection :
// 	tcpSettings -
// 	isValidConfiguration -
type TcpConnection struct {
	TcpSettings         TcpSettings
	ActiveConfiguration bool
}
