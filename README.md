# Agenda
[![BuildStatus](https://travis-ci.org/James-Yip/Agenda.svg?branch=v0.5)](https://travis-ci.org/James-Yip/Agenda)

Agenda is a CLI meeting manager program based on cobra.

This application is a tool to support various operations on meetings
including user register, meeting creation & query, etc.

## Collaborators
15331376 叶政
15331395 张力越
15331377 伊少波

## Usage
```
Usage:
  agenda [command]

Available Commands:
  cancelMeeting       Cancel a meeting
  changeParticipators Change(add/delete) meeting participators.
  clearMeetings       Clear all meetings
  createMeeting       Create a meeting
  deleteUser          Delete current login user
  help                Help about any command
  listMeetings        List all meetings
  listUsers           List all registered users
  login               User login
  logout              User logout
  quitMeeting         Quit a meeting
  register            Register user

Flags:
  -h, --help   help for agenda

Use "agenda [command] --help" for more information about a command.

```

See `cmd-design.md` for information about each command.


## Examples

### register
```
$ ./agenda register -u james -p 123456 -e james@qq.com -t 13623887454
用户名可使用，正在注册......
注册成功！当前用户是：james
$ ./agenda register
请检查输入信息是否正确
$ ./agenda register -u james -p 123456 -e james@qq.com -t 13623887454
此用户名已被注册
```

### login
```
$ ./agenda login -u james -p 123456
登录成功！当前用户是：james
$ ./agenda login -u bob -p 123456
用户名错误！登录失败！
$ ./agenda login -u james -p 123
密码错误！登录失败！
```

### logout
```
$ ./agenda logout
退出登录！
$ ./agenda logout
您未登录！
```

### listUsers
```
$ ./agenda listUsers
所有用户信息：
编号  用户名  电子邮箱  电话号码
0 | james | james@qq.com | 13623887454
```


### deleteUser
```
$ ./agenda deleteUser
删除用户james成功
退出登陆
$ ./agenda deleteUser
请先登录！！！
```

### createMeeting
```
$ ./agenda createMeeting -t singleDog -p "bob" -s 2017-11-11/11:11 -e 2017-11-11/11:22

$ ./agenda createMeeting -t coding -p "alice" -s 2017-11-11/11:00 -e 2017-11-11/12:00

$ ./agenda createMeeting -t coding -p "alice" -s 2017-11-12/11:00 -e 2017-11-12/12:00
```

### changeParticipators
```
$ ./agenda changeParticipators -t singleDog -a alice

$ ./agenda changeParticipators -t singleDog -d bob
```

### listMeetings
```
$ ./agenda listMeetings -s 2017-11-11/11:11 -e 2017-11-11/11:22
```

### quitMeeting
```
$ ./agenda quitMeeting -t singleDog
```

### cancelMeeting
```
$ ./agenda cancelMeeting -t coding

$ ./agenda listMeetings -s 2017-11-11/11:10 -e 2017-11-13/11:30
```

### clearMeetings
```
$ ./agenda clearMeetings

$ ./agenda listMeetings -s 2017-11-11/11:10 -e 2017-11-13/11:30
```
