# CMD-DESIGN
The command design of Agenda.

## Table of Contents

- [Supporting Commands](#supportingCommands)
    * [register](#register)
    * [login](#login)
    * [logout](#logout)
    * [listUsers](#listUsers)
    * [deleteUser](#deleteUser)
    * [createMeeting](#createMeeting)
    * [changeParticipants](#changeParticipants)
    * [listMeetings](#listMeetings)
    * [quitMeeting](#quitMeeting)
    * [cancelMeeting](#cancelMeeting)
    * [clearMeetings](#clearMeetings)


 <span id="supportingCommands"/>
## Supporting Commands
```
register            Register user.
login               User login.
logout              User logout.
listUsers           List all registered users.
deleteUser          Delete current login user.
createMeeting       Create a meeting.
changeParticipants  Change(add/delete) meeting Participants.
listMeetings        List all meetings.
quitMeeting         Quit a meeting.
cancelMeeting       Cancel a meeting.
clearMeetings       Clear all meetings.
```

<span id="register"/>
### register
Register user account with specific infomation.

Usage: `agenda register -u user -p password -e email -p phone`

```
Flags:
    -u, --user     string   Username
    -p, --password string   Password
    -e, --email    string   Email
    -p, --phone    string   Phone
```


<span id="login"/>
### login
User login.

Note: you need to login before using most of the functions in Agenda(except register command)

Usage: `agenda login -u user -p password -e email -p phone`

```
Flags:
    -u, --user     string   Username
    -p, --password string   Password
```


<span id="logout"/>
### logout
Logout the current login user.

Usage: `agenda logout`


<span id="listUsers"/>
### listUsers
List all registered users' information except for passwords.

Use these information to invite others to attend your meetings.

Usage: `agenda listUsers`


<span id="deleteUser"/>
### deleteUser
Delete current login user.

**Warning**: Once you do that, you can not find back the deleted user information (including the corresponding meeting information).

Usage: `agenda deleteUser`


<span id="createMeeting"/>
### createMeeting
Create a meeting.

At least one participant should be provided.

Usage: `agenda createMeeting -t title -p participants -s startTime -e endTime`

```
Flags:
    -t, --title        string   Title
    -p, --participants string   Participants
    -s, --start        string   start time  (format: yyyy-mm-dd)
    -e, --end          string   end time    (format: yyyy-mm-dd)
```


<span id="changeParticipants"/>
### changeParticipants
Change(add/delete) Participants of a meeting created by current login user.

Usage: `agenda changeParticipants -t title [-d|-a] participants`

```
Flags:
    -t, --title        string   Title
    -d, --delete       string   Participants that you intend to delete
    -a, --add          string   Participants that you intend to add
```


<span id="listMeetings"/>
### listMeetings
List all meetings attended by current login user.

Usage: `agenda listMeetings -s startTime -e endTime`

```
Flags:
    -s, --start        string   start time  (format: yyyy-mm-dd)
    -e, --end          string   end time    (format: yyyy-mm-dd)
```

<span id="quitMeeting"/>
### quitMeeting
Quit a meeting attended by current login user.

Usage: `agenda quitMeeting -t title`

```
Flags:
    -t, --title        string   Title of the meeting you intend to quit
```


<span id="cancelMeeting"/>
### cancelMeeting
Cancel a meeting created by current login user.

Usage: `agenda cancelMeeting -t title`

```
Flags:
    -t, --title        string   Title of the meeting you intend to cancel
```


<span id="clearMeetings"/>
### clearMeetings
clear all meetings created by current login user.

Warning: Once you do that, you can not find back the  information of meetings you created.

Usage: `agenda clearMeetings`
