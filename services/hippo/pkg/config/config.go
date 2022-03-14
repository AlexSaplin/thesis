package config


type HippoConfig struct {
	GRPC     GRPCConfig           `json:"grpc"`
	Ardea    ArdeaClientConfig    `json:"ardea"`
	Selachii SelachiiClientConfig `json:"selachii"`
	Nalogi   NalogiClientConfig   `json:"nalogi"`
}

type GRPCConfig struct {
	Bind string `json:"bind"`
}

type ArdeaClientConfig struct {
	Target string `json:"target"`
}

type SelachiiClientConfig struct {
	Target string `json:"target"`
}

type NalogiClientConfig struct {
	Topic  string `json:"topic"`
	Target string `json:"target"`
}
