package utils

// Configuration is hold the Database
type Configuration struct {
	Database DbSettings
}

// DbSettings is hold the database configurations
type DbSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	Timeout  int
}
