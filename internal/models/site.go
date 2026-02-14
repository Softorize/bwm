package models

import "strconv"

type Site struct {
	URL                string `json:"Url"`
	IsVerified         bool   `json:"IsVerified"`
	AuthenticationCode string `json:"AuthenticationCode"`
}

type SiteList []Site

func (s SiteList) Headers() []string {
	return []string{"URL", "Verified"}
}

func (s SiteList) Rows() [][]string {
	rows := make([][]string, len(s))
	for i, site := range s {
		rows[i] = []string{site.URL, strconv.FormatBool(site.IsVerified)}
	}
	return rows
}
