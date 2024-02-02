package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Patrignani/audit-flow/src/data"
	"github.com/Patrignani/audit-flow/src/models"
)

const (
	collection = "auditory"
)

type AuditoryHandler struct {
	mongo data.IMongoContext
}

func NewAuditoryHandler(mongo data.IMongoContext) *AuditoryHandler {
	return &AuditoryHandler{
		mongo: mongo,
	}
}

func (c *AuditoryHandler) Run(ctx context.Context, body []byte) {

	var auditory models.Auditory
	err := json.Unmarshal(body, &auditory)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	id, err := c.mongo.Insert(ctx, collection, auditory)

	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	println("Id criado " + id)
}
