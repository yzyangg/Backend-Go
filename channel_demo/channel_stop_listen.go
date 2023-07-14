package main

//func main() {
//	var (
//		keep = make(chan int)
//		stop = make(chan int)
//	)
//	go func() {
//		time.Sleep(3 * time.Second)
//		stop <- 1
//	}()
//	go func() {
//		for {
//			time.Sleep(1 * time.Second)
//			keep <- 1
//		}
//	}()
//loop:
//	for {
//		select {
//		case <-keep:
//			print("keep\n")
//
//		case <-stop:
//			print("stop")
//			break loop
//		}
//	}
//
//}
