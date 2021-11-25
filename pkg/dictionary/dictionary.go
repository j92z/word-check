package dictionary

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sensitive_words_check/constant"
	"sensitive_words_check/model/word_model"
	"strings"
	"sync"
)

var dictionaryInstance *Dictionary

var fileMutex sync.Mutex

type Dictionary struct {
	Words     []string
	path      string
	version   int
	storeType constant.DictionaryStoreType
}

func OpenDictionary(store constant.DictionaryStoreType, resource ...string) *Dictionary {
	if dictionaryInstance == nil || (store == constant.FileStore && len(resource) > 0 && dictionaryInstance.path != resource[0]) || dictionaryInstance.storeType != store {
		if store == constant.FileStore {
			var filePath string
			if len(resource) == 0 {
				_, filename, _, _ := runtime.Caller(0)
				filePath = path.Join(path.Dir(filename), "../../resource/sensitive_words.txt")
			} else {
				filePath = resource[0]
			}
			file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			buf := bufio.NewReader(file)
			dictionary := Dictionary{
				path: filePath,
				storeType: constant.FileStore,
			}
			for {
				line, _, err := buf.ReadLine()
				if err != nil {
					if err == io.EOF {
						if string(line) != "" {
							dictionary.Words = append(dictionary.Words, string(line))
						}
						fmt.Println("dictionary read done!")
						break
					} else {
						panic(err)
					}
				}
				if string(line) != "" {
					fileMutex.Lock()
					dictionary.Words = append(dictionary.Words, string(line))
					fileMutex.Unlock()
				}
			}
			dictionaryInstance = &dictionary
		} else {
			sensitiveWords := word_model.GetAll()
			var words []string
			for _, w := range sensitiveWords {
				words = append(words, w.Word)
			}
			dictionaryInstance = &Dictionary{
				Words: words,
				storeType: constant.MysqlStore,
			}
		}
	}

	return dictionaryInstance
}

func (d *Dictionary) AddWord(word string) error {
	if len(word) == 0 {
		return errors.New("word can't empty")
	}
	if d.storeType == constant.FileStore {
		flag := 0
		for _, v := range d.Words {
			if v == word {
				flag = 1
				break
			}
		}
		if flag == 0 {
			file, err := os.OpenFile(d.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()
			d.Words = append(d.Words, word)
			fileMutex.Lock()
			_, err = io.WriteString(file, strings.Join(d.Words, "\n"))
			d.incVersion()
			fileMutex.Unlock()
			return err
		}
	} else {
		wordInfo := word_model.FindWord(word)
		if wordInfo.ID <= 0 {
			_ = word_model.AddWord(word)
			d.Words = append(d.Words, word)
			d.incVersion()
		}
	}
	return nil
}

func (d *Dictionary) RemoveWord(word string) error {
	if len(word) == 0 {
		return errors.New("word can't empty")
	}
	if d.storeType == constant.FileStore {
		flag := 0
		for i, v := range d.Words {
			if v == word {
				flag = 1
				d.Words = append(d.Words[:i], d.Words[i+1:]...)
				break
			}
		}
		if flag == 1 {
			file, err := os.OpenFile(d.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()
			fileMutex.Lock()
			_, err = io.WriteString(file, strings.Join(d.Words, "\n"))
			d.incVersion()
			fileMutex.Unlock()
			return err
		}
	} else {
		wordInfo := word_model.FindWord(word)
		if wordInfo.ID > 0 {
			_ = word_model.RemoveWord(word)
			for i, v := range d.Words {
				if v == word {
					d.Words = append(d.Words[:i], d.Words[i+1:]...)
					break
				}
			}
			d.incVersion()
		}
	}
	return nil
}

func (d *Dictionary) incVersion() {
	d.version++
}

func (d *Dictionary) Version() int {
	return d.version
}
