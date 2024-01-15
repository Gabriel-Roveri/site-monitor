package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	for {
		fmt.Println("1 - Start monitoring\n2 - Show logs\n0 - Exit")
		var id int
		fmt.Scan(&id)

		switch id {
		case 1:
			startMonitoring()
			fmt.Println("Monitorando...")
		case 2:
			showLogs()

		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}

}

func startMonitoring() {

	site := fileReader()

	for i := 0; i < 3; i++ {
		for i, sites := range site {
			fmt.Println("Site ", i, ":", sites)
			siteTesting(sites)
		}
		time.Sleep(5 * time.Second) //timer
	}

}

func siteTesting(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("On air")
		logRegister(site, true)
	} else {
		fmt.Println("Non air")
		logRegister(site, false)
	}
}

func fileReader() []string {

	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file) //leitor

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	for err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("Monday 02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n") //conversão bool > string
	// https://go.dev/src/time/format.go
	file.Close()
}

func showLogs() {

	file, err := ioutil.ReadFile("log.txt")

	for err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
