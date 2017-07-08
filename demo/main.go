package main

import (
	"fmt"

	"github.com/gunsluo/kagome/tokenizer"
)

func main() {
	tokenize, err := newTokenizer("", "")
	if err != nil {
		panic(err)
	}

	input := "関西国際空港"
	mode := tokenizer.Search //tokenizer.Extended
	tokens := tokenize.Analyze(input, mode)
	for _, tok := range tokens {
		if tok.ID == tokenizer.BosEosID {
			continue
		}
		fs := tok.Features()

		fmt.Println(tok, fs)
	}
}

func newDic(typ string) tokenizer.Dic {
	switch typ {
	case "ipa":
		return tokenizer.SysDicIPA()
	case "uni":
		return tokenizer.SysDicUni()
	default:
		return tokenizer.SysDic()
	}
}

func newUserDic(typ string) (tokenizer.UserDic, error) {
	if typ == "" {
		return tokenizer.UserDic{}, nil
	}

	return tokenizer.NewUserDic(typ)
}

func newTokenizer(sysdic, usrdic string) (tokenizer.Tokenizer, error) {
	dic := newDic(sysdic)
	udic, err := newUserDic(usrdic)
	if err != nil {
		return tokenizer.Tokenizer{}, err
	}

	tokenize := tokenizer.NewWithDic(dic)
	tokenize.SetUserDic(udic)

	return tokenize, nil
}
