package config

type Environment struct {
	Environment            string `env:"Environment,default=local"`
	LogLevel               string `env:"LOG_LEVEL,default=warn"`
	RabbitUrl              string `env:"RABBIT_URL,default=amqp://guest:guest@localhost:5672/"`
	MongodbAddrs           string `env:"MONGO_DATABASE_ADDRS,default=mongodb+srv://adminApp:vWQuBeGLwDtr6B3I@yourfinances.slmnk.mongodb.net/?retryWrites=true&w=majority"`
	MongodbUser            string `env:"MONGO_DATABASE_USERNAME ,default=adminApp"`
	MongodbDatabase        string `env:"MONGO_DATABASE_NAME,default=audit-flow"`
	MongodbPassword        string `env:"MONGO_DATABASE_PASSWORD,default=vWQuBeGLwDtr6B3I"`
	MongodbMaxPoolSize     uint64 `env:"MONGO_MAX_POOL_SIZE,default=100"`
	MongodbMaxConnIdleTine int    `env:"MONGO_MAX_CONN_IDLE_TIME,default=60000"`
}
