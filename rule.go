package puck

import (
	"crypto/md5"
	"fmt"
)

type Rule struct {
	Name, Command, Description, Defaults, Params string
	Connection, Inputs, Dependencies, Variables  interface{}
}

func (r *Rule) Hash() [16]byte {
	return md5.Sum([]byte(fmt.Sprintf(r.Name, r.Command, r.Description, r.Defaults, r.Params)))
}
