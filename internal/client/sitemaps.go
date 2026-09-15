package client

import (
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

// Bing Webmaster's JSON API calls sitemaps "feeds": the methods are GetFeeds,
// SubmitFeed, RemoveFeed and GetFeedDetails, and the sitemap parameter is
// feedUrl. The earlier GetSitemaps/SubmitSitemap names answered 404.

func (c *Client) GetSitemaps(siteURL string) (models.SitemapList, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var sitemaps models.SitemapList
	err := c.get("GetFeeds", params, &sitemaps)
	return sitemaps, err
}

func (c *Client) SubmitSitemap(siteURL, sitemapURL string) error {
	return c.post("SubmitFeed", map[string]string{
		"siteUrl": siteURL,
		"feedUrl": sitemapURL,
	}, nil)
}

func (c *Client) RemoveSitemap(siteURL, sitemapURL string) error {
	return c.post("RemoveFeed", map[string]string{
		"siteUrl": siteURL,
		"feedUrl": sitemapURL,
	}, nil)
}

// GetSitemapDetail returns the feed itself followed by the feeds it indexes
// (for a sitemap index), which is what GetFeedDetails answers.
func (c *Client) GetSitemapDetail(siteURL, sitemapURL string) (models.SitemapList, error) {
	params := url.Values{
		"siteUrl": {siteURL},
		"feedUrl": {sitemapURL},
	}
	var detail models.SitemapList
	err := c.get("GetFeedDetails", params, &detail)
	return detail, err
}
