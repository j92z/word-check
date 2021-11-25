package config

type DatabaseSetting struct {
	Type     string
	User     string
	Password string
	Port     int
	Host     string
	DbName   string
	ShowDbLog bool
	MaxIdleConns int
	MaxOpenConns int
}

var DbConfig DatabaseSetting

func (d *DatabaseSetting) Init(env EnvType) {
	switch env {
	case DevEnv:
		DbConfig.Host = "192.168.10.247"
		DbConfig.Type = "mysql"
		DbConfig.User = "root"
		DbConfig.Password = "123456"
		DbConfig.Port = 3306
		DbConfig.DbName = "sensitive-check"
		DbConfig.ShowDbLog = true
		DbConfig.MaxIdleConns = 10
		DbConfig.MaxOpenConns = 100
		break
	case ProdEnv:
		DbConfig.Host = "192.168.10.247"
		DbConfig.Type = "mysql"
		DbConfig.User = "root"
		DbConfig.Password = "123456"
		DbConfig.Port = 3306
		DbConfig.DbName = "sensitive-check"
		DbConfig.ShowDbLog = false
		DbConfig.MaxIdleConns = 10
		DbConfig.MaxOpenConns = 100
		break
	default:
		DbConfig.Host = "192.168.10.247"
		DbConfig.Type = "mysql"
		DbConfig.User = "root"
		DbConfig.Password = "123456"
		DbConfig.Port = 3306
		DbConfig.DbName = "sensitive-check"
		DbConfig.ShowDbLog = false
		DbConfig.MaxIdleConns = 10
		DbConfig.MaxOpenConns = 100
	}
}