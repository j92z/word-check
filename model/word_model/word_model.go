package word_model

import "sensitive_words_check/model"

type Word struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Word string `gorm:"index;comment:'敏感词';not null" json:"word"`
}

func AddWord(word string) error {
	return model.Db.Create(&Word{Word: word}).Error
}

func RemoveWord(word string) error {
	return model.Db.Where("word = ?", word).Delete(&Word{}).Error
}

func FindWord(word string) Word {
	var info Word
	model.Db.Where("word = ?", word).First(&info)
	return info
}

func GetAll() []*Word {
	var list []*Word
	model.Db.Find(&list)
	return list
}
