package models

import (
	"fmt"
	"strconv"
)

type CrawlStats struct {
	Date             BingTime `json:"Date"`
	CrawledPages     int64    `json:"CrawledPages"`
	CrawlErrors      int64    `json:"CrawlErrors"`
	InIndex          int64    `json:"InIndex"`
	InLinks          int64    `json:"InLinks"`
	BlockedByRobots  int64    `json:"BlockedByRobots"`
	ContainsMalware  int64    `json:"ContainsMalware"`
	CrawlBandwidth   int64    `json:"CrawlBandwidth"`
}

type CrawlStatsList []CrawlStats

func (c CrawlStatsList) Headers() []string {
	return []string{"Date", "Crawled", "Errors", "In Index", "In Links", "Blocked", "Malware"}
}

func (c CrawlStatsList) Rows() [][]string {
	rows := make([][]string, len(c))
	for i, s := range c {
		rows[i] = []string{
			s.Date.String(),
			strconv.FormatInt(s.CrawledPages, 10),
			strconv.FormatInt(s.CrawlErrors, 10),
			strconv.FormatInt(s.InIndex, 10),
			strconv.FormatInt(s.InLinks, 10),
			strconv.FormatInt(s.BlockedByRobots, 10),
			strconv.FormatInt(s.ContainsMalware, 10),
		}
	}
	return rows
}

type CrawlIssue struct {
	Issue    string `json:"Issue"`
	Count    int64  `json:"Count"`
	Severity int    `json:"Severity"`
}

type CrawlIssueList []CrawlIssue

func (c CrawlIssueList) Headers() []string {
	return []string{"Issue", "Count", "Severity"}
}

func (c CrawlIssueList) Rows() [][]string {
	rows := make([][]string, len(c))
	for i, issue := range c {
		rows[i] = []string{
			issue.Issue,
			strconv.FormatInt(issue.Count, 10),
			strconv.Itoa(issue.Severity),
		}
	}
	return rows
}

type CrawlSettings struct {
	CrawlBoostAvailable bool  `json:"CrawlBoostAvailable"`
	CrawlBoostEnabled   bool  `json:"CrawlBoostEnabled"`
	CrawlRate           int   `json:"CrawlRate"`
}

func (s *CrawlSettings) Headers() []string {
	return []string{"Crawl Rate", "Boost Available", "Boost Enabled"}
}

func (s *CrawlSettings) Rows() [][]string {
	return [][]string{{
		strconv.Itoa(s.CrawlRate),
		fmt.Sprintf("%v", s.CrawlBoostAvailable),
		fmt.Sprintf("%v", s.CrawlBoostEnabled),
	}}
}
