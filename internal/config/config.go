package config

import (
	"log"

	"github.com/caarlos0/env"

	_ "time/tzdata"
)

//Configuration ...
type Configuration struct {
	DBConCampaign string `env:"DB_CON_CAMPAIGN" envDefault:"sqlserver://sa:example-db123@localhost:1433?database=test_lh_campaign"`
	// DBConRoboConductor     string `env:"DB_CON_ROBO_CONDUCTOR" envDefault:"sqlserver://sa:example-db123@localhost:1433?database=test_lh_robo_conductor"`
	ENV             string `env:"ENV"`
	Port            string `env:"PORT"`
	PortRestful     string `env:"PORT_RESTFUL"`
	Locale          string `env:"LOCALE" envDefault:"Asia/Bangkok"`
	LogFormat       string `env:"LOG_FORMAT" envDefault:""`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	IsEnableProtoV1 bool   `env:"IS_ENABLE_PROTO_V1" envDefault:"true"`
	IsEnableProtoV2 bool   `env:"IS_ENABLE_PROTO_V2" envDefault:"true"`
	IsDebugDB       bool   `env:"IS_DEBUG_DB" envDefault:"true"`
	//JWT CONFIG
	AESSecretKey     string `env:"AES_SECRET_KEY"`
	SessionTimeout   int    `env:"SESSION_TIMEOUT"`
	PrivateKeyPath   string `env:"PRIVATE_KEY_PATH"`
	PublicKeyPath    string `env:"PUBLIC_KEY_PATH"`
	MessagingAMQPUri string `env:"MESSAGING_AMQP_URI"`
}

//NewConfiguration ...
func NewConfiguration() *Configuration {
	c := &Configuration{}

	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	return c
}
