package gjson

import (
	"errors"
	"github.com/tidwall/gjson"
	"regexp"
)

type Selector struct {
	str string
}

func (s *Selector) One(path string) *Result {
	return NewResult(gjson.Get(s.str, path).Raw)
}

func (s *Selector) Many(path string) (results []*Result) {
	json := gjson.Get(s.str, path)
	if json.IsArray() {
		for _, i := range json.Array() {
			results = append(results, NewResult(i.Raw))
		}
	} else {
		results = append(results, NewResult(json.Raw))
	}
	return
}
func (s *Selector) ManyString(str string) []string {
	return regexp.MustCompile(str).FindAllString(s.str, -1)
}

// NewSelectorFromBytes selector from bytes
func NewSelectorFromBytes(bs []byte) (selector *Selector, err error) {
	if len(bs) == 0 {
		err = errors.New("bs is empty")
		return
	}
	selector = &Selector{
		str: string(bs),
	}
	return
}

// NewSelectorFromStr selector from str
func NewSelectorFromStr(str string) (selector *Selector, err error) {
	if str == "" {
		err = errors.New("str is empty")
		return
	}
	selector = &Selector{
		str: str,
	}
	return
}
