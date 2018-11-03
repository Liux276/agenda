# models模块介绍

## logger.go

此模块添加日志服务，定义日志格式及日志存储文件。

```go
func init() {
    var logHandler *os.File
    logHandler, err := os.OpenFile(ExecPath+"github.com/sysu-615/agenda/log/logFile.txt", os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    Logger = log.New(logHandler, "", log.LstdFlags)
}
```
之后，import此模块之后，就可以使用Logger来输出日志。
除此之外，由于go语言执行命令式的路径为执行命令时的目录，并不是可执行文件的目录。所以我在此文件中定义了执行路径，使得agenda可以在任何目录下正常执行。

```go
var ExecPath string
func init() {
    // for循环部分可以忽略，这部分是为了go test测试使用
    for _, arg := range os.Args {
        if strings.Split(arg, "=")[0] == "go_test" {
            ExecPath = strings.Replace(strings.Split(arg, "=")[1], "\\", "/", -1)
            if ExecPath[len(ExecPath)-1] != '/' {
                ExecPath += "/"
            }
            return
        }
    }
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))

    // the current path will be $GOPATH/bin
    // so I here truncate the bin and add src to get $GOPATH/src
    ExecPath = path[:index-3] + "src/"
    ExecPath = strings.Replace(ExecPath, "\\", "/", -1)
}
```

上述函数中利用exec、filepath等库来获取agenda.exe所在的路径，由于我们是通过go install安装的此项目，所以该可执行文件的路径固定。经过简单的处理之后定位到$GOPATH中，然后再找到所需的文件即可。

## meeting.go

定义Meeting，包括会议主题、会议主持者、会议参与者、开始时间、结束时间。

```go
type Meeting struct {
    Title           string
    Originator      string
    Participants    string
    StartTime       string
    EndTime         string
}
```

## user.go

定义User，包括用户名、密码、手机号、邮箱。

```go
type User struct {
    Username    string
    Password    string
    Telephone   string
    Email       string
}
```

## partake.go

此数据结构的设计是为了方便查询用户参与的所有会议，类似于关系型数据库中的关系表，但实际实现中并未使用到。

```go
type Partake struct {
    participator    string
    MeetingTitle    string
    Initiate        bool
}
```