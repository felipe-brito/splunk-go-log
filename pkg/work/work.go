package work

import "fmt"

const (
	NumberOfWorkersDefault = 10
	MinimumOfWorkers       = 1
)

type Work struct {
	workersChannel  chan interface{}
	numberOfWorkers int
}

func New(numberOfWorkers int) *Work {
	work := &Work{}
	work.setNumberOfWorkers(numberOfWorkers)
	work.initWorkersChannel()
	return work
}

func (w *Work) initWorkersChannel() {
	w.workersChannel = make(chan interface{}, w.getNumberOfWorkers())
}

func (w *Work) setNumberOfWorkers(numberOfWorkers int) {
	if numberOfWorkers > MinimumOfWorkers {
		w.numberOfWorkers = numberOfWorkers
	}
	w.numberOfWorkers = NumberOfWorkersDefault
}

func (w *Work) getNumberOfWorkers() int {
	return w.numberOfWorkers
}

func (w *Work) SendToWork(obj interface{}) {
	w.workersChannel <- obj
}

func (w *Work) ChannelIsClosed() bool {
	return w.workersChannel == nil
}

func (w *Work) CloseWorkers() {
	if w.workersChannel != nil {
		close(w.workersChannel)
	}
}

func (w *Work) StartWorkers() {
	for i := 0; i < w.getNumberOfWorkers(); i++ {
		go func() {
			for msg := range w.workersChannel {
				fmt.Println(msg)
			}
		}()
	}
}
