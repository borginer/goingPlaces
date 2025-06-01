package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type xkcdData struct {
	Num        int
	Transcript string
	URL        string
}

func (d *xkcdData) String() string {
	var s []string
	s = []string{string(d.Num), d.Transcript, d.URL}
	return strings.Join(s, "@")

}

const URLFMT = `https://xkcd.com/%d/info.0.json`

func GetData(url string) *xkcdData {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xkcd: %v", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "xkcd: bad status")
		return nil
	}
	var msg xkcdData

	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "xkcd: %v", err)
	}
	resp.Body.Close()
	return &msg
}

func main() {
	for i := 500; i < 550; i++ {
		url := fmt.Sprintf(URLFMT, i)
		data := GetData(url)
		if data != nil {
			os.WriteFile("Data/"+strconv.Itoa(i), []byte(data.String()), 0777)
		}
	}
}
