package main

import (
	"context"
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

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Printf("Процесс %v отменен\n", num)
		return
	case <-timer.C:
		fmt.Printf("Данные процесса %v успешно отправлены\n", num)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := range 10 {
		go sendData(ctx, i)
	}
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}
