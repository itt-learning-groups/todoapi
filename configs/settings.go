package configs

// Settings contains app environment configuration settings
type Settings struct {
	ServerAddr string `envcfg:"SERVER_ADDR" envcfgDefault:""`
	ServerPort string `envcfg:"SERVER_PORT" envcfgDefault:"8080"`

	DatabaseHostName          string `envcfg:"DB_HOSTNAME" envcfgDefault:""`
	DatabaseDBName            string `envcfg:"DB_DBNAME" envcfgDefault:""`
	DatabaseTodosCollection   string `envcfg:"DB_TODOS_COLLECTION" envcfgDefault:""`
	DatabaseUserName          string `envcfg:"DB_USERNAME" envcfgDefault:""`
	DatabasePswd              string `envcfg:"DB_PSWD" envcfgDefault:""`
	DatabaseCxnTimeoutSeconds int64  `envcfg:"DB_TIMEOUT" envcfgDefault:"10"`
}
