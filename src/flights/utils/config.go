package utils

type DBConfig struct {
	Type string `json:"type"`
	Name string `json:"name"`

	User     string `json:"user"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Configuration struct {
	DB      DBConfig `json:"db"`
	LogFile string   `json:"log_file"`
	Port    uint16   `json:"port"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		DBConfig{
			"postgres",
			"flights",
			"program",
			"test",
			"postgres",
			"5432",
		},
		"logs/server.log",
		8060,
	}
}
