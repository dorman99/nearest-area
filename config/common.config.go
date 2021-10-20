package config

type databaseConfig struct {
	Host     string
	User     string
	Port     int
	Password string
	Database string
}

func GetDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		Host:     "localhost",
		User:     "postgres",
		Port:     5435,
		Password: "password",
		Database: "postgres",
	}
}
