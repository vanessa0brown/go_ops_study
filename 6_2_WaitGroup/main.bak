package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func startGoroutines() error {
	wg := sync.WaitGroup{}
	timer := time.After(2 * time.Second)
	err := errors.New("Прошло больше 2ух секунд")

	f := func(w *sync.WaitGroup) {
		// Какая-то работа
		time.Sleep(1 * time.Second)
		w.Done()
	}

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go f(&wg)
	}

	wg.Wait()

	select {
	case <-timer:
		return err
	default:
		fmt.Println("Успешно")
		return nil
	}
}

func main() {
	err := startGoroutines()

	if err != nil {
		fmt.Printf("Программа завершилась с ошибкой: %v", err)
	}
}
