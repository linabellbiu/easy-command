package cmd

type commandYaml struct {
	Parent []parentCmd `yaml:"parent_cmd"`
}

type parentCmd struct {
	ParentName string     `yaml:"parent_name"`
	Child      []childCmd `yaml:"child_cmd"`
}

type childCmd struct {
	ChildName string `yaml:"child_name"`
	Def       string `yaml:"def"`
	Usage     string `yaml:"usage"`
}

type parentCmdFunc func(par parentCmd)

type childCmdFunc func(s []childCmd) []ChildCommandInterFace

//存放全局的父命令
type cliMap map[string]CommandInterFace

type FlagValueMap map[string]ChildCommandInterFace
