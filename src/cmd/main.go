package main

import (
	"context"
	"log"

	"github.com/Patrignani/audit-flow/src/config"
	"github.com/Patrignani/audit-flow/src/handlers"
	rabbitmqhelper "github.com/Patrignani/rabbit-mq-helper"

	"github.com/olivere/elastic/v7"
)

func main() {

	ctx := context.Background()

	elasticClient, err := elastic.NewClient(elastic.SetURL(config.Env.ElasticUrl))
	if err != nil {
		log.Panicf("Erro ao conectar ao Elasticsearch: %v", err)
	}

	cashFlow := getCashFlowConfig(elasticClient)

	rabbitmqhelper.NewRabbitEventBuider(config.Env.ElasticUrl).
		Subscribe("cash-flow-audit", cashFlow).
		Run(ctx)

}

func getCashFlowConfig(client *elastic.Client) *rabbitmqhelper.Subscribe {
	auditCashFlow := handlers.NewAuditoryHandler(client)

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
