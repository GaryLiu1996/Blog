package codes

import "fmt"

func ex() {
	ch := make(chan int)
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 1)
	//ch <- x
	//x := <- ch
	//<- ch
	close(ch)
	fmt.Println("chan: ", ch, ch1, ch2)
}
