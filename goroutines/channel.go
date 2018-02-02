package main

func main() {
	ch := make(chan int, 4)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()
	go func() {
		ch <- 3
	}()
	go func() {
		ch <- 4
	}()
	for i := 0; i < 4; i++ {
		<-ch
	}

}
