package models

import (
	"strconv"
)

// Sitemap is one Bing "Feed" as returned by GetFeeds / GetFeedDetails.
type Sitemap struct {
	URL         string   `json:"Url"`
	Type        string   `json:"Type"`   // "Sitemap", "Sitemap Index", "RSS", ...
	Status      string   `json:"Status"` // "Success", "Pending", ...
	Compressed  bool     `json:"Compressed"`
	FileSize    int64    `json:"FileSize"`
	LastCrawled BingTime `json:"LastCrawled"`
	Submitted   BingTime `json:"Submitted"`
	URLCount    int64    `json:"UrlCount"`
}

type SitemapList []Sitemap

func (s SitemapList) Headers() []string {
	return []string{"URL", "Type", "Status", "Last Crawled", "Submitted", "URLs"}
}

func (s SitemapList) Rows() [][]string {
	rows := make([][]string, len(s))
	for i, sm := range s {
		rows[i] = []string{
			sm.URL,
			sm.Type,
			sm.Status,
			bingDate(sm.LastCrawled),
			bingDate(sm.Submitted),
			strconv.FormatInt(sm.URLCount, 10),
		}
	}
	return rows
}

// Bing reports an unset date as /Date(-11644473600000)/ (year 1601).
func bingDate(t BingTime) string {
	if t.Year() < 1970 {
		return "-"
	}
	return t.String()
}
