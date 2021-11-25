package dfa

import (
	"sensitive_words_check/config"
	"sensitive_words_check/pkg/dictionary"
	"strings"
	"sync"
)

const (
	defaultInvalidWorlds = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,¥,·…"
	defaultReplaceStr    = "****"
)

type DFA struct {
	l                 sync.Mutex
	trie              *Trie
	replaceStr        string
	invalidWords      map[string]struct{}
	dictionaryVersion int
}

var dfa *DFA

func GetDFA() *DFA {
	if dfa == nil {
		dfa = newDFA()
	}
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	if dfa.trie.size == 0 || dfa.dictionaryVersion < dty.Version() {
		dfa.AddBadWords(dty.Words)
		dfa.dictionaryVersion = dty.Version()
	}
	return dfa
}

func newDFA() *DFA {
	f := &DFA{
		trie:         NewTrie(),
		replaceStr:   defaultReplaceStr,
		invalidWords: make(map[string]struct{}),
	}
	for _, s := range defaultInvalidWorlds {
		f.invalidWords[string(s)] = struct{}{}
	}
	return f
}

func (f *DFA) AddBadWords(words []string) {
	f.l.Lock()
	defer f.l.Unlock()
	if len(words) > 0 {
		for _, s := range words {
			f.trie.Insert(s)
		}
	}
}

func (f *DFA) SetInvalidChar(chars string) {
	f.l.Lock()
	defer f.l.Unlock()
	f.invalidWords = make(map[string]struct{})
	for _, s := range chars {
		f.invalidWords[string(s)] = struct{}{}
	}
}

func (f *DFA) SetReplaceStr(str string) {
	f.l.Lock()
	defer f.l.Unlock()

	f.replaceStr = str
}

func (f *DFA) Check(txt string) ([]string, []string, bool) {
	_, found, target, b := f.check(txt, false)
	return found, target, b
}

func (f *DFA) CheckAndReplace(txt string) (string, []string, []string, bool) {
	return f.check(txt, true)
}

func (f *DFA) FilterInvalidChar(txt ...string) []string {
	res := make([]string, 0, len(txt))
	for _, s := range txt {
		str := []rune(s)
		for i, c := range str {
			if _, ok := f.invalidWords[string(c)]; ok {
				str = append(str[:i], str[i+1:]...)
			}
		}
		res = append(res, string(str))
	}
	return res
}

type hitNodeStruct struct {
	word  string
	node  map[rune]*Node
	start int
	end   int
}

func (f *DFA) check(txt string, replace bool) (dist string, found []string, target []string, b bool) {
	var (
		str = []rune(txt)
		//ok         bool
		hitNodes []hitNodeStruct
		result   string
	)
	target = make([]string, 0, 0)
	f.l.Lock()
	defer f.l.Unlock()

	for i, val := range str {
		if _, ok := f.invalidWords[string(val)]; ok {
			continue
		}
		for j := 0; j < len(hitNodes); j++ {
			if node, ok := hitNodes[j].node[val]; ok {
				hitNodes[j].word += node.Value
				if !node.IsEnd {
					hitNodes[j].node = node.Child
				} else {
					target = append(target, hitNodes[j].word)
					found = append(found, string(str[hitNodes[j].start:i+1]))
					if replace {
						result = strings.Replace(result, string(str[hitNodes[j].start:i+1]), f.replaceStr, 1)
						if result == "" {
							result = strings.Replace(txt, string(str[hitNodes[j].start:i+1]), f.replaceStr, 1)
						}
					}
					if j+1 >= len(hitNodes) {
						hitNodes = hitNodes[:j]
					} else {
						hitNodes = append(hitNodes[:j], hitNodes[j+1:]...)
					}
					j-- //删除节点导致元素长度下降的问题
				}
			} else {
				if j+1 >= len(hitNodes) {
					hitNodes = hitNodes[:j]
				} else {
					hitNodes = append(hitNodes[:j], hitNodes[j+1:]...)
				}
				j-- //删除节点导致元素长度下降的问题
			}
		}

		var node = f.trie.Child(string(val))
		if node != nil {
			if !node.IsEnd {
				hitNodes = append(hitNodes, hitNodeStruct{
					word:  node.Value,
					node:  node.Child,
					start: i,
				})
			} else {
				target = append(target, node.Value)
				//tmp = ""
				found = append(found, string(str[i:i+1]))
				if replace {
					result = strings.Replace(result, string(str[i:i+1]), f.replaceStr, 1)
					if result == "" {
						result = strings.Replace(txt, string(str[i:i+1]), f.replaceStr, 1)
					}
				}
			}
		}
	}
	b = len(found) > 0
	return
}
