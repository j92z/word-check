package config

import (
	"sensitive_words_check/constant"
)

type DictionarySetting struct {
	StoreType constant.DictionaryStoreType
}

var DictionaryConfig DictionarySetting

func (d *DictionarySetting) Init(env EnvType) {
	switch env {
	case DevEnv:
		DictionaryConfig.StoreType = constant.FileStore
		break
	case ProdEnv:
		DictionaryConfig.StoreType = constant.FileStore
		break
	default:
		DictionaryConfig.StoreType = constant.FileStore
	}
}