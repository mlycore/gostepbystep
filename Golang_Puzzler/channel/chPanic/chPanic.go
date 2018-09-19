package main

func main(){
	// example 1
	ch1 := make(chan int, 1)
	ch1 <- 1
	//ch1 <- 2 //here will have a panic, because the channel is already full.

	// example 2
	ch2 := make(chan int, 1)
	//elem,ok := <- ch2 //here will have a panic , because the channel is empty.
	//_,_ = elem, ok
	//ch2 <- 1

	// emample 3
	var ch3 chan int
	//ch3 <- 1  //here will have a panic, because the channel's value is nil.
	<- ch2  //here will have a panic, because the channel's value is nil.
	_ = ch3


}
