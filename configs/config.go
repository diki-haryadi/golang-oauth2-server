package config

type MainConfig struct {
	Rest     RestConfig     `fig:"rest"`
	Mongo    MongoConfig    `fig:"mongo"`
	Postgres PostgresConfig `fig:"postgres"`
}

type (
	RestConfig struct {
		ListenAddress   string `fig:"listenAddress"`
		Port            string `fig:"port"`
		GracefulTimeout int    `fig:"gracefulTimeout"`
		AppName         string `fig:"appName"`
		ReadTimeout     int    `fig:"readTimeout"`
		WriteTimeout    int    `fig:"writeTimeout"`
		EnableSwagger   bool   `fig:"enableSwagger"`
	}
	MongoConfig struct {
		URI               string `yaml:"uri"`
		DB                string `yaml:"db"`
		ConnectionTimeout int    `yaml:"connectionTimeout"`
		PingTimeout       int    `yaml:"pingTimeout"`
	}
	PostgresConfig struct {
		Host     string `yaml:"host"`
		User     string `yml:"user"`
		Password string `yaml:"password"`
		Schema   string `yaml:"schema"`
	}
)
