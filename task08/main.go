package main

import (
	"fmt"
	"strconv"
)

func bitNum(num int64, position uint8, one bool) int64 {
	// на этот тип выделяется не больше 64 бит,
	// если залезть дальше, то будет ошибка
	if position > 64 {
		fmt.Println("Number is too small for this")
		return 0
	}

	fmt.Printf("num: %d \n", num)
	fmt.Printf("bits: %s \n", strconv.FormatInt(num, 2))

	mask := int64(1)
	if one {
		// делаем маску, побитово сдвигая единицу на нужную позицию аля 0001000
		mask = mask << position
		// выполняем операцию or, тк она всегда вернет 1 в разряд, если оба операнда не ноль
		// (у нас на нужной позиции единица, значит 1 в бит поставится)
		num = num | mask
	} else {
		// делаем инвертированную маску, тоже самое,
		// только потом применяем к битам not аля 1110111
		mask = ^(mask << position)
		// выполняем операцию and, тк она всегда вернет 0 в разряд,
		// если только оба операнда не единица (у нас на нужной позиции 0, значит он поставится)
		num = num & mask
	}

	fmt.Printf("result: %d \n", num)
	fmt.Printf("resBits: %s \n", strconv.FormatInt(num, 2))
	return num
}

func main() {
	num := int64(5)
	_ = bitNum(num, 62, true)
}
