package pointers

import (
	"fmt"
)

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayı : ", sayilar[0])
}

func test() {
	numaralar := []int{1, 2, 3, 4, 5}
	Demo2(numaralar)
}
