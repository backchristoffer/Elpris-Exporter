package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	//"github.com/gin-gonic/gin"
)

// how elpris is structured -
type ElprisRegion struct {
	Year        int
	Month       int
	Day         int
	PriceRegion string
}

func main() {

	currentTime := time.Now()

	elpris := ElprisRegion{
		Year:        currentTime.Year(),
		Month:       int(currentTime.Month()),
		Day:         currentTime.Day(),
		PriceRegion: "SE2",
	}

	url := fmt.Sprintf("https://www.elprisetjustnu.se/api/v1/prices/%d/%02d-%02d_%s.json", elpris.Year, elpris.Month, elpris.Day, elpris.PriceRegion)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Print(sb)
}
