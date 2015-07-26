package main

import (
	"log"
	"fmt"
	"time"
	"strings"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty.usbmodem621", Baud: 9600, ReadTimeout: time.Second * 15}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
    }
    defer s.Close()

    for {
		reader := bufio.NewReader(s)
		reply, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		writeToCSV(reply)
	}
}

func writeToCSV(buf string) {
	
	csvfile, err := os.OpenFile("/Users/andrewlb/go/src/readywater/culturemonitor/culture.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("find culture.csv failed")
		panic(err)
	}
	data := fmt.Sprintf("%s,%s",time.Now().Local(),buf)

	// data = []string{data}
	fmt.Println("recording",data)
	_,err = io.WriteString(csvfile, data)
	if err != nil {
		fmt.Println("write to file failed")
		panic(err)
	}
}

func findArduino() string {
	contents, err := ioutil.ReadDir("/dev")
	if err != nil {
		fmt.Println("find arduino failed")
		panic(err)
	}

	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbmodem") {
			return "/dev/" + f.Name()
		}
	}

	fmt.Println("No arduino apparently")
	return ""
}
