package cmd

import (
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type CommandInterFace interface {
	CmdPares(args []string)
	Usage()
	LoadChildCmd(childCmd ...ChildCommandInterFace)
	Exec()
	SetFunc(func(FlagValueMap))
}

//子命令接口
type ChildCommandInterFace interface {
	ChildCmdPares(*flag.FlagSet)
	GetName() string
	SetName(string)
	SetDef(string)
	SetUsage(string)
	GetValueString() string
}

var cli cliMap = make(map[string]CommandInterFace)

var cmdYaml = &commandYaml{}

func parse() {
	yamlFile, err := ioutil.ReadFile("./cmd.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	if err = yaml.Unmarshal(yamlFile, &cmdYaml); err != nil {
		log.Fatalln(err.Error())
	}
}

//加载命令
//只提供这个加载入口
func LoadCmd(m map[string]func(FlagValueMap)) {
	parse()

	registerCmd(newParentCmd(newChildCmd()))

	run(m)
}

func run(m map[string]func(FlagValueMap)) {
	if len(os.Args) < 2 {
		help()
	}

	if c := findCmd(os.Args[1]); c != nil {
		c.CmdPares(os.Args[2:])
		if f, ok := m[os.Args[1]]; ok {
			c.SetFunc(f)
		} else {
			panic(errors.New(fmt.Sprintf("not have this `%s` command function", os.Args[1])))
		}
		//执行命令
		c.Exec()
	} else {
		help()
	}
}

func help() {
	for _, c := range cli {
		c.Usage()
	}
	os.Exit(0)
}

func findCmd(cmdName string) CommandInterFace {
	if v, ok := cli[cmdName]; ok {
		return v
	}
	return nil
}

//注册命令
func registerCmd(cmdFunc parentCmdFunc) {
	for _, par := range cmdYaml.Parent {
		cmdFunc(par)
	}
}
