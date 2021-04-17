package cmd

import (
	"flag"
	"fmt"
)

type Parent struct {
	flag     *flag.FlagSet
	childCmd map[string]ChildCommandInterFace
	f        func(FlagValueMap)
}

func (s *Parent) CmdPares(args []string) {
	if err := s.flag.Parse(args); err != nil {
		panic(fmt.Sprintf("cmd pares error:%s", err.Error()))
	}
}

func (s *Parent) Usage() {
	s.flag.Usage()
}

//加载子命令
//如果没有子命令就不加载
func (s *Parent) LoadChildCmd(childCmdList ...ChildCommandInterFace) {
	for _, childCmd := range childCmdList {
		//解析子命令
		childCmd.ChildCmdPares(s.flag)
		s.childCmd[childCmd.GetName()] = childCmd
	}
}

func (s *Parent) SetFunc(f func(FlagValueMap)) {
	s.f = f
}

//执行程序
func (s *Parent) Exec() {
	s.f(s.childCmd)
}

func newParentCmd(newChildFunc childCmdFunc) parentCmdFunc {
	return func(par parentCmd) {
		s := &Parent{
			flag:     flag.NewFlagSet(par.ParentName, flag.ExitOnError),
			childCmd: make(map[string]ChildCommandInterFace),
		}
		cli[par.ParentName] = s
		//初始化一个子命令
		s.LoadChildCmd(newChildFunc(par.Child)...)
	}
}
