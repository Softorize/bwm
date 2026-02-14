package client

import (
	"fmt"
	"net/url"

	"github.com/Softorize/bwm/internal/models"
)

func (c *Client) GetLinkCounts(siteURL string) (*models.LinkCounts, error) {
	params := url.Values{"siteUrl": {siteURL}}
	var counts models.LinkCounts
	err := c.get("GetLinkCounts", params, &counts)
	return &counts, err
}

func (c *Client) GetLinkDetails(siteURL string, page int) (models.LinkDetailList, error) {
	params := url.Values{
		"siteUrl": {siteURL},
		"page":    {url.QueryEscape(fmt.Sprintf("%d", page))},
	}
	var details models.LinkDetailList
	err := c.get("GetInboundLinks", params, &details)
	return details, err
}
