package main

import "fmt"

var justString string
//проблема такого
//1. в памяти остается вся строка v
//2. GC не сможет освободить ее, так как justString держит ссылку на v
//Итог: Утечка памяти - если вызывать ее много раз
// Чтобы исправить, лучше создать новый срез и копирать туда все данные
//justString - указывает на только первые 100 символов, но удерживает ВЕСЬ МАССИВ большой строки V === утечка памяти
func someFunc() {
	v := createHugeString(1 &lt;&lt; 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
// - конкретный пример реализации
//var justString string
//
//func createHugeString(size int) string {
//	return string(make([]byte, size))
//}
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = string([]byte(v[:100]))
//}
//
//func main() {
//	someFunc()
//	fmt.Println("Длина justString:", len(justString))
//}
