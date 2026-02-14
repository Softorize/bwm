package models

import (
	"fmt"
	"strconv"
)

type TrafficStats struct {
	Date        BingTime `json:"Date"`
	Clicks      int64    `json:"Clicks"`
	Impressions int64    `json:"Impressions"`
	AvgCTR      float64  `json:"AvgCTR,omitempty"`
	AvgPosition float64  `json:"AvgImpressionPosition,omitempty"`
}

type TrafficStatsList []TrafficStats

func (t TrafficStatsList) Headers() []string {
	return []string{"Date", "Clicks", "Impressions", "CTR", "Avg Position"}
}

func (t TrafficStatsList) Rows() [][]string {
	rows := make([][]string, len(t))
	for i, s := range t {
		rows[i] = []string{
			s.Date.String(),
			strconv.FormatInt(s.Clicks, 10),
			strconv.FormatInt(s.Impressions, 10),
			fmt.Sprintf("%.2f%%", s.AvgCTR*100),
			fmt.Sprintf("%.1f", s.AvgPosition),
		}
	}
	return rows
}

type QueryStats struct {
	Query       string   `json:"Query"`
	Date        BingTime `json:"Date"`
	Clicks      int64    `json:"Clicks"`
	Impressions int64    `json:"Impressions"`
	AvgCTR      float64  `json:"AvgCTR,omitempty"`
	AvgPosition float64  `json:"AvgImpressionPosition,omitempty"`
}

type QueryStatsList []QueryStats

func (q QueryStatsList) Headers() []string {
	return []string{"Query", "Date", "Clicks", "Impressions", "CTR", "Avg Position"}
}

func (q QueryStatsList) Rows() [][]string {
	rows := make([][]string, len(q))
	for i, s := range q {
		rows[i] = []string{
			s.Query,
			s.Date.String(),
			strconv.FormatInt(s.Clicks, 10),
			strconv.FormatInt(s.Impressions, 10),
			fmt.Sprintf("%.2f%%", s.AvgCTR*100),
			fmt.Sprintf("%.1f", s.AvgPosition),
		}
	}
	return rows
}

type PageStats struct {
	Query       string   `json:"Query"`
	Page        string   `json:"Page,omitempty"`
	Date        BingTime `json:"Date"`
	Clicks      int64    `json:"Clicks"`
	Impressions int64    `json:"Impressions"`
}

type PageStatsList []PageStats

func (p PageStatsList) Headers() []string {
	return []string{"Page", "Query", "Date", "Clicks", "Impressions"}
}

func (p PageStatsList) Rows() [][]string {
	rows := make([][]string, len(p))
	for i, s := range p {
		rows[i] = []string{
			s.Page,
			s.Query,
			s.Date.String(),
			strconv.FormatInt(s.Clicks, 10),
			strconv.FormatInt(s.Impressions, 10),
		}
	}
	return rows
}
