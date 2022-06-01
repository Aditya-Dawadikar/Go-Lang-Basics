package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func fetchData1(wg *sync.WaitGroup) {

	if wg != nil {
		defer wg.Done()
	}

	fmt.Println("API call 1")

	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println("some error occured")
		os.Exit(1)
	}

	var _, err1 = ioutil.ReadAll(res.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("API call 1 completed")
}

func fetchData2(wg *sync.WaitGroup) {

	if wg != nil {
		defer wg.Done()
	}
	fmt.Println("API call 2")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("some error occured")
		os.Exit(1)
	}

	var _, err1 = ioutil.ReadAll(res.Body)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("API call 2 completed")
}

func fetchData3(wg *sync.WaitGroup) {

	if wg != nil {
		defer wg.Done()
	}
	fmt.Println("API call 3")

	res, err := http.Get("https://catfact.ninja/fact")

	if err != nil {
		fmt.Println("some error occured")
		os.Exit(1)
	}

	var _, err1 = ioutil.ReadAll(res.Body)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("API call 3 completed")
}

func nonConcurentFunc() {
	fmt.Println("non concurent API calls")
	start := time.Now().UnixNano()

	fetchData1(nil)
	fetchData2(nil)
	fetchData3(nil)

	end := time.Now().UnixNano()

	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
	fmt.Println("duration: ", (end - start))
}

func concurentFunc(wg *sync.WaitGroup) {

	wg.Add(3)

	fmt.Println("concurent API calls")
	start := time.Now().UnixNano()

	go fetchData1(wg)
	go fetchData2(wg)
	go fetchData3(wg)

	end := time.Now().UnixNano()

	wg.Wait()

	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
	fmt.Println("duration: ", (end - start))
}

func main() {
	wg := new(sync.WaitGroup)

	nonConcurentFunc()
	fmt.Println()
	concurentFunc(wg)

}
