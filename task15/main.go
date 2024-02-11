package main

// тк это глобальная переменная пакета,
// то сборщик мусора ее не удалит.
var justString string

func createHugeString(l int64) string {
	return string(make([]byte, l))
}

/*
выделенная лишняя память под массив не соберется сборщиком мусора,
тк переменная глобальная, мы будем держать эту ссылку до первого измнения
переменной
func someFunc() {
	v := createHugeString(1 << 10)
	// емкость среза равна его длине,
	// но мы по-прежнему ссылаемся на большой базовый массив
	justString = v[:100]
}
*/

func someFunc() {
	v := createHugeString(1 << 10)

	justByte := make([]byte, 100)
	copy(justByte, v) // эта функция создаст новый массив
	justString = string(justByte)
}

func main() {
	someFunc()
}
