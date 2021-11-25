package dictionary

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sensitive_words_check/config"
	"sensitive_words_check/constant"
	"sensitive_words_check/model"
	"sensitive_words_check/model/word_model"
	"sensitive_words_check/setup"
	"testing"
)

func init() {
	config.InitConfig()
	model.Setup()
	setup.CheckTable()
}

func TestOpenDictionary(t *testing.T) {
	pathInfo := "test.txt"
	dct := OpenDictionary(constant.FileStore, pathInfo)
	defer os.Remove(pathInfo)
	assert.Equal(t, dct.path, "test.txt")

	dct2 := OpenDictionary(constant.MysqlStore)
	assert.Equal(t, dct2.storeType, constant.MysqlStore)
}

func TestDictionary_AddWord(t *testing.T) {
	pathInfo := "test.txt"
	dct := OpenDictionary(constant.FileStore, pathInfo)
	defer os.Remove(pathInfo)
	word := "xxxx"
	err := dct.AddWord(word)
	assert.Nil(t, err)
	assert.Equal(t, dct.Words[0], word)

	dct2 := OpenDictionary(constant.MysqlStore)
	err = dct2.AddWord(word)
	assert.Nil(t, err)
	defer word_model.RemoveWord(word)
	wordInfo := word_model.FindWord(word)
	assert.Equal(t, wordInfo.Word, word)
}

func TestDictionary_RemoveWord(t *testing.T) {
	pathInfo := "test.txt"
	dct := OpenDictionary(constant.FileStore, pathInfo)
	defer os.Remove(pathInfo)
	word := "xxxx"
	err := dct.AddWord(word)
	assert.Nil(t, err)
	err = dct.RemoveWord(word)
	assert.Nil(t, err)
	assert.Equal(t, len(dct.Words), 0)

	dct2 := OpenDictionary(constant.MysqlStore)
	err = word_model.AddWord(word)
	assert.Nil(t, err)
	err = dct2.RemoveWord(word)
	assert.Nil(t, err)
}