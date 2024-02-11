package main

import "time"

// в исходниках написано, что time.Sleep
// блокирует горутину на время, значит
// используем небуферизированный канал, при чтении
// также блокирующий горутину
func sleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	sleep(2 * time.Second)
}
