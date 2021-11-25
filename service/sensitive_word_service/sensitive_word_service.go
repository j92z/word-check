package sensitive_word_service

import (
	"sensitive_words_check/config"
	"sensitive_words_check/pkg/dfa"
	"sensitive_words_check/pkg/dictionary"
)

type CheckSensitiveWordResult struct {
	Text      string   `json:"text"`
	Words     []string `json:"words"`
	HitWords  []string `json:"hit_words"`
	Sensitive bool     `json:"sensitive"`
}

func CheckSensitiveWord(words []string) []CheckSensitiveWordResult {
	dfaTool := dfa.GetDFA()
	var result = make([]CheckSensitiveWordResult, len(words))
	for i, word := range words {
		hitOriginWords, hitWords, sensitive := dfaTool.Check(word)
		result[i] = CheckSensitiveWordResult{
			Text:      word,
			Words:     hitOriginWords,
			HitWords:  hitWords,
			Sensitive: sensitive,
		}
	}
	return result
}

func AddSensitiveWord(word string) error {
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	return dty.AddWord(word)
}

func RemoveSensitiveWord(word string) error {
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	return dty.RemoveWord(word)
}