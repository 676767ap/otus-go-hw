package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DevMode bool `env:"PROJECT_DEV_MODE" env-default:"true"`

	App struct {
		Storage string `env:"PROJECT_STORAGE"`
	}

	PostgreSQL struct {
		User     string `env:"PROJECT_PSQL_USER" env-required:"true"`
		Password string `env:"PROJECT_PSQL_PASSWORD" env-required:"true"`
		Host     string `env:"PROJECT_PSQL_HOST" env-required:"true"`
		Port     string `env:"PROJECT_PSQL_PORT" env-default:"5432"`
		Database string `env:"PROJECT_PSQL_DATABASE" env-required:"true"`
	}

	Session struct {
		MaxAge     int    `env:"PROJECT_MAXAGE" env-default:"3600"`
		CookieName string `env:"PROJECT_COOKIE_NAME" env-default:"gosession"`
		SessionKey string `env:"PROJECT_SESSION_KEY" env-default:"7ca21d0aae117345c7eabac70cea6fa9133af82201b2cb4186754e5945764de9408b5f31a1620aa8434bd586942f600f261aff9a5cb842c1773a03aeb46c11a7"`
	}

	Server struct {
		Address string `env:"PROJECT_SERVER_ADDRESS" env-default:""`
		Port    string `env:"PROJECT_SERVER_PORT" env-default:"8086"`
	}

	RabbitMQ struct {
		Protocol  string `env:"RABBITMQ_PROTOCOL" required:"true"`
		Username  string `env:"RABBITMQ_USERNAME" required:"true"`
		Password  string `env:"RABBITMQ_PASSWORD" required:"true"`
		Host      string `env:"RABBITMQ_HOST" required:"true"`
		Port      int    `env:"RABBITMQ_PORT" required:"true"`
		Exchange  RabbitMQExchange
		Queue     RabbitMQQueue
		Publisher RabbitMQPublisher
		Consumer  RabbitMQConsumer
	}
}

func (c *Config) BuildDSNRabbitMQ() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", c.RabbitMQ.Protocol, c.RabbitMQ.Username, c.RabbitMQ.Password, c.RabbitMQ.Host, c.RabbitMQ.Port)
}

type RabbitMQExchange struct {
	Name       string `env:"RABBITMQ_EXCHANGE_NAME" required:"true"`
	Kind       string `env:"RABBITMQ_EXCHANGE_KIND" required:"true"`
	Durable    bool   `env:"RABBITMQ_EXCHANGE_DURABLE" required:"true"`
	AutoDelete bool   `env:"RABBITMQ_EXCHANGE_AUTO_DELETE" required:"true"`
	Internal   bool   `env:"RABBITMQ_EXCHANGE_INTERNAL" required:"true"`
	NoWait     bool   `env:"RABBITMQ_EXCHANGE_NO_WAIT" required:"true"`
}

type RabbitMQQueue struct {
	Name       string `env:"RABBITMQ_QUEUE_NAME" required:"true"`
	Durable    bool   `env:"RABBITMQ_QUEUE_DURABLE" required:"true"`
	AutoDelete bool   `env:"RABBITMQ_QUEUE_AUTO_DELETE" required:"true"`
	Exclusive  bool   `env:"RABBITMQ_QUEUE_EXCLUSIVE" required:"true"`
	NoWait     bool   `env:"RABBITMQ_QUEUE_NO_WAIT" required:"true"`
	BindNoWait bool   `env:"RABBITMQ_QUEUE_BIND_NO_WAIT" required:"true"`
	BindingKey string `env:"RABBITMQ_QUEUE_BINDING_KEY" required:"true"`
}

type RabbitMQPublisher struct {
	Mandatory  bool   `env:"RABBITMQ_PUBLISH_MANDATORY" required:"true"`
	Immediate  bool   `env:"RABBITMQ_PUBLISH_IMMEDIATE" required:"true"`
	RoutingKey string `env:"RABBITMQ_PUBLISH_ROUTING_KEY" required:"true"`
}

type RabbitMQConsumer struct {
	Name      string `env:"RABBITMQ_CONSUMER_NAME" required:"true"`
	AutoAck   bool   `env:"RABBITMQ_CONSUMER_AUTO_ACK" required:"true"`
	Exclusive bool   `env:"RABBITMQ_CONSUMER_EXCLUSIVE" required:"true"`
	NoLocal   bool   `env:"RABBITMQ_CONSUMER_NO_LOCAL" required:"true"`
	NoWait    bool   `env:"RABBITMQ_CONSUMER_NO_WAIT" required:"true"`
}

func (cfg *Config) GetAddressWithPort() string {
	return fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)
}

var instance *Config

func LoadConfig() *Config {
	instance = &Config{}
	if err := cleanenv.ReadEnv(instance); err != nil {
		helpText := "OTUS student project"
		help, _ := cleanenv.GetDescription(instance, &helpText)
		fmt.Println(help)
		fmt.Println(err)
		os.Exit(1)
	}
	return instance
}

func GetConfig() *Config {
	return instance
}
