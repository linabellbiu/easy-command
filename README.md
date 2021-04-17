# easy-command

# 超级简单的获取子命令参数工具

### 通过`yaml`配置项来控制命令行参数获取，很方便删除添加修改子命令，好用又简单

- 配置`cmd.yaml`。注意要放到当前项目根目录下

```yaml
# 父命令结构
parent_cmd:
  # 父命令名
  - parent_name: "create"
    # 子命令结构
    child_cmd:
      # 子命令名
      - child_name: "a"
        # 默认值
        def: "create-a"
        # 说明
        usage: "is create-a"
      - child_name: "b"
        def: "create-b"
        usage: "is create-b"
  - parent_name: "update"
    child_cmd:
      - child_name: "c"
        def: "update-a"
        usage: "is update-c"
```

- 配置执行方法

```go

// 设置一个map，用来映射父命令对应的执行方法
var run = make(map[string]func (cmd.FlagValueMap))

func main() {
    
    run["create"] = HelloWord1
    run["update"] = HelloWord2
    
    //加载命令
    cmd.LoadCmd(run)
}

var HelloWord1 = func (s cmd.FlagValueMap) {
    fmt.Println("this is func: hello world1")
    //获取子命令值
    fmt.Printf("a: %s\n", s["a"].GetValueString())
    fmt.Printf("b: %s", s["b"].GetValueString())
}

var HelloWord2 = func (s cmd.FlagValueMap) {
    fmt.Println("this is func: hello world2")
    fmt.Printf("c: %s\n", s["c"].GetValueString())
}

```

- 执行
```
go run main.go create -a "haha"
```

- 结果

```
this is func: hello world1
a: haha
b: create-b
```

- 支持显示帮助 `-h`