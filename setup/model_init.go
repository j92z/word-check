package setup

import (
	"sensitive_words_check/model"
	"sensitive_words_check/model/word_model"
)

func CheckTable() {
	model.Db.AutoMigrate(&word_model.Word{})
}
