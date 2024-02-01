package handlers

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type AuditoryHandler struct {
	client *elastic.Client
}

func NewAuditoryHandler(client *elastic.Client) *AuditoryHandler {
	return &AuditoryHandler{
		client: client,
	}
}

func (c *AuditoryHandler) Run(ctx context.Context, body []byte) {
	c.client.Index().
		Index("auditory").
		BodyJson(body).
		Do(ctx)
}
