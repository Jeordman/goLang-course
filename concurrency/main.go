// check websites are running concurrently - with at timer
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a string channel
	c := make(chan string)

	for _, link := range links {
		// run this function in its own go routine
		go checkLink(link, c)
	}

	// for loop
	// for i := 0; i < len(links); i++ {
	// infinite loop

	// for every link recieved from a channel, run this loop
	for l := range c {
		// the print line is blocked until a message comes in to relieve it
		// fmt.Println(<-c)
		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// there was an error hitting this website
		fmt.Println(link, "might be down!")
		// c <- "Failed"
		c <- link
		return
	}

	fmt.Println(link, "is up and working")
	// c <- "Success"
	c <- link
}
