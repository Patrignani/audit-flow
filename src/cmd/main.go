package main

import (
	"context"
	"log"
	"time"

	"github.com/Patrignani/audit-flow/src/config"
	"github.com/Patrignani/audit-flow/src/data"
	"github.com/Patrignani/audit-flow/src/handlers"
	"github.com/Patrignani/audit-flow/src/logs"
	"github.com/Patrignani/audit-flow/src/models"
	rabbitmqhelper "github.com/Patrignani/rabbit-mq-helper"
	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx := context.Background()

	logger := logs.NewLogger(models.GetLoggingConfig())

	mongo := getMongoContext()

	cashFlow := getCashFlowConfig(mongo, logger)

	rabbitmqhelper.NewRabbitEventBuider(config.Env.RabbitUrl).
		Subscribe("cash-flow-audit", cashFlow).
		Run(ctx)

}

func getCashFlowConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	auditCashFlow := handlers.NewAuditoryHandler(mongo, logger)

	cashFlow := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "cash-flow",
			Kind:       "fanout",
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		Queue: rabbitmqhelper.QueueOptions{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  true,
			NoWait:     false,
			Args:       nil,
		},
		Bind: rabbitmqhelper.BindOptions{
			Key:      "",
			Exchange: "cash-flow",
			NoWait:   false,
			Args:     nil,
		},
		Consume: rabbitmqhelper.ConsumeOptions{
			Consumer:  "",
			AutoAck:   true,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
			Action:    auditCashFlow.Run,
		},
	}

	return cashFlow
}

func getMongoContext() data.IMongoContext {

	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	mongo := data.GetInstance()
	credential := options.Credential{
		Username:      config.Env.MongodbUser,
		Password:      config.Env.MongodbPassword,
		PasswordSet:   true,
		AuthSource:    config.Env.MongodbDatabase,
		AuthMechanism: config.Env.MongodbAuth,
	}

	if err := mongo.Initialize(ctx, credential, "mongodb://"+config.Env.MongodbHosts+":"+config.Env.MongodbPort,
		config.Env.MongodbDatabase, &config.Env.MongodbReplicaset); err != nil {
		log.Println("Could not resolve Data access layer", err)
	}

	return mongo
}
