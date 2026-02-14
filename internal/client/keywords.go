package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetKeywordStats(siteURL string) (models.KeywordStatsList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var stats models.KeywordStatsList
	err := c.get("GetQueryStats", params, &stats)
	return stats, err
}

func (c *Client) GetKeywordResearch(query string) (models.RelatedKeywordList, error) {
	params := url.Values{"query": {query}}
	var keywords models.RelatedKeywordList
	err := c.get("GetKeywordResearch", params, &keywords)
	return keywords, err
}

func (c *Client) GetRelatedKeywords(query string) (models.RelatedKeywordList, error) {
	params := url.Values{"query": {query}}
	var keywords models.RelatedKeywordList
	err := c.get("GetRelatedKeywords", params, &keywords)
	return keywords, err
}
