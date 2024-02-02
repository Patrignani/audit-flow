package models

import "github.com/Patrignani/audit-flow/src/config"

type Environment string

const (
	LocalEnv            Environment = "local"
	DevelopmentEnv      Environment = "dev"
	QualityAssuranceEnv Environment = "qa"
	ProductionEnv       Environment = "prod"
)

func GetEnvironmentConfig() Environment {
	return Environment(config.Env.Environment)
}
