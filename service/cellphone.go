package service

import (
	"go-crawl/types"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchCellPhoneData() ([]types.CellPhoneData, error) { // Use 'types.CellPhoneData' instead of 'Data'
	var allData []types.CellPhoneData
	doc, err := goquery.NewDocument("https://cellphones.com.vn/sforum/link-nhan-code-coin-master-spin/amp") // Replace with the actual URL
	if err != nil {
		return nil, err
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		date := s.Find("td").Eq(0).Text()
		if date == "Ngày" || strings.TrimSpace(date) == "" {
			return
		}
		gift := s.Find("td").Eq(1).Text()
		link, _ := s.Find("td a").Attr("href")

		// Parse the date from the string
		parsedDate, err := time.Parse("02/01/2006", date)
		if err != nil {
			log.Fatal(err)
		}

		// Get the current date
		now := time.Now()

		// Check if the date is within the last 4 days
		if now.Sub(parsedDate).Hours() <= 4*24 {
			allData = append(allData, types.CellPhoneData{Date: date, Gift: gift, Link: link})
		}
	})

	return allData, nil
}
