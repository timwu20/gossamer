package offchain

import (
	"errors"
	"net/http"
	"sync"
)

const maxConcurrentRequests = 1000

var (
	errIntBufferEmpty        = errors.New("int buffer exhausted")
	errIntBufferFull         = errors.New("int buffer is full") //nolint:unused
	errRequestIDNotAvailable = errors.New("request id not available")
)

type requestIDBuffer chan int16

func newIntBuffer(buffSize int16) *requestIDBuffer {
	b := make(chan int16, buffSize)
	for i := int16(0); i < buffSize; i++ {
		b <- i
	}

	intb := requestIDBuffer(b)
	return &intb
}

func (b *requestIDBuffer) get() (int16, error) {
	select {
	case v := <-*b:
		return v, nil
	default:
		return 0, errIntBufferEmpty
	}
}

// nolint:unused
func (b *requestIDBuffer) put(i int16) error {
	select {
	case *b <- i:
		return nil
	default:
		return errIntBufferFull
	}
}

// HTTPSet holds a pool of concurrent http request calls
type HTTPSet struct {
	mtx    *sync.Mutex
	reqs   map[int16]*http.Request
	idBuff *requestIDBuffer
}

// NewHTTPSet creates a offchain http set that can be used
// by runtime as HTTP clients, the max concurrent requests is 1000
func NewHTTPSet() *HTTPSet {
	return &HTTPSet{
		mtx:    new(sync.Mutex),
		reqs:   make(map[int16]*http.Request),
		idBuff: newIntBuffer(maxConcurrentRequests),
	}
}

// StartRequest create a new request using the method and the uri, adds the request into the list
// and then return the position of the request inside the list
func (p *HTTPSet) StartRequest(method, uri string) (int16, error) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	id, err := p.idBuff.get()
	if err != nil {
		return 0, err
	}

	if _, ok := p.reqs[id]; ok {
		return 0, errRequestIDNotAvailable
	}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return 0, err
	}

	p.reqs[id] = req
	return id, nil
}

// Remove just remove a expecific request from reqs
func (p *HTTPSet) Remove(id int16) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	delete(p.reqs, id)
}

// Get returns a request or nil if request not found
func (p *HTTPSet) Get(id int16) *http.Request {
	return p.reqs[id]
}