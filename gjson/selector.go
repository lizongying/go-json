package gjson

import (
	"errors"
	"github.com/tidwall/gjson"
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

func (s *Selector) OneSelector(path string) *Selector {
	s.str = gjson.Get(s.str, path).Raw
	return s
}

func (s *Selector) ManySelector(path string) (selectors []*Selector) {
	json := gjson.Get(s.str, path)
	if json.IsArray() {
		for _, i := range json.Array() {
			selectors = append(selectors, &Selector{str: i.Raw})
		}
	} else {
		selectors = append(selectors, &Selector{str: json.Raw})
	}
	return
}

func (s *Selector) String() string {
	return s.str
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
