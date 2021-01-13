package model

import (
	"encoding/json"
	"github.com/felipe-brito/splunk-go-log/pkg/util"
	"time"
)

// 	Message :
// 	Time - time of upload to the server
// 	Host
// 	Source
// 	SourceType
// 	Index
// 	Event
type Message struct {
	Time       time.Time   `json:"time"`
	Host       string      `json:"host"`
	Source     string      `json:"source,omitempty"`
	SourceType string      `json:"sourcetype,omitempty"`
	Index      string      `json:"index,omitempty"`
	Event      interface{} `json:"event"`
}

func (m *Message) SetTime() {
	m.Time = time.Now()
}

func (m *Message) SetHost() {
	m.Host = util.HostName()
}

func (m *Message) GetBytes() []byte {
	b, _ := json.Marshal(m)
	return b
}
