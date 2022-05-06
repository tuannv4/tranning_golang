package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		+ Goroutines:
			- Là function xử lý đồng thời với các goroutine khác
			- Có thể coi là 1 thread siêu gọn nhẹ, chỉ tốn 2KB bộ nhớ stack, có thể tăng/giảm vùng nhớ tùy theo nhu cầu sử dụng
			- Một chương trình Go có thể có hàng ngàn Goroutine được xử lý đồng thời

		+ Channel:
			- coi là các đường ống sử dụng mà Goroutines giao tiếp
			- dữ liệu có thể được gửi từ một đầu và nhận từ đầu kia bằng channels
			- Mỗi channel có một loại liên kết với nó. Loại này là loại dữ liệu mà channel được phép vận chuyển. Không có loại khác được phép vận chuyển bằng cách sử dụng channel
			- chan T là một channel loại T
			- Giá trị 0 của một channel là không
			- Channels nil không được sử dụng và do đó channel phải được xác định bằng cách sử dụng make tương tự như map và slice
	*/
	// sử dụng time.Sleep cho hàm main chờ để chạy 2 func này
	// go PrintChar()
	// go PrintNumber()
	// time.Sleep(3 * time.Second)

	// sử dụng WaitGroup
	/*var wg sync.WaitGroup
			wg.Add(2)
			go PrintChar(&wg)
			go PrintNumber(&wg)
			wg.Wait()
			log.Println("==============================main finish==============================")
		}

		func PrintNumber(wg *sync.WaitGroup) {
			for i := 0; i <= 100; i++ {
				result += i
			}
			wg.Done()
		}
		func PrintChar(wg *sync.WaitGroup) {
			for i := 'A'; i < 'A'+26; i++ {
				result += string(i)
			}
			wg.Done()
		}

		// channel
		chanNumber := make(chan int)
		chanChar := make(chan string)
		go PrintChar(chanChar)
		go PrintNumber(chanNumber)
		log.Println("kết quả từ printNumber là: ", <-chanNumber)
		log.Println("kết quả từ printChar là: ", <-chanChar)
		log.Println("==============================main finish==============================")
	}

		func PrintNumber(numChan chan int) {
			var result int
			for i := 0; i <= 100; i++ {
				result += i
			}
			numChan <- result
		}
		func PrintChar(strChan chan string) {
			var result string
			for i := 'A'; i < 'A'+26; i++ {
				result += string(i)
			}
			strChan <- result
		}
		// deadlock
		- thiếu gửi hoặc nhận sẽ bị deadlock
			chanInt := make(chan int)
			chanInt <- 5
		- sửa lại cho đúng

		func dumy(dumyChan chan int){
			dumyChan <- 10
		}
		chanInt := make(chan int)
		go dummy(chanInt)
		fmt.Println("nhận dữ liệu", <- chanInt)

		- đọc data từ URL

		fisrtChanUrl := make(chan string)
		secondChanUrl := make(chan string)
		thirdChanUrl := make(chan string)

		listChanUrl := []chan string{fisrtChanUrl, secondChanUrl, thirdChanUrl}
		listUrl := []string{"https://youtube.com", "https://google.com", "https://vnexpress.net"}

		for i := 0; i < len(listUrl); i++ {
			go getUrl(listUrl[i], listChanUrl[i])
		}
		for i := 0; i < len(listUrl); i++ {
			fmt.Println(<-listChanUrl[i])
			fmt.Println("================================================================")
		}
		func getUrl(url string, result chan string) {
			resp, errGetUrl := http.Get(url)
			if errGetUrl != nil {
				log.Println(errGetUrl)
				return
			}
			urlContent, errReadBody := ioutil.ReadAll(resp.Body)
			if errReadBody != nil {
				log.Println(errReadBody)
				return
			}
			result <- string(urlContent)
			defer resp.Body.Close()
		}

	- function return channel
		func printfSth(msg string) chan string {
			result := make(chan string)
			go func() {
				for i := 0; i <= 5; i++ {
					result <- fmt.Sprintf("%s %d", msg, i)
				}
			}()
			return result
		}

		bridge := printfSth("hello")
		for i := 0; i <= 5; i++ {
			log.Println("receive from", <-bridge)
		}
		log.Println("main finished")
	- Fan-in pattern
		coffee := printfSth("order coffee")
		braed := printfSth("order bread")
		serve := fanIn(coffee, braed)
		for i := 0; i < 5; i++ {
			log.Println("receive from", <-serve)
		}
	- Pipeline

		firstChan := firstInput(1, 2, 3, 4, 5, 6)
		secondChan := secondInput(firstChan)

		for item := range secondChan {
			log.Println("receive", item)
		}
		log.Println("main finished")
	*/
	chanGg := make(chan string)
	chanBing := make(chan string)

	go googleSearch(chanGg)
	go bingSearch(chanBing)

	select {
	case result := <-chanGg:
		log.Println(result)
	case result := <-chanBing:
		log.Println(result)
	}
	fmt.Print("xxx")

	// unbuffered channel
	ch := make(chan int)
	go func() {
		ch <- 100
		fmt.Print("sent")
	}()
	fmt.Println(<-ch)
	fmt.Println("done")

	// buffered channel
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

func printfSth(msg string) chan string {
	result := make(chan string)
	go func() {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			// select channel
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}

func firstInput(num ...int) chan int {
	result := make(chan int)
	go func() {
		for i := 0; i < len(num); i++ {
			result <- num[i]
		}
		close(result)
	}()
	return result
}

func secondInput(formFisrt chan int) chan int {
	result := make(chan int)
	go func() {
		for item := range formFisrt {
			result <- item * item
		}
		close(result)
	}()
	return result
}

func googleSearch(result chan string) {
	time.Sleep(1 * time.Second)
	result <- "found from Bing"
}

func bingSearch(result chan string) {
	time.Sleep(1 * time.Second)
	result <- "foung from Bing"
}
