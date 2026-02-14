package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetSitemaps(siteURL string) (models.SitemapList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var sitemaps models.SitemapList
	err := c.get("GetSitemaps", params, &sitemaps)
	return sitemaps, err
}

func (c *Client) SubmitSitemap(siteURL, sitemapURL string) error {
	return c.post("SubmitSitemap", map[string]string{
		"siteUrl":    siteURL,
		"sitemapUrl": sitemapURL,
	}, nil)
}

func (c *Client) RemoveSitemap(siteURL, sitemapURL string) error {
	return c.post("RemoveSitemap", map[string]string{
		"siteUrl":    siteURL,
		"sitemapUrl": sitemapURL,
	}, nil)
}

func (c *Client) GetSitemapDetail(siteURL, sitemapURL string) (*models.SitemapDetail, error) {
	params := url.Values{
		"siteUrl":    {siteURL},
		"feedpath":   {sitemapURL},
	}
	var detail models.SitemapDetail
	err := c.get("GetSitemapDetailedInfo", params, &detail)
	return &detail, err
}
