package main

import "fmt"

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Произошла паника: %v \n", r)
		}
	}()

	panic("Error")

	return nil
}

func main() {

	if err := foo(); err != nil {
		fmt.Println("Ошибка при выполнении программы:", err)
	} else {
		fmt.Println("Успешно")
	}
}
