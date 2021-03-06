package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %sn", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(rand.Intn(10)) + "n"
		c.Write([]byte(string(result)))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide a port number!")
			return
	}
	TestRing()

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
			c, err := l.Accept()
			if err != nil {
					fmt.Println(err)
					return
			}
			go handleConnection(c)
	}
}
