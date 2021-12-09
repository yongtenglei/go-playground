package config

type Config struct {
	KafkaConfig   `ini:"kafka"`
	EtcdConfig    `ini:"etcd"`
	CollectConfig `ini:"collect"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogfilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}
