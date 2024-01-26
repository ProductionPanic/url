package url

import (
	"fmt"
	"regexp"
)

var (
	protocol_pattern = regexp.MustCompile("^[a-z]+://")
	domain_pattern   = regexp.MustCompile("^[a-z]+://[^/]+")
	path_pattern     = regexp.MustCompile("^[a-z]+://[^/]+(/[^?#]*)?")
	query_pattern    = regexp.MustCompile("^[a-z]+://[^/]+(/[^?#]*)?\\?[^#]*")
	hash_pattern     = regexp.MustCompile("^[a-z]+://[^/]+(/[^?#]*)?(\\?[^#]*)?#.*")
)

type Url struct {
	protocol string
	domain   string
	path     string
	search   *UrlSearchParams
	hash     string
}

func New(url string) *Url {
	u := &Url{}
	u.parse(url)
	return u
}

func (u *Url) parse(url string) {
	u.protocol = protocol_pattern.FindString(url)
	u.domain = domain_pattern.FindString(url)
	u.path = path_pattern.FindString(url)
	u.hash = hash_pattern.FindString(url)
	query := query_pattern.FindString(url)
	if query != "" {
		u.search = NewUrlSearchParams(query)
	}
}

func (u *Url) String() string {
	return fmt.Sprintf("%s://%s%s%s%s", u.protocol, u.domain, u.path, u.search.String(), u.hash)
}

func (u *Url) Search() *UrlSearchParams {
	return u.search
}

func (u *Url) SetSearch(search *UrlSearchParams) {
	u.search = search
}

func (u *Url) Protocol() string {
	return u.protocol
}

func (u *Url) SetProtocol(protocol string) {
	u.protocol = protocol
}

func (u *Url) Domain() string {
	return u.domain
}

func (u *Url) SetDomain(domain string) {
	u.domain = domain
}

func (u *Url) Path() string {
	return u.path
}

func (u *Url) SetPath(path string) {
	u.path = path
}

func (u *Url) Hash() string {
	return u.hash
}

func (u *Url) SetHash(hash string) {
	u.hash = hash
}

func (u *Url) Copy() *Url {
	return &Url{
		protocol: u.protocol,
		domain:   u.domain,
		path:     u.path,
		search:   u.search,
		hash:     u.hash,
	}
}
