package model

// 	HECSettings :
// 	Url -
// 	Source -
// 	SourceType -
// 	Index -
type HECSettings struct {
	Url          string
	Source       string
	SourceType   string
	Index        string
	WorkSettings WorkSettings
}

// 	TCPSettings :
// 	Url -
type TCPSettings struct {
	Url          string
	WorkSettings WorkSettings
}

// 	WorkSettings :
// 	NumberOfWorkers -
// 	Retry -
type WorkSettings struct {
	NumberOfWorkers int
	Retry           int
}
