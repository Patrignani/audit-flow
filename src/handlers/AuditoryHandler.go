package handlers

import (
	"context"
	"encoding/json"

	"github.com/Patrignani/audit-flow/src/data"
	"github.com/Patrignani/audit-flow/src/models"
	"go.uber.org/zap"
)

const (
	collection = "auditory"
)

type AuditoryHandler struct {
	mongo  data.IMongoContext
	logger *zap.Logger
}

func NewAuditoryHandler(mongo data.IMongoContext, log *zap.Logger) *AuditoryHandler {
	return &AuditoryHandler{
		mongo:  mongo,
		logger: log,
	}
}

func (c *AuditoryHandler) Run(ctx context.Context, body []byte) {

	var auditory models.Auditory
	err := json.Unmarshal(body, &auditory)
	if err != nil {
		c.logger.Error("Error json", zap.Any("error", err))
		return
	}

	id, err := c.mongo.Insert(ctx, collection, auditory)

	if err != nil {
		c.logger.Error("Error inser audit", zap.Any("error", err))
		return
	}

	c.logger.Warn("Audit insert", zap.String("id", id))
}
