package main

import (
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/splunk"
)

func main() {
	//TcpConnect()
	//HecConnect()
	s := splunk.GetInstance(splunk.HEC)

	tcp := model.HecSettings{}
	tcp.Url = "http://localhost:8080"

	_ = s.Configuration(tcp)
	s.Initialize()
}
