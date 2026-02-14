package models

import "strconv"

type URLSubmissionQuota struct {
	DailyQuota    int `json:"DailyQuota"`
	MonthlyQuota  int `json:"MonthlyQuota"`
	DailyUsed     int `json:"DailyUsed,omitempty"`
	MonthlyUsed   int `json:"MonthlyUsed,omitempty"`
}

func (q *URLSubmissionQuota) Headers() []string {
	return []string{"Daily Quota", "Monthly Quota", "Daily Used", "Monthly Used"}
}

func (q *URLSubmissionQuota) Rows() [][]string {
	return [][]string{{
		strconv.Itoa(q.DailyQuota),
		strconv.Itoa(q.MonthlyQuota),
		strconv.Itoa(q.DailyUsed),
		strconv.Itoa(q.MonthlyUsed),
	}}
}
