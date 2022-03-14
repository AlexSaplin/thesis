package config

type GorillaConfig struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

type DBConfig struct {
	Target string `json:"target"`
}

type ServerConfig struct {
	Bind string `json:"bind"`
}
