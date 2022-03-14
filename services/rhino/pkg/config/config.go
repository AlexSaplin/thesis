package config

type RhinoConfig struct {
	GRPC   GRPCConfig         `json:"grpc"`
	Ardea  IbisClientConfig   `json:"ibis"`
	Runner RunnerConfig       `json:"runner"`
	Nalogi NalogiClientConfig `json:"nalogi"`
}

type GRPCConfig struct {
	Bind string `json:"bind"`
}

type IbisClientConfig struct {
	Target string `json:"target"`
}

type RunnerConfig struct {
	MaxClientsPerFunction int                  `json:"max_clients_per_function"`
	Clients               []DockerClientConfig `json:"clients"`
}

type DockerClientConfig struct {
	Target        string `json:"target"`
	ContainerHost string `json:"container_host"`
}

type NalogiClientConfig struct {
	Topic  string `json:"topic"`
	Target string `json:"target"`
}
