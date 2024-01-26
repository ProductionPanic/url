package url

import "strings"

type UrlSearchParams struct {
	params map[string]string
}

func NewUrlSearchParams(query string) *UrlSearchParams {
	u := &UrlSearchParams{}
	u.parse(query)
	return u
}

func (u *UrlSearchParams) parse(query string) {
	if len(query) == 0 {
		return
	}
	if query[0] == '?' {
		query = query[1:]
	}
	u.params = make(map[string]string)
	for _, param := range strings.Split(query, "&") {
		kv := strings.Split(param, "=")
		if len(kv) == 2 {
			u.params[kv[0]] = kv[1]
		}
	}
}

func (u *UrlSearchParams) Get(key string) string {
	return u.params[key]
}

func (u *UrlSearchParams) Set(key string, value string) {
	u.params[key] = value
}

func (u *UrlSearchParams) Delete(key string) {
	delete(u.params, key)
}

func (u *UrlSearchParams) String() string {
	if len(u.params) == 0 {
		return ""
	}
	result := "?"
	for k, v := range u.params {
		result += k + "=" + v + "&"
	}
	return result[:len(result)-1]
}

func (u *UrlSearchParams) Has(key string) bool {
	_, ok := u.params[key]
	return ok
}

func (u *UrlSearchParams) Keys() []string {
	keys := make([]string, len(u.params))
	i := 0
	for k := range u.params {
		keys[i] = k
		i++
	}
	return keys
}
