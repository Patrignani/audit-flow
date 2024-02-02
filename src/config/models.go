package config

type Environment struct {
	Environment       string `env:"Environment,default=local"`
	LogLevel          string `env:"LOG_LEVEL,default=warn"`
	RabbitUrl         string `env:"RABBIT_URL,default=amqp://guest:guest@localhost:5672/"`
	MongodbUser       string `env:"MONGODB_USER,default=auditAdmim"`
	MongodbPassword   string `env:"MONGODB_PASSWORD,default=f0cd47b4b7364a7e9b87e1a377b7dddf"`
	MongodbHosts      string `env:"MONGODB_HOST,default=localhost"`
	MongodbPort       string `env:"MONGODB_PORT,default=27017"`
	MongodbAuth       string `env:"MONGODB_AUTH,default=SCRAM-SHA-1"`
	MongodbDatabase   string `env:"MONGODB,default=audit-flow"`
	MongodbReplicaset string `env:"MONGODB_REPLICASET"`
}
