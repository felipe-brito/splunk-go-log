package main

import (
	"fmt"
	"github.com/felipe-brito/splunk-go-log/pkg/hec"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"github.com/felipe-brito/splunk-go-log/pkg/splunk"
	"time"
)

type M struct {
	Msg string
}

func main() {
	//TcpConnect()
	//HecConnect()
	s := splunk.GetInstance(splunk.HEC)

	tcp := model.HECSettings{}
	tcp.Url = "http://localhost:8080"

	_ = s.Configuration(tcp)
	s.Initialize()

	time.Sleep(time.Second * 5)

	a := true
	count := 0

	for a {
		count++
		fmt.Println("Count: ", count)
		hec.Send(M{Msg: "Felipe de Brito Lira"})
		time.Sleep(time.Second * 5)

		if count == 12 {
			a = false
		}

	}

}
