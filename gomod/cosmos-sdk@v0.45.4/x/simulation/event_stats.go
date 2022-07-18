package simulation

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)



type EventStats map[string]map[string]map[string]int


func NewEventStats() EventStats {
	return make(EventStats)
}


func (es EventStats) Tally(route, op, evResult string) {
	_, ok := es[route]
	if !ok {
		es[route] = make(map[string]map[string]int)
	}

	_, ok = es[route][op]
	if !ok {
		es[route][op] = make(map[string]int)
	}

	es[route][op][evResult]++
}


func (es EventStats) Print(w io.Writer) {
	obj, err := json.MarshalIndent(es, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(obj))
}


func (es EventStats) ExportJSON(path string) {
	bz, err := json.MarshalIndent(es, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, bz, 0600)
	if err != nil {
		panic(err)
	}
}
