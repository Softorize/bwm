package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetSites() (models.SiteList, error) {
	var sites models.SiteList
	err := c.get("GetUserSites", nil, &sites)
	return sites, err
}

func (c *Client) AddSite(siteURL string) error {
	return c.post("AddSite", map[string]string{"siteUrl": siteURL}, nil)
}

func (c *Client) RemoveSite(siteURL string) error {
	params := url.Values{"siteUrl": {siteURL}}
	var result any
	return c.get("RemoveSite", params, &result)
}

func (c *Client) VerifySite(siteURL string) error {
	return c.post("VerifySite", map[string]string{"siteUrl": siteURL}, nil)
}
