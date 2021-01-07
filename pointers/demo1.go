package pointers

import "fmt"

func demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı : ", *sayi)
}

func Test() {
	sayi := 10
	demo1(&sayi)
	//numara?? 10
}
