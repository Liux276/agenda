# entity功能函数介绍：
## meetingOp.go
-   ReadMeetingFromFile()
    ```bash
    /**
     * @arguments: nil
     * @return: []models.Meeting
     */
    ```
    此函数用于获取文件中所存放的所有会议。
    通过利用文件读操作，包括os、bufio、json-iterator/go等库的使用，我们可以解析meetings.json文件中存储的所有会议并作为models.Meeting切片来返回。

-   WriteMeetingToFile()
    ```bash
    /**
     * @arguments: []models.Meeting
     * @return: nil
     */
    ```
    此函数用于将当前列表中的更新后的所有会议重新写入meetings.json文件中。
    通过利用文件写操作，包括os、bufio、json-iterator/go等库的使用，我们可以将所有会议编码为json格式的字符串并存储到meetings.json文件中。

-   FetchMeetingsByName()
    ```bash
    /**
     * @arguments: name string
     * @return: []models.Meeting
     */
    ```
    此函数用于根据用户名字来获取该用户参与过的所有会议。
    通过利用文件读操作，包括os、bufio、json-iterator/go等库的使用，我们遍历整个会议文件，查询该用户参加/主持的所有会议作为modles.Meeting切片来返回。

-   RemoveParticipantsByName()
    ```bash
    /**
     * @arguments:  name string, meeting models.Meeting
     * @return: models.Meeting
     */
    ```
    此函数用于将一个参与人员从一个会议中移除。
    通过传入的meeting来查询该会议中是否包含该参与人员。如果包含，则将改参与人员删除；否则，不进行任何操作。该函数主要是一个中间过程函数，并不将出以后的meeting存储到文件中，在其他函数中可以调用该函数来完成相应的操作之后再存储。

## meetingOp.go测试
关于测试，我使用go语言自带的测试框架go test，相关内容不在展开叙述。
首先定义数据：
```go
var meetings = []models.Meeting{
	{
		Title:        "first",
		Originator:   "liuyh73",
		Participants: "liuyh74,liuyh75",
		StartTime:    "2018/11/1 10:00:00",
		EndTime:      "2018/11/1 10:30:00",
	},
	{
		Title:        "second",
		Originator:   "liu",
		Participants: "liuyh73,wang",
		StartTime:    "2018/10/13 21:00:00",
		EndTime:      "2018/10/13 21:30:00",
	},
}
```
-   ReadMeetingFromFile和WriteMeetingToFile测试
    在测试中，我将二者同时进行测试，首先将上面定义的两个会议写入文件中，然后再从文件中读取。如果读取出的数据与上述数据一致，则证明函数测试正确。
    ```go
    func TestReadMeetingFromFile_WriteMeetingToFile(t *testing.T) {
        WriteMeetingToFile(meetings)
        meetingsRead := ReadMeetingFromFile()
        if len(meetingsRead) == 2 && meetings[0] == meetingsRead[0] && meetings[1] == meetingsRead[1] {
            t.Log("ReadMeetingFromFile 和 WriteMeetingToFile 测试通过")
        } else {
            t.Error("ReadMeetingFromFile 或者 WriteMeetingToFile 测试失败")
        }
    }
    ```
-   FetchMeetingsByName测试
    此函数测试，我找寻两组不同的测试案例：一个为已经参加会议的liuyh73，另一个为未参加会议的liuyh76
    ```bash
    func TestFetchMeetingsByName(t *testing.T) {
        meetingsFetchedByName := FetchMeetingsByName("liuyh73")
        meetingsFetchedByName2 := FetchMeetingsByName("liuyh76")
        if len(meetingsFetchedByName) == 2 && len(meetingsFetchedByName2) == 0 {
            t.Log("FetchMeetingsByName 测试通过")
        } else {
            t.Error("FetchMeetingsByName 测试失败")
        }
    }
    ```
