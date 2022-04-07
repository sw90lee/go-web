package forms

// 오류 발생시 오류유형 정의
type errors map[string][]string

// 오류 메세지를 추가하고 특정항목과 연결된 오류를 추가함
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// 오류 메세지를 확인하고 오류가 있으면 오류를 반환함
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
