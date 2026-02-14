package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) SubmitURL(siteURL, pageURL string) error {
	return c.post("SubmitUrl", map[string]string{
		"siteUrl": siteURL,
		"url":     pageURL,
	}, nil)
}

func (c *Client) SubmitURLBatch(siteURL string, urls []string) error {
	return c.post("SubmitUrlbatch", map[string]any{
		"siteUrl": siteURL,
		"urlList": urls,
	}, nil)
}

func (c *Client) GetURLSubmissionQuota(siteURL string) (*models.URLSubmissionQuota, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var quota models.URLSubmissionQuota
	err := c.get("GetUrlSubmissionQuota", params, &quota)
	return &quota, err
}
