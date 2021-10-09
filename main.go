package main

import "fmt"

func main() {
	limit := 5
	count := 0
	var digit1 int
	var digit2 int
	var digit3 int
	fmt.Println("Три числа.")

	fmt.Println("Введите первое число:")
	fmt.Scan(&digit1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&digit2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&digit3)

	if digit1 > limit {
		count++
	}
  if digit2 > limit {
		count++
	}
  if digit3 > limit {
		count++
	} 
	
	fmt.Println("Среди введённых чисел", count, "больше или равны", limit, ".")
}
