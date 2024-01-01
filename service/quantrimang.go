package service

import (
	"go-crawl/types"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchQuanTriMangData() ([]types.DataQuanTriMangData, error) {
	var allData []types.DataQuanTriMangData
	doc, err := goquery.NewDocument("https://quantrimang.com/cong-nghe/code-coin-master-200712")
	if err != nil {
		return nil, err
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		date := s.Find("td").Eq(0).Text()
		if date == "Ngày" || strings.TrimSpace(date) == "" {
			return
		}
		spins := s.Find("td").Eq(1).Text()
		link, _ := s.Find("td a").Attr("href")

		link = "https://quantrimang.com" + link

		// Parse the date from the string
		parsedDate, err := time.Parse("02/01/2006", date)
		if err != nil {
			log.Fatal(err)
		}

		// Get the current date
		now := time.Now()

		// Check if the date is within the last 4 days
		if now.Sub(parsedDate).Hours() <= 4*24 {
			allData = append(allData, types.DataQuanTriMangData{Date: date, Spins: spins, Link: link})
		}
	})

	return allData, nil
}
