package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Event struct {
	App        string `json:"app"`
	AppVersion string `json:"appVersion"`
	Status     string `json:"status"`
}

func TcpConnect() {

	s := Event{
		App:        "test",
		AppVersion: "test",
		Status:     "test",
	}

	m := model.Message{
		Time:       time.Now(),
		Host:       "test",
		Source:     "test-t",
		SourceType: "_json",
		Index:      "hec_index",
		Event:      s,
	}

	jsons, _ := json.Marshal(&m)

	conn, err := net.DialTimeout("tcp", "localhost:8086", time.Duration(30)*time.Second)
	defer closeConnection(conn)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = conn.Write(jsons)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func closeConnection(conn net.Conn) {
	if conn != nil {
		_ = conn.Close()
	}
}

func HecConnect() {
	s := Event{
		App:        "test",
		AppVersion: "test",
		Status:     "test",
	}

	m := model.Message{
		Time:       time.Now(),
		Host:       "test",
		Source:     "test-t",
		SourceType: "_json",
		Index:      "hec_index",
		Event:      s,
	}

	jsons, _ := json.Marshal(&m)

	b := new(bytes.Buffer)
	b.Write(jsons)

	client := &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest("POST", "http://localhost:8088/services/collector", b)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Splunk 15c45bc7-1f90-404a-a0b3-a23c9e4874e8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		io.Copy(ioutil.Discard, res.Body)
	default:
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		fmt.Println(responseBody)
		err = errors.New(responseBody)
	}

}
