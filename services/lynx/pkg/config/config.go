package config

type LynxConfig struct {
	API        APIConfig           `json:"api"`
	Ardea      ArdeaClientConfig   `json:"ardea"`
	Hippo      HippoClientConfig   `json:"hippo"`
	Ovis       OvisClientConfig    `json:"ovis"`
	S3         S3ClientConfig      `json:"s3"`
	Auth       AuthClientConfig    `json:"auth"`
	Gorilla    GorillaClientConfig `json:"gorilla"`
	Slav       SlavClientConfig    `json:"slav"`
	Ibis       IbisClientConfig    `json:"ibis"`
	Rhino      RhinoClientConfig   `json:"rhino"`
	Arietes    ArietesClientConfig `json:"arietes"`
	Picus      PicusClientConfig   `json:"picus"`
	Prometheus PromClientConfig    `json:"prometheus"`
}

type APIConfig struct {
	Bind          string      `json:"bind"`
	Users         []UserToken `json:"users"`
	ContainersURL string      `json:"containers_url"`
}

type UserToken struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

type ArdeaClientConfig struct {
	Target string `json:"target"`
}

type IbisClientConfig struct {
	Target string `json:"target"`
}

type HippoClientConfig struct {
	Target string `json:"target"`
}

type RhinoClientConfig struct {
	Target string `json:"target"`
}

type OvisClientConfig struct {
	Target string `json:"target"`
	Queue  string `json:"queue"`
}

type S3ClientConfig struct {
	Target    string `json:"target"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
	Secure    bool   `json:"secure"`

	ModelsBucket    string `json:"models_bucket"`
	FunctionsBucket string `json:"functions_bucket"`
}

type PromClientConfig struct {
	Target string `json:"target"`
}

type AuthClientConfig struct {
	Target string `json:"target"`
}

type GorillaClientConfig struct {
	Target string `json:"target"`
}

type SlavClientConfig struct {
	Target string `json:"target"`
}

type ArietesClientConfig struct {
	Target string `json:"target"`
	Topic  string `json:"topic"`
}

type PicusClientConfig struct {
	Target string `json:"target"`
}
