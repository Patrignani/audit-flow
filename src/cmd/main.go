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
)

func main() {

	ctx := context.Background()

	logger := logs.NewLogger(models.GetLoggingConfig())

	mongo := getMongoContext()

	cashFlow := getCashFlowConfig(mongo, logger)

	expenseControl := getExpenseControlConfig(mongo, logger)

	rabbitmqhelper.NewRabbitEventBuider(config.Env.RabbitUrl).
		Subscribe("cash-flow-audit", cashFlow).
		Subscribe("expense-control-audit", expenseControl).
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

func getExpenseControlConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	expenseControlFlow := handlers.NewAuditoryHandler(mongo, logger)

	expenseControlFlowSub := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "expense-control",
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
			Exchange: "expense-control",
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
			Action:    expenseControlFlow.Run,
		},
	}

	return expenseControlFlowSub
}

func getMongoContext() data.IMongoContext {

	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	mongo := data.GetInstance()

	if err := mongo.Initialize(ctx, config.Env.MongodbAddrs, config.Env.MongodbDatabase, config.Env.MongodbMaxPoolSize, time.Duration(config.Env.MongodbMaxConnIdleTine)*time.Millisecond); err != nil {

		log.Panic("Could not resolve Data access layer", err)
	}

	return mongo
}
