package config

// Settings app settings
type Config struct {
	AppParams        Params           `mapstructure:"app"`
	PostgresDbParams PostgresDbParams `mapstructure:"db"`
}

// Params contains params of server meta data
type Params struct {
	Host string `mapstructure:"host"`
	Port    int `mapstructure:"port"`
}

// PostgresMegafonDbParams conteins params of postgresql db server
type PostgresDbParams struct {
	Host   	 string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Database string `mapstructure:"name"`
}


