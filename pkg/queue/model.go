package queue

import (
	"fmt"
	"github.com/felipe-brito/splunk-go-log/pkg/model"
	"sync"
)

type Queue struct {
	bufferSize         int
	MessageBloc        []*Bloc
	currentBloc        *Bloc
	mutexCurrentBuffer sync.Mutex
	mutexBuffers       sync.Mutex
}

type Bloc struct {
	messages []model.Message
	size     int
	retry    int
}

// função new para queue
// função para add uma nova mensagem

func (q *Queue) Push(message interface{}) {

	_, success := message.(model.Message)
	if !success {
		fmt.Println("não foi possível realizar o cast para a estrutura message")
	}

	//q.currentBloc

	// add

	// verifica se o buffer corrente chegou ao limete

	// caso sim, um novo buffer é criado e o corrente é colocado na fila de envio

}

func (b *Bloc) add(message model.Message) {
	b.messages = append(b.messages, message)
}

func (b *Bloc) isEmpty() bool {
	return len(b.messages) < 1
}
