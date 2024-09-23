package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция которая имитирует работу
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker #%v закончил работу\n", id)
}

// Функция которая ожидает завершения работы
func runWorkers(count int) error {
	wg := sync.WaitGroup{}
	timer := time.After(2 * time.Second)

	numWorkers := count
	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i, &wg)
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Успешно")
		return nil
	case <-timer:
		return fmt.Errorf("Не успела выполнится за 2 секунды")
	}
}

func main() {
	err := runWorkers(200000)

	if err != nil {
		fmt.Printf("При выполнении возникла ошибка: %v", err)
	}

}
