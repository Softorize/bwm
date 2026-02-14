package models

import (
	"strconv"
)

type Sitemap struct {
	URL             string   `json:"Url"`
	IsPending       bool     `json:"IsPending"`
	LastCrawledDate BingTime `json:"LastCrawledDate"`
	LastSubmitted    BingTime `json:"LastSubmittedDate"`
	URLCount        int64    `json:"UrlCount"`
	Warnings        int64    `json:"Warnings"`
	Errors          int64    `json:"Errors"`
}

type SitemapList []Sitemap

func (s SitemapList) Headers() []string {
	return []string{"URL", "Pending", "Last Crawled", "URLs", "Warnings", "Errors"}
}

func (s SitemapList) Rows() [][]string {
	rows := make([][]string, len(s))
	for i, sm := range s {
		rows[i] = []string{
			sm.URL,
			strconv.FormatBool(sm.IsPending),
			sm.LastCrawledDate.String(),
			strconv.FormatInt(sm.URLCount, 10),
			strconv.FormatInt(sm.Warnings, 10),
			strconv.FormatInt(sm.Errors, 10),
		}
	}
	return rows
}

type SitemapDetail struct {
	URL             string   `json:"Url"`
	IsPending       bool     `json:"IsPending"`
	IsSitemapIndex  bool     `json:"IsSitemapIndex"`
	LastCrawledDate BingTime `json:"LastCrawledDate"`
	LastSubmitted    BingTime `json:"LastSubmittedDate"`
	URLCount        int64    `json:"UrlCount"`
	Warnings        int64    `json:"Warnings"`
	Errors          int64    `json:"Errors"`
	TotalURLs       int64    `json:"TotalUrls"`
}

func (d *SitemapDetail) Headers() []string {
	return []string{"URL", "Pending", "Index", "Last Crawled", "Total URLs", "Warnings", "Errors"}
}

func (d *SitemapDetail) Rows() [][]string {
	return [][]string{{
		d.URL,
		strconv.FormatBool(d.IsPending),
		strconv.FormatBool(d.IsSitemapIndex),
		d.LastCrawledDate.String(),
		strconv.FormatInt(d.TotalURLs, 10),
		strconv.FormatInt(d.Warnings, 10),
		strconv.FormatInt(d.Errors, 10),
	}}
}
