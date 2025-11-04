package main

import "fmt"

/*
Пример 1: Объявление sum внутри и вне цикла (для int)
Видимый эффект такой же, компилятор оптимизирует одинаково.
*/

// Вне цикла
func SumOutside(arr []int) int {
	sum := 0
	for i := 1; i < len(arr); i++ {
		sum = arr[i-1] + arr[i]
	}
	return sum
}

// Внутри цикла
func SumInside(arr []int) int {
	result := 0
	for i := 1; i < len(arr); i++ {
		sum := arr[i-1] + arr[i]
		result = sum // просто, чтобы использовать sum, иначе компилятор может оптимизировать сильнее
	}
	return result
}

/*
Пример 2: Разница для больших структур или слайсов (slice)
Здесь важно объявлять внешние переменные, чтобы не было лишних выделений памяти!
*/

// Медленно: внутри цикла — новый slice каждый раз
func SlowSlice() {
	for i := 0; i < 1000; i++ {
		data := make([]int, 10000) // 1000 выделений памяти!
		_ = data
	}
}

// Быстро: один slice, переиспользуем
func FastSlice() {
	data := make([]int, 10000)
	for i := 0; i < 1000; i++ {
		// можно занулять элементы, если нужно
		_ = data
	}
}

/*
Пример 3: Стиль кода и “принцип минимальной видимости”
*/

// Видимость шире, чем надо — хуже для чтения
func LessClear(arr []int) int {
	sum := 0 // sum видна во всей функции
	for i := 1; i < len(arr); i++ {
		sum = arr[i-1] + arr[i]
	}
	// кто-то мог бы случайно использовать sum здесь (непонятно зачем)
	return sum
}

// Лучшее решение — объявлять переменную максимально близко к месту использования!
func ClearScope(arr []int) int {
	result := 0
	for i := 1; i < len(arr); i++ {
		sum := arr[i-1] + arr[i] // sum живёт только в одной итерации
		result = sum
	}
	return result
}

// Мини-кейс для проверки работы (можно запустить через go run)
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(SumOutside(arr)) // 7
	fmt.Println(SumInside(arr))  // 7
	FastSlice()
	SlowSlice()
	fmt.Println(ClearScope(arr)) // 7
}