-   RemoveParticipantsByName测试
    此函数测试，我也找寻两组不同的测试案例：一个为参与会议meeting[0]的liuyh74，另一个为未参与会议meeting[1]（非主持人员）的liu:
    ```bash
    func TestRemoveParticipantsByName(t *testing.T) {
        meetingRemovedByName := RemoveParticipantsByName("liuyh74", meetings[0])
        meetingRemovedByName2 := RemoveParticipantsByName("liu", meetings[1])
        if (len(strings.Split(meetingRemovedByName.Participants, ",")) == 1 && meetingRemovedByName.Participants == "liuyh75") &&
            (len(strings.Split(meetingRemovedByName2.Participants, ",")) == 2 && meetingRemovedByName2.Participants == "liuyh73,wang") {
            t.Log("RemoveParticipantsByName 测试通过")
        } else {
            t.Error("RemoveParticipantsByName 测试失败")
        }
    }
    ```

## userInfoOp.go
-   ReadUserInfoFromFile()
    ```bash
    /**
     * @arguments: nil
     * @return: []models.User
     */
    ```
    此函数用于从users.json文件中读取所有用户信息。
    通过利用文件读操作，包括os、bufio、json-iterator/go等库的使用，我们遍历整个用户信息文件，获取所有用户的models.Meeting切片然后返回。
-   WriteUserInfoToFile()
    ```bash
    /**
     * @arguments: []models.User
     * @return: nil
     */
    ```
    此函数用于将当前列表中的更新后的所有用户信息重新写入users.json文件中。
    通过利用文件写操作，包括os、bufio、json-iterator/go等库的使用，我们可以将所有用户信息编码为json格式的字符串并存储到users.json文件中。

-   SaveCurUserInfo()
    ```bash
    /**
     * @arguments: loginUser models.User
     * @return: nil
     */
    ```
    此函数用于将当前登陆的用户信息存储到curUser.txt文件中，方便登陆用户信息的存储。

-   ClearCurUserInfo()
    ```bash
    /**
     * @arguments: nil
     * @return: nil
     */
    ```
    当登陆用户登出的时候，我们利用os库Truncate函数来将登录用户信息从curUser.txt文件中删除。

-   IsLoggedIn()
    ```bash
    /**
     * @arguments: nil
     * @return: bool, models.User
     */
    ```
    此函数判断当前是否已经已经有用户登录，并且返回登录用户信息。
    我们可以利用此函数来加一些限定，因为未登录的用户不能进行cm、mtcancel等操作。

-   IsUser()
    ```bash
    /**
     * @arguments: name string
     * @return: bool
     */
    ```
    此函数用于判端当前用户名是否为已注册的用户，调用ReadUserInfoFromFile并加以判断即可。可以用于在创建、删除会议时判断用户是否存在；或者注册用户时判断该用户名是否已经被注册等。

-   RemoveUser()
    ```bash
    /**
     * @arguments: name string
     * @return: nil
     */
    ```
    此函数用于移除用处，主要是方便ru操作。调用ReadUserInfoFromFile获取用户信息，加以处理后再调用WriteUserInfoToFile更新用户信息即可。


## userInfoOp.go测试
首先，定义数据：
```go
var users = []models.User{
	{
		Username:  "liuyh73",
		Password:  "123",
		Telephone: "12364579823",
		Email:     "a@163.com",
	},
	{
		Username:  "liuyh74",
		Password:  "123",
		Telephone: "12364579823",
		Email:     "a@163.com",
	},
}

```
-   ReadUserInfoFromFile和WriteUserInfoToFile测试
    在测试中，我将二者同时进行测试，首先将上面定义的两个用户写入文件中，然后再从文件中读取。如果读取出的数据与上述数据一致，则证明函数测试正确。
    具体代码与测试会议的代码基本一致，不再赘述。
-   SaveCurUserInfo测试
    假设"liuyh73"为登录用户，将其保存在curUser.txt中，然后再从中读取出数据，若该数据等于users[0]，则测试正确。
    ```go
    func TestSaveCurUserInfo(t *testing.T) {
        // 假设登陆用户为liuyh73
        SaveCurUserInfo(users[0])
        file, err := os.OpenFile(models.ExecPath+"github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
        defer file.Close()

        if err != nil {
            panic(err)
        }
        var loginUser models.User
        reader := bufio.NewReader(file)
        data, _ := reader.ReadBytes('\n')

        jsoniter.Unmarshal(data, &loginUser)
        if loginUser == users[0] {
            t.Log("SaveCurUserInfo 测试通过")
        } else {
            t.Error("SaveCurUserInfo 测试失败")
        }
    }
    ```
