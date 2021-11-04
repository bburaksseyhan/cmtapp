package utils

type Configuration struct {
	Database DbSettings
}

type DbSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	Timeout  int
}
