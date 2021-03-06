package config

// Config contains app configutarion
type Config struct {
	Port     string
	Database DB
}

// DB contains database credentials
type DB struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}
