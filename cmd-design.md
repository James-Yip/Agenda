# CMD-DESIGN
The command design of Agenda.

## Supporting Commands
```
register            Register user.
login               User login.
logout              User logout.
listUsers           List all registered users.
deleteUser          Delete current login user.
createMeeting       Create a meeting.
changeParticipators  Change(add/delete) meeting participators.
listMeetings        List all meetings.
quitMeeting         Quit a meeting.
cancelMeeting       Cancel a meeting.
clearMeetings       Clear all meetings.
```


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


### login
User login.

Note: you need to login before using most of the functions in Agenda(except register command)

Usage: `agenda login -u user -p password -e email -p phone`

```
Flags:
    -u, --user     string   Username
    -p, --password string   Password
```


### logout
Logout the current login user.

Usage: `agenda logout`


### listUsers
List all registered users' information except for passwords.

Use these information to invite others to attend your meetings.

Usage: `agenda listUsers`


### deleteUser
Delete current login user.

**Warning**: Once you do that, you can not find back the deleted user information (including the corresponding meeting information).

Usage: `agenda deleteUser`


### createMeeting
Create a meeting.

At least one participator should be provided.

Usage: `agenda createMeeting -t title -p participators -s startTime -e endTime`

```
Flags:
    -t, --title         string   Title
    -p, --participators string   participators
    -s, --start         string   start time  (format: yyyy-mm-dd/hh:mm)
    -e, --end           string   end time    (format: yyyy-mm-dd/hh:mm)
```


### changeParticipators
Change(add/delete) participators of a meeting created by current login user.

Usage: `agenda changeParticipators -t title [-d|-a] participators`

```
Flags:
    -t, --title        string   Title
    -d, --delete       string   participators that you intend to delete
    -a, --add          string   participators that you intend to add
```


### listMeetings
List all meetings attended by current login user.

Usage: `agenda listMeetings -s startTime -e endTime`

```
Flags:
    -s, --start        string   start time  (format: yyyy-mm-dd/hh:mm)
    -e, --end          string   end time    (format: yyyy-mm-dd/hh:mm)
```


### quitMeeting
Quit a meeting attended by current login user.

Usage: `agenda quitMeeting -t title`

```
Flags:
    -t, --title        string   Title of the meeting you intend to quit
```


### cancelMeeting
Cancel a meeting created by current login user.

Usage: `agenda cancelMeeting -t title`

```
Flags:
    -t, --title        string   Title of the meeting you intend to cancel
```


### clearMeetings
clear all meetings created by current login user.

Warning: Once you do that, you can not find back the  information of meetings you created.

Usage: `agenda clearMeetings`
