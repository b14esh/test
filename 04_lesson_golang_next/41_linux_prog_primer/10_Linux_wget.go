//wget https://chto_scachati/example.txt example.txt

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func wget(url string, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf(
			"status code: %d [%s]",
			resp.StatusCode, resp.Status,
		)

		panic(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}

func main() {

	wget(
		"https://chto_scachati/example.txt", // что скачать
		"example.txt",                       // как сохронить
	)
}
