package types

import "fmt"





type Invariant func(ctx Context) (string, bool)


type Invariants []Invariant


type InvariantRegistry interface {
	RegisterRoute(moduleName, route string, invar Invariant)
}


func FormatInvariant(module, name, msg string) string {
	return fmt.Sprintf("%s: %s invariant\n%s\n", module, name, msg)
}