-   IsLoggedIn测试
    判断登陆的用户是否为liuyh73即可。
    ```go
    func TestIsLoggedIn(t *testing.T) {
        login, loginUser := IsLoggedIn()
        if login && loginUser == users[0] {
            t.Log("IsLoggedIn 测试通过")
        } else {
            t.Error("IsLoggedIn 测试失败")
        }
    }
    ```
-   ClearCurUserInfo测试
    将当前登陆的用户从curUser.txt中移除，重新读取该文件，若该文件为空，则测试正确。
    ```go
    func TestClearCurUserInfo(t *testing.T) {
        ClearCurUserInfo()
        file, err := os.OpenFile(models.ExecPath+"github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
        defer file.Close()

        if err != nil {
            panic(err)
        }
        reader := bufio.NewReader(file)
        data, err := reader.ReadBytes('\n')
        if string(data) == "" && err == io.EOF {
            t.Log("ClearCurUserInfo 测试通过")
        } else {
            t.Error("ClearCurUserInfo 测试失败")
        }
    }
    ```
-   IsUser测试
    找寻两组数据：一个是已注册用户liuyh73，另一个为未注册用户liuyh75.
    ```go
    func TestIsUser(t *testing.T) {
        isUser := IsUser("liuyh73")
        isNotUser := IsUser("liuyh75")
        if isUser && !isNotUser {
            t.Log("IsUser 测试通过")
        } else {
            t.Error("IsUser 测试失败")
        }
    }
    ```
-   RemoveUser测试
    删除用户，然后读取文件获取用户数量，判断是否删除成功。
    ```go
    func TestRemoveUser(t *testing.T) {
        RemoveUser("liuyh73")
        users := ReadUserInfoFromFile()
        RemoveUser("liuyh74")
        users2 := ReadUserInfoFromFile()
        if len(users) == 1 && len(users2) == 0 {
            t.Log("RemoveUser 测试通过")
        } else {
            t.Error("RemoveUser 测试失败")
        }
    }
    ```
## 测试结果
```bash
=== RUN   TestReadMeetingFromFile_WriteMeetingToFile
--- PASS: TestReadMeetingFromFile_WriteMeetingToFile (0.00s)
    meetingOp_test.go:31: ReadMeetingFromFile 和 WriteMeetingToFile 测试通过
=== RUN   TestFetchMeetingsByName
--- PASS: TestFetchMeetingsByName (0.00s)
    meetingOp_test.go:41: FetchMeetingsByName 测试通过
=== RUN   TestRemoveParticipantsByName
--- PASS: TestRemoveParticipantsByName (0.00s)
    meetingOp_test.go:52: RemoveParticipantsByName 测试通过
=== RUN   TestReadUserInfoFromFile_WriteUserInfoToFile
--- PASS: TestReadUserInfoFromFile_WriteUserInfoToFile (0.00s)
    userInfoOp_test.go:33: ReadUserInfoFromFile 和 WriteUserInfoToFile 测试通过
=== RUN   TestSaveCurUserInfo
--- PASS: TestSaveCurUserInfo (0.00s)
    userInfoOp_test.go:54: SaveCurUserInfo 测试通过
=== RUN   TestIsLoggedIn
--- PASS: TestIsLoggedIn (0.00s)
    userInfoOp_test.go:63: IsLoggedIn 测试通过
=== RUN   TestClearCurUserInfo
--- PASS: TestClearCurUserInfo (0.00s)
    userInfoOp_test.go:80: ClearCurUserInfo 测试通过
=== RUN   TestIsUser
--- PASS: TestIsUser (0.00s)
    userInfoOp_test.go:90: IsUser 测试通过
=== RUN   TestRemoveUser
--- PASS: TestRemoveUser (0.00s)
    userInfoOp_test.go:102: RemoveUser 测试通过
PASS
ok      github.com/sysu-615/agenda/entity       0.480s
```