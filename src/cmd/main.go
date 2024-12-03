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

	estimatedSpend := getEstimatedSpendConfig(mongo, logger)

	moneyMovement := getMoneyMovementConfig(mongo, logger)

	consolidated := getConsolidatedConfig(mongo, logger)

	rabbitmqhelper.NewRabbitEventBuider(config.Env.RabbitUrl).
		Subscribe("cash-flow-audit", cashFlow).
		Subscribe("expense-control-audit", expenseControl).
		Subscribe("estimated-spend-audit", estimatedSpend).
		Subscribe("money-movement-audit", moneyMovement).
		Subscribe("consolidated-audit", consolidated).
		Run(ctx)
}

func getCashFlowConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	auditCashFlow := handlers.NewAuditoryHandler(mongo, logger)

	cashFlow := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "cash-flow",
			Kind:       "topic",
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
			Key:      "*",
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
	expenseControl := handlers.NewAuditoryHandler(mongo, logger)

	expenseControlSub := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "expense-control",
			Kind:       "topic",
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
			Key:      "*",
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
			Action:    expenseControl.Run,
		},
	}

	return expenseControlSub
}

func getEstimatedSpendConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	estimatedSend := handlers.NewAuditoryHandler(mongo, logger)

	estimatedSendSub := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "estimated-spend",
			Kind:       "topic",
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
			Key:      "*",
			Exchange: "estimated-spend",
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
			Action:    estimatedSend.Run,
		},
	}

	return estimatedSendSub
}

func getMoneyMovementConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	moneyMovementSend := handlers.NewAuditoryHandler(mongo, logger)

	moneyMovementSendSub := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "money-movement",
			Kind:       "topic",
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
			Key:      "*",
			Exchange: "money-movement",
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
			Action:    moneyMovementSend.Run,
		},
	}

	return moneyMovementSendSub
}

func getConsolidatedConfig(mongo data.IMongoContext, logger *zap.Logger) *rabbitmqhelper.Subscribe {
	consolidatedSend := handlers.NewAuditoryHandler(mongo, logger)

	consolidatedSendSub := &rabbitmqhelper.Subscribe{
		Exchange: rabbitmqhelper.ExchangeOptions{
			Name:       "consolidated",
			Kind:       "topic",
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
			Key:      "*",
			Exchange: "consolidated",
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
			Action:    consolidatedSend.Run,
		},
	}

	return consolidatedSendSub
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
