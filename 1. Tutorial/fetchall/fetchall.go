package main

import (
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"bufio"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	ioutil.WriteFile("output.txt", []byte(""), 0600)
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("while opening %s: %v", f.Name(), err)
	}

	defer f.Close()
	
	f2, err := os.OpenFile("websites", 0, 0600)
	if err != nil {
		fmt.Printf("while opening %s: %v", f2.Name(), err)
	}

	defer f2.Close()

	scanner := bufio.NewScanner(f2)

	i := 0
	for scanner.Scan() {
		go fetch(scanner.Text(), ch)
		i++
	}

	for j := 0; j < i; j++ {
		_, err = f.WriteString(<-ch)
		if err != nil {
			fmt.Printf("while writing %s %v", f.Name(), err)
		}
	}
	_, err = f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	if err != nil {
		fmt.Printf("while writing %s %v", f.Name(), err)
	}
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	if !strings.HasPrefix("http://", url){
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}