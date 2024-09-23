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
	err := foo()

	if err != nil {
		fmt.Println("Ошибка при выполнении программы")
	} else {
		fmt.Println("Успешно")
	}
}
