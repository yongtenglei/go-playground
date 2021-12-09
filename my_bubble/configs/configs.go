package configs

var Conf = new(Config)

type Config struct {
	*AppConf   `mapstructure:"app"`
	*LogConf   `mapstructure:"log"`
	*MysqlConf `mapstructure:"mysql"`
}

type AppConf struct {
	Name     string `mapstructure:"name"`
	Mode     string `mapstructure:"mode"`
	Version  string `mapstructure:"version"`
	Language string `mapstructure:"language"`
	Port     int    `mapstructure:"port"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConf struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxLifeTime  int    `mapstructure:"max_life_time"`
}
