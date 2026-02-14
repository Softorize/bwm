package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetCrawlStats(siteURL string) (models.CrawlStatsList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var stats models.CrawlStatsList
	err := c.get("GetCrawlStats", params, &stats)
	return stats, err
}

func (c *Client) GetCrawlIssues(siteURL string) (models.CrawlIssueList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var issues models.CrawlIssueList
	err := c.get("GetCrawlIssues", params, &issues)
	return issues, err
}

func (c *Client) GetCrawlSettings(siteURL string) (*models.CrawlSettings, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var settings models.CrawlSettings
	err := c.get("GetCrawlSettings", params, &settings)
	return &settings, err
}
