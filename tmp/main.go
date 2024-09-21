package main

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
	"tmp/client"
)

type ServerError struct {
	err error
}

func (s *ServerError) Error() string {
	return s.err.Error()
}

func NewServerError(msg string) error {
	return &ServerError{err: errors.New(msg)}
}

var internalErr = NewServerError("внутренняя ошибка")

func GetClient() (client.Client, error) {
	cl := client.Client{}
	err := fmt.Errorf("ошибка получения клиента: %w", internalErr)
	return cl, err
}

func addAge(cl *client.Client, add int) {
	atomic.AddInt64(&cl.Age, int64(add))
}
func main() {
	cl := &client.Client{Name: "string", Age: 1}
	//mu := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go addAge(cl, 1)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(cl.Age)
}
