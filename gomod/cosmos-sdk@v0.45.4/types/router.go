package types

import (
	"regexp"
	"strings"
)

var (


	IsAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString



	IsAlphaLower = regexp.MustCompile(`^[a-z]+$`).MatchString



	IsAlphaUpper = regexp.MustCompile(`^[A-Z]+$`).MatchString



	IsAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString



	IsNumeric = regexp.MustCompile(`^[0-9]+$`).MatchString
)


type Router interface {
	AddRoute(r Route) Router
	Route(ctx Context, path string) Handler
}

type Route struct {
	path    string
	handler Handler
}


func NewRoute(p string, h Handler) Route {
	return Route{path: strings.TrimSpace(p), handler: h}
}


func (r Route) Path() string {
	return r.path
}


func (r Route) Handler() Handler {
	return r.handler
}


func (r Route) Empty() bool {
	return r.handler == nil || r.path == ""
}


type QueryRouter interface {
	AddRoute(r string, h Querier) QueryRouter
	Route(path string) Querier
}
