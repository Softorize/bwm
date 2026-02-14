package models

import "strconv"

type LinkCounts struct {
	InboundLinkCount int64 `json:"InboundLinkCount"`
	IndexedPages     int64 `json:"IndexedPages,omitempty"`
}

func (l *LinkCounts) Headers() []string {
	return []string{"Inbound Links", "Indexed Pages"}
}

func (l *LinkCounts) Rows() [][]string {
	return [][]string{{
		strconv.FormatInt(l.InboundLinkCount, 10),
		strconv.FormatInt(l.IndexedPages, 10),
	}}
}

type LinkDetail struct {
	SourceURL string `json:"SourceUrl"`
	TargetURL string `json:"TargetUrl"`
	AnchorText string `json:"AnchorText"`
}

type LinkDetailList []LinkDetail

func (l LinkDetailList) Headers() []string {
	return []string{"Source URL", "Target URL", "Anchor Text"}
}

func (l LinkDetailList) Rows() [][]string {
	rows := make([][]string, len(l))
	for i, link := range l {
		rows[i] = []string{link.SourceURL, link.TargetURL, link.AnchorText}
	}
	return rows
}
