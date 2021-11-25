package router

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sensitive_words_check/config"
	"sensitive_words_check/pkg/dictionary"
	"strings"
	"testing"
)

func TestAddSensitiveWord(t *testing.T) {
	word := "测试数据"
	defer func() {
		dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
		_ = dty.RemoveWord(word)
	}()
	router := NewRouter()
	w := httptest.NewRecorder()
	reader := strings.NewReader("{\"word\": \""+word+"\"}")
	req, _ := http.NewRequest("POST", "/sensitive_word_service", reader)
	router.ServeHTTP(w, req)
	var resInfo map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resInfo)
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, float64(0), resInfo["code"].(float64))
}

func TestRemoveSensitiveWord(t *testing.T) {
	word := "测试数据"
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	_ = dty.AddWord(word)
	router := NewRouter()
	w := httptest.NewRecorder()
	reader := strings.NewReader("{\"word\": \""+word+"\"}")
	req, _ := http.NewRequest("DELETE", "/sensitive_word_service", reader)
	router.ServeHTTP(w, req)
	var resInfo map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resInfo)
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, float64(0), resInfo["code"].(float64))
}

func TestCheckSensitiveWord(t *testing.T) {
	word := "测试数据"
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	_ = dty.AddWord(word)
	defer func() {
		_ = dty.RemoveWord(word)
	}()
	router := NewRouter()
	w := httptest.NewRecorder()
	reader := strings.NewReader("{\"words\": [\""+word+"\"]}")
	req, _ := http.NewRequest("GET", "/sensitive_word_service", reader)
	router.ServeHTTP(w, req)
	var resInfo map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resInfo)
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, float64(0), resInfo["code"].(float64))
	assert.Equal(t, true, resInfo["info"].([]interface{})[0].(map[string]interface{})["sensitive"].(bool))


}