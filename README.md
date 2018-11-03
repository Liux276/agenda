# agenda

agenda是一个用于帮助用户创建会议，并且管理会议的命令行工具。

## cobra使用

Cobra既是一个用来创建强大的现代CLI命令行的golang库，也是一个生成程序应用和命令行文件的程序。此命令行工具基于cobra开发：

### cobra安装

直接执行以下命令，可能安装不成功:（因为cobra用到的一些依赖包被墙了）
> go get -v github.com/spf13/cobra/cobra

所以我们可以首先安装其依赖包：
在`$GOPATH/src/golang.org/x`目录下（如果没有，则自行创建）用`git clone`下载sys和text项目：
> git clone https://github.com/golang/sys
> git clone https://github.com/golang/text

然后执行`go get -v github.com/spf13/cobra/cobra`安装即可。
若成功安装，则在`$GOBIN`下会出现cobra可执行程序，如果没有配置`$GOBIN`,则可自行去`$GOPATH/bin`下寻找该文件。
然后在命令行中输入`cobra`:

```bash
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra [command]

Available Commands:
  add         Add a command to a Cobra Application
  help        Help about any command
  init        Initialize a Cobra Application
```

如果出现以上提示则表示安装成功。

### cobra使用

1. 生成agenda项目

```bash
$ cobra init agenda
```

2. 添加agenda工具命令

```bash
$ cobra add [命令名称]
```

3. cobra工作原理

```bash
agenda
  |─ cmd
  |  |─ register
  |  |─ login
  |  |─ 
  |  └─ ……
  |─ LICENSE
  └─ main.go
```

**main.go**文件如下：

```go
package main
import "agenda/cmd"
func main() {
  cmd.Execute()
}
```
主函数中调用了cmd.Execute()，此函数便启动了整个项目。

**cmd/root.go**:

```go
var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
```

root.go中定义了Execute()函数，在该函数中，启动了rootCmd.Execute()函数，内部实现中监听了所有命令。

除此之外，我们还可以定义flag来处理命令行参数：

```go
func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
  rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
  rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
  rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
  rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
  viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
  viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
  viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
  viper.SetDefault("license", "apache")
}
```

上述函数中定义了一些变量，这些变量我在agenda实现中并未接触到，接下来需要进一步学习。

## 命令介绍

有关agenda的命令，详见[commandIntro.md](./docs/commandIntro.md)。

## 文件结构介绍

```bash
agenda
├─ cmd
|   |─ ap.go
|   |─ cm.go
|   |─ login.go
|   |─ logout.go
|   |─ mtcancel.go
|   |─ mtquery.go
|   |─ mtquit.go
|   |─ register.go
|   |─ root.go
|   |─ rp.go
|   |─ ru.go
|   └─ userquery.go
├─ docs
|   |─ commandIntro.md
|   |─ entity.md
|   |─ models.md
├─ entity
|   |─ meetingOp.gp
|   |─ userInfoOp.go
|   |─ meetingOp_test.go
|   └─ userInfoOp.go
├─ log
|   └─ logFile.txt
├─ models
|   ├─ logger.go
|   ├─ meeting.go
|   └─ user.go
├─ storage
|   ├─ curUser.txt
|   ├─ meetings.json
|   └─ users.json
├─ LICENSE
├─ main.go
└─ README.md
```

## entity介绍

有关entity中函数定义，详见[entity.md](./docs/entity.md)。

## 测试结果

简单测试：（有些命令并没有测试到）

```bash
// 注册liuyh73
λ agenda register -uliuyh73 -p123 -P15976541234 -ea@163.com
Register liuyh73 successfully!

// 登录liuyh73
λ agenda login -uliuyh73 -p123
Login successfully

// 注册liuyt
λ agenda register -uliuyt -p234 -P15912345432 -eb@163.com
Register liuyt successfully!

// 注册liux
λ agenda register -uliux -p345 -P15978342332 -ec@163.com
Register liux successfully!

// 登录liuyt
λ agenda login -uliuyt -p234
liuyh73 has already in

// 创建会议test
λ agenda cm -ttest -oliuyh73 -pliuyt,liux -s="2018/11/1 10:20:00" -e="2018/11/1 11:00:00"
Create meeting: test successfully

// 船舰会议时间冲突
λ agenda cm -ttest -oliuyh73 -pliuyt -s="2018/11/1 10:30:00" -e="2018/11/1 11:10:00"
Some meetings of the sponsor("liuyh73")conflict with the meeting in terms of time.

// 创建会议liu未注册
λ agenda cm -ttest -oliuyh73 -pliuyt,liux,liu -s="2018/11/1 10:20:00" -e="2018/11/1 11:00:00"
The participator liu has not yet registered

// 移除参与者liux
λ agenda rp -ttest -pliux
rp called
Remove participators liux from meeting test success!

// 取消会议test
λ agenda mtcancel -ttest
mtcancel called
the meeting test are cancelled!

// 移除用户
λ agenda ru liuyh73
Remove user [liuyh73] successfully

// 退出用户
λ agenda logout
No user login

```

最终，meetings.json、curUser.txt文件为空，users.json存在以下两条记录:

```bash
[
	{"Username":"liuyt","Password":"234","Telephone":"15912345432","Email":"b@163.com"},
	{"Username":"liux","Password":"345","Telephone":"15978342332","Email":"c@163.com"}
]
```
