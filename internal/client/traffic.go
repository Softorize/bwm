package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetTrafficStats(siteURL string) (models.TrafficStatsList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var stats models.TrafficStatsList
	err := c.get("GetRankAndTrafficStats", params, &stats)
	return stats, err
}

func (c *Client) GetQueryStats(siteURL string) (models.QueryStatsList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var stats models.QueryStatsList
	err := c.get("GetQueryStats", params, &stats)
	return stats, err
}

func (c *Client) GetPageStats(siteURL, pageURL string) (models.PageStatsList, error) {
	params := url.Values{"siteUrl": {siteURL}, "page": {pageURL}}
	var stats models.PageStatsList
	err := c.get("GetPageStats", params, &stats)
	return stats, err
}
