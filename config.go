package scylla

// Config is scylla config.
type Config struct {
	Hosts       []string `json:"hosts" env:"SCYLLA_HOSTS"`
	Username    string   `json:"username" env:"SCYLLA_USERNAME"`
	Password    string   `json:"password" env:"SCYLLA_PASSWORD"`
	Keyspace    string   `json:"keyspace" env:"SCYLLA_KEYSPACE"`
	Consistency string   `json:"consistency" env:"SCYLLA_CONSISTENCY"`
}
