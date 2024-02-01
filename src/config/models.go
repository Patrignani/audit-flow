package config

type Environment struct {
	ElasticUrl string `env:"ELASTIC_URL,default=http://127.0.0.1:9200/"`
	RabbitUrl  string `env:"RABBIT_URL,default=amqp://guest:guest@localhost:5672/"`
}
