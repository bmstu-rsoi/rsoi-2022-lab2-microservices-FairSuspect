package utils

type Configuration struct {
	LogFile            string `json:"log_file"`
	Port               uint16 `json:"port"`
	FlightsEndpoint    string `json:"flights-endpoint"`
	TicketsEndpoint    string `json:"tickets-endpoint"`
	PrivilegesEndpoint string `json:"privileges-endpoint"`
}

var (
	Config Configuration
)

func InitConfig() {
	Config = Configuration{
		"logs/server.log",
		8080,

		"http://flights:8060",
		"http://tickets:8070",
		"http://privileges:8050",
	}
}
