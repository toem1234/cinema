package main

import (
	"fmt"
	"github.com/toem1234/cinema/movie"
	"github.com/toem1234/cinema/ticket"
)

// function ตัวใหญ่ นำหน้า public
// function ตัวเล็ก นำหน้า private
// init by order
// not allow import cycle
// how to publish to other

// how to create gomod
// cmd => go mod init "name package"
// example go mod init "github.com/toem1234/cinema

func init() {
	// is convention like constructor
	fmt.Println("init : main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
