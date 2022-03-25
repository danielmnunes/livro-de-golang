package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	records, err := readData("top-1m.csv")

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	ch := make(chan string)

	for _, record := range records {
		url := fmt.Sprint("https://www.", record[1])
		go fetch(url, ch)
	}
	for range records {
		fmt.Println(<-ch)
	}
	fmt.Printf("%v elapsed\n", time.Since(start))

}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%s | %7d | %s", time.Since(start), nbytes, url)
}
