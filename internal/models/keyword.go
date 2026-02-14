package models

import (
	"fmt"
	"strconv"
)

type KeywordStats struct {
	Query       string   `json:"Query"`
	Date        BingTime `json:"Date"`
	Impressions int64    `json:"Impressions"`
	Clicks      int64    `json:"Clicks"`
	Position    float64  `json:"AvgImpressionPosition,omitempty"`
}

type KeywordStatsList []KeywordStats

func (k KeywordStatsList) Headers() []string {
	return []string{"Query", "Date", "Impressions", "Clicks", "Avg Position"}
}

func (k KeywordStatsList) Rows() [][]string {
	rows := make([][]string, len(k))
	for i, s := range k {
		rows[i] = []string{
			s.Query,
			s.Date.String(),
			strconv.FormatInt(s.Impressions, 10),
			strconv.FormatInt(s.Clicks, 10),
			fmt.Sprintf("%.1f", s.Position),
		}
	}
	return rows
}

type RelatedKeyword struct {
	Query       string  `json:"Query"`
	Impressions int64   `json:"Impressions"`
	BroadCTR    float64 `json:"BroadClickThrough,omitempty"`
}

type RelatedKeywordList []RelatedKeyword

func (r RelatedKeywordList) Headers() []string {
	return []string{"Query", "Impressions", "Broad CTR"}
}

func (r RelatedKeywordList) Rows() [][]string {
	rows := make([][]string, len(r))
	for i, k := range r {
		rows[i] = []string{
			k.Query,
			strconv.FormatInt(k.Impressions, 10),
			fmt.Sprintf("%.2f%%", k.BroadCTR*100),
		}
	}
	return rows
}
