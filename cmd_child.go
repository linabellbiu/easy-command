package cmd

import (
	"flag"
)

type Child struct {
	name  string
	usage string
	Value *string
	def   string
}

func (c *Child) ChildCmdPares(set *flag.FlagSet) {
	c.Value = set.String(c.name, c.def, c.usage)
}

func (c *Child) GetValueString() string {
	return *c.Value
}

func (c *Child) GetName() string {
	return c.name
}

func (c *Child) SetName(name string) {
	c.name = name
}

func (c *Child) SetDef(def string) {
	c.def = def
}
func (c *Child) SetUsage(usage string) {
	c.usage = usage
}

func newChildCmd() childCmdFunc {
	return func(s []childCmd) []ChildCommandInterFace {
		var newChildList []ChildCommandInterFace
		for _, child := range s {
			newChild := &Child{}
			newChild.name = child.ChildName
			newChild.usage = child.Usage
			newChild.def = child.Def
			newChildList = append(newChildList, newChild)
		}
		return newChildList
	}
}
