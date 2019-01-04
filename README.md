# Agenda
[![Build Status](https://travis-ci.org/James-Yip/Agenda.svg?branch=master)](https://travis-ci.org/James-Yip/Agenda)

Agenda is a CLI meeting manager program based on cobra.

This application is a tool to support various operations on meetings
including user register, meeting creation & query, etc.

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
Username available，registering......
Register: success
Current login user: james

$ ./agenda register
You have been registered！Please don't register repeatedly！
Current login user: james

$ ./agenda register -u james -p 123456 -e james@qq.com -t 13623887454
Username repeats
```

### login
```
$ ./agenda login -u james -p 123456
Login: success
Current login user: james

$ ./agenda login -u jame -p 123456
Login: fail
User unfound

$ ./agenda login -u james -p 123
Login: fail
Password incorrect
```

### logout
```
$ ./agenda logout
Logout: success

$ ./agenda logout
Logout: fail
You aren't logged-in
```

### listUsers
```
$ ./agenda listUsers
All User Information:
index  username  email  phone
0 | bob | bob@qq.com | 15044569825
1 | james | james@qq.com | 13623887454
2 | alice | alice@qq.com | 13765845784
```


### deleteUser
```
$ ./agenda deleteUser
DeleteUser(james): success
Logout

$ ./agenda deleteUser
Please login first
```

### createMeeting
```
$ ./agenda createMeeting -t singleDog -p "bob alice" -s 2017-11-11/11:11 -e 2017-11-11/11:22
create meeting success.
singleDog | james | 2017-11-11/11:11 | 2017-11-11/11:22 | [bob]

$ ./agenda createMeeting -t coding -p "alice" -s 2017-11-11/11:00 -e 2017-11-11/12:00
alice have not time.

$ ./agenda createMeeting -t coding -p "alice" -s 2017-11-12/11:00 -e 2017-11-12/12:00
create meeting success.
coding | james | 2017-11-12/11:00 | 2017-11-12/12:00 | [alice]
```

### changeParticipators
```
$ ./agenda changeParticipators -t coding -a bob
add participators success.
coding | james | 2017-11-12/11:00 | 2017-11-12/12:00 | [alice bob]

$ ./agenda changeParticipators -t singleDog -d alice
delete participators success.
singleDog | james | 2017-11-11/11:11 | 2017-11-11/11:22 | [bob]
```

### listMeetings
```
$ ./agenda listMeetings -s 2017-11-11/11:11 -e 2017-11-11/11:22
Your meeting information：
title  sponsor  startTime  endTime  participators
singleDog | james | 2017-11-11/11:11 | 2017-11-11/11:22 | [bob]

Total meeting amount:  1
```

### quitMeeting
```
$ ./agenda quitMeeting -t coding
quit done.
```

### cancelMeeting
```
$ ./agenda cancelMeeting -t coding
cancel done.

$ ./agenda listMeetings -s 2017-11-11/11:10 -e 2017-11-13/11:30
Your meeting information：
title  sponsor  startTime  endTime  participators
singleDog | james | 2017-11-11/11:11 | 2017-11-11/11:22 | [bob]
Total meeting amount:  1

```

### clearMeetings
```
$ ./agenda clearMeetings
clear total 1 meetings done.

$ ./agenda listMeetings -s 2017-11-11/11:10 -e 2017-11-13/11:30
Your meeting information：
title  sponsor  startTime  endTime  participators
Total meeting amount:  0
```
