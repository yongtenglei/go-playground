package conf

var Conf = new(Config)

type Config struct {
	*AppConf `mapstructure:"app"`
	//*LogConf   `mapstructure:"log"`
	*MysqlConf  `mapstructure:"mysql"`
	*ServerConf `mapstructure:"server"`
}

type AppConf struct {
	RunMode   string `mapstructure:"run_mode"`
	JwtSecret string `mapstructure:"jwt_secret"`
	PageSize  int    `mapstructure:"page_size"`
}

//type LogConf struct {
//Level      string `mapstructure:"level"`
//Filename   string `mapstructure:"filename"`
//MaxSize    int    `mapstructure:"max_size"`
//MaxAge     int    `mapstructure:"max_age"`
//MaxBackups int    `mapstructure:"max_backups"`
//}

type ServerConf struct {
	HttpPort     int `mapstructure:"http_port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
}

type MysqlConf struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	DbName       string `mapstructure:"db_name"`
	TablePrefix  string `mapstructure:"table_prefix"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxLifeTime  int    `mapstructure:"max_life_time"`
}
