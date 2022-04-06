package forms

import (
	"net/http"
	"net/url"
)

// Form 구조체 생성
type Form struct {
	url.Values
	Errors errors
}

// Form을 통해 데이터를 새로 만듬
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// field에 값이 존재하는지 안하는지 검사하는 함수
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
