package main

import "fmt"

func main(){
	ch1 := make(chan int, 10)
	go func() {
		for i := 0;i<10;i ++ {
			fmt.Printf("Sender : send element %v\n",i)
			ch1 <- i

		}
		fmt.Println("Sender: close channel.")
		close(ch1)
	}()

	for{
		elem, ok := <- ch1
		if !ok{
			fmt.Println("Receiver: close channel.")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n",elem)
	}
	fmt.Println("end")
}
