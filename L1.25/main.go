package main

import (
	"fmt"
	"time"
)

// mySleep приостанавливает выполнение текущей горутины на длительность d
func mySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало", time.Now().Format("15:04:05.000"))
	mySleep(3 * time.Second)
	fmt.Println("Конец", time.Now().Format("15:04:05.000"))
}
