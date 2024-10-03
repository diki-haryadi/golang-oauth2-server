package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang-oauth2-server/internal/pkg/constant"
	"golang-oauth2-server/internal/pkg/env"
	"log"
)

type Config struct {
	App              AppConfig
	Grpc             GrpcConfig
	Http             HttpConfig
	Postgres         PostgresConfig
	SampleExtService GrpcConfig
	Kafka            KafkaConfig
	Sentry           SentryConfig
}

var BaseConfig *Config

type AppConfig struct {
	AppEnv  string `json:"app_env" envconfig:"APP_ENV"`
	AppName string `json:"app_name" envconfig:"APP_NAME"`
}

type PostgresConfig struct {
	Host            string `json:"host" envconfig:"PG_HOST"`
	Port            string `json:"port" envconfig:"PG_PORT"`
	User            string `json:"user" envconfig:"PG_USER"`
	Pass            string `json:"pass" envconfig:"PG_PASS"`
	DBName          string `json:"db_name" envconfig:"PG_DB"`
	MaxConn         int    `json:"max_conn" envconfig:"PG_MAX_CONNECTIONS"`
	MaxIdleConn     int    `json:"max_idle_conn" envconfig:"PG_MAX_IDLE_CONNECTIONS"`
	MaxLifeTimeConn int    `json:"max_life_time_conn" envconfig:"PG_MAX_LIFETIME_CONNECTIONS"`
	SslMode         string `json:"ssl_mode" envconfig:"PG_SSL_MODE"`
}
type GrpcConfig struct {
	Port int    `json:"port" envconfig:"GRPC_PORT"`
	Host string `json:"host" envconfig:"GRPC_HOST" `
}

type HttpConfig struct {
	Port int    `json:"port" envconfig:"HTTP_PORT"`
	Host string `json:"host" envconfig:"HTTP_HOST"`
}

type KafkaConfig struct {
	Enabled       bool     `json:"enabled" envconfig:"KAFKA_ENABLED"`
	LogEvents     bool     `json:"log_events" envconfig:"KAFKA_LOG_EVENTS"`
	ClientId      string   `json:"client_id" envconfig:"KAFKA_CLIENT_ID"`
	ClientGroupId string   `json:"client_group_id" envconfig:"KAFKA_CLIENT_GROUP_ID"`
	ClientBrokers []string `json:"client_brokers" envconfig:"KAFKA_CLIENT_BROKERS"`
	Topic         string   `json:"topic" envconfig:"KAFKA_TOPIC"`
}

type SentryConfig struct {
	Dsn string `json:"dsn" envconfig:"SENTRY_DSN"`
}

func init() {
	//BaseConfig = newConfig()
	BaseConfig = LoadConfig()
}

func newConfig() *Config {
	return &Config{
		App: AppConfig{
			AppEnv:  env.New("APP_ENV", constant.AppEnvDev).AsString(),
			AppName: env.New("APP_NAME", constant.AppName).AsString(),
		},
		Grpc: GrpcConfig{
			Port: env.New("GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Http: HttpConfig{
			Port: env.New("HTTP_PORT", constant.HttpPort).AsInt(),
			Host: env.New("HTTP_HOST", constant.HttpHost).AsString(),
		},
		Postgres: PostgresConfig{
			Host:            env.New("PG_HOST", nil).AsString(),
			Port:            env.New("PG_PORT", nil).AsString(),
			User:            env.New("PG_USER", nil).AsString(),
			Pass:            env.New("PG_PASS", nil).AsString(),
			DBName:          env.New("PG_DB", nil).AsString(),
			MaxConn:         env.New("PG_MAX_CONNECTIONS", constant.PgMaxConn).AsInt(),
			MaxIdleConn:     env.New("PG_MAX_IDLE_CONNECTIONS", constant.PgMaxIdleConn).AsInt(),
			MaxLifeTimeConn: env.New("PG_MAX_LIFETIME_CONNECTIONS", constant.PgMaxLifeTimeConn).AsInt(),
			SslMode:         env.New("PG_SSL_MODE", constant.PgSslMode).AsString(),
		},
		SampleExtService: GrpcConfig{
			Port: env.New("SAMPLE_EXT_SERVICE_GRPC_PORT", constant.GrpcPort).AsInt(),
			Host: env.New("SAMPLE_EXT_SERVICE_GRPC_HOST", constant.GrpcHost).AsString(),
		},
		Kafka: KafkaConfig{
			Enabled:       env.New("KAFKA_ENABLED", nil).AsBool(),
			LogEvents:     env.New("KAFKA_LOG_EVENTS", nil).AsBool(),
			ClientId:      env.New("KAFKA_CLIENT_ID", nil).AsString(),
			ClientGroupId: env.New("KAFKA_CLIENT_GROUP_ID", nil).AsString(),
			ClientBrokers: env.New("KAFKA_CLIENT_BROKERS", nil).AsStringSlice(","),
			Topic:         env.New("KAFKA_TOPIC", nil).AsString(),
		},
		Sentry: SentryConfig{
			Dsn: env.New("SENTRY_DSN", nil).AsString(),
		},
	}
}

func LoadConfig() *Config {
	_ = godotenv.Overload()
	var configLoader Config

	if err := envconfig.Process("BaseConfig", &configLoader); err != nil {
		log.Printf("error load configs: %v", err)
	}

	BaseConfig = &configLoader
	spew.Dump(configLoader)
	return &configLoader
}

func IsDevEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvDev
}

func IsProdEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvProd
}

func IsTestEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvTest
}
