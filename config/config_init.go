package config

type cngInterface interface {
	Init(env EnvType)
}

type EnvType int

const (
	DevEnv	EnvType = iota
	ProdEnv
)

func InitConfig(env ...string) {
	var envInfo EnvType
	if len(env) > 0 {
		switch env[0] {
		case "dev":
			envInfo = DevEnv
			break
		case "prod":
			envInfo = ProdEnv
			break
		default:
			envInfo = DevEnv
		}
	} else {
		envInfo = DevEnv
	}
	var configs = []cngInterface{&DbConfig, &DictionaryConfig}
	for _, c := range configs {
		c.Init(envInfo)
	}
}
