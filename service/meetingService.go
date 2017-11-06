package service

import (
	"fmt"

	"github.com/James-Yip/Agenda/entity"
	"github.com/James-Yip/Agenda/util"
)

func CreateMeeting(title string, participators_str string, startTime string, endTime string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	meetings := entity.GetMeetings()
	participators := util.Str2slice(participators_str)
	if title == "" {
		fmt.Println("title is not normative.")
		return
	}
	for _, meeting := range meetings {
		if meeting.Title == title {
			fmt.Println("title is repeated.")
			return
		}
	}
	if !util.IsTimeValid(startTime) || !util.IsTimeValid(endTime) {
		fmt.Println("time is not normative.")
		return
	}
	if util.Time2str(startTime) > util.Time2str(endTime) ||
		util.Time2str(startTime) == util.Time2str(endTime) {
		fmt.Println("startTime should < endTime")
		return
	}
	if len(participators) <= 0 {
		fmt.Println("participators should >= 1")
		return
	}
	for _, user := range participators {
		if user == CurUser {
			fmt.Println("login user can't be the participator.")
			return
		}
		if !IsUserRegistered(user) {
			fmt.Println(user + " is not registered.")
			return
		}
		if !IsUserHaveTime(user, startTime, endTime) {
			fmt.Println(user + " have not time.")
			return
		}
	}
	entity.AddMeeting(title, CurUser, participators, startTime, endTime)
	fmt.Println("create meeting success.")
	_, meeting, _ := entity.GetMeeting(title)
	PrintMeeting(meeting)
}

func AddParticipators(title string, add_participators_str string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		fmt.Println("can't find this meeting.")
		return
	}
	if meeting.Sponsor != CurUser {
		fmt.Println("the meeting's sponsor is't login user.")
		return
	}
	add_par := util.Str2slice(add_participators_str)
	changed_par := meeting.Participators
	for _, user := range add_par {
		if !IsUserRegistered(user) {
			fmt.Println(user + " is not registered")
			return
		}
		if !IsUserHaveTime(user, meeting.StartTime, meeting.EndTime) {
			fmt.Println(user + " have not time.")
			return
		}
		changed_par = append(changed_par, user)
	}
	entity.UpdateParticipators(title, changed_par)
	fmt.Println("add participators success.")
	_, meeting, _ = entity.GetMeeting(title)
	PrintMeeting(meeting)
}

func DeleteParticipators(title string, delete_participators_str string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		fmt.Println("can't find this meeting.")
		return
	}
	if meeting.Sponsor != CurUser {
		fmt.Println("the meeting's sponsor is't login user.")
		return
	}
	del_par := util.Str2slice(delete_participators_str)
	changed_par := meeting.Participators
	for _, user := range del_par {
		if !IsUserRegistered(user) {
			fmt.Println(user + " is not registered")
			return
		}
		if !IsUserAttend(meeting, user) {
			fmt.Println(user + " don't attend the meeting.")
			return
		}
		changed_par = Remove(changed_par, user)
	}
	if len(changed_par) == 0 {
		entity.DeleteMeeting(meeting.Title)
		fmt.Println("delete participators done(the meeting is del because no participators).")
	} else {
		entity.UpdateParticipators(title, changed_par)
		fmt.Println("delete participators success.")
		_, meeting, _ = entity.GetMeeting(title)
		PrintMeeting(meeting)
	}
}

func ListMeetings(startTime, endTime string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	if !util.IsTimeValid(startTime) || !util.IsTimeValid(endTime) {
		fmt.Println("time is not valid")
		return
	}
	meetings := entity.GetMeetings()
	ss := util.Time2str(startTime)
	ee := util.Time2str(endTime)
	meetingnum := 0
	fmt.Println("Your meeting information：")
	fmt.Println("title  sponsor  startTime  endTime  participators")
	for _, meeting := range meetings {
		//fmt.Println(len(meeting.Participators))
		s := util.Time2str(meeting.StartTime)
		e := util.Time2str(meeting.EndTime)
		if (s > ss && s < ee) ||
			(e > ss && (e < ee || e == ee)) ||
			(s < ss && e > e) {
			if meeting.Sponsor == CurUser {
				meetingnum++
				PrintMeeting(meeting)
			}
			for _, str := range meeting.Participators {
				if CurUser == str {
					meetingnum++
					PrintMeeting(meeting)
					break
				}
			}
		}
	}
	fmt.Println("Total meeting amount: ", meetingnum)
}

func QuitMeeting(title string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	_, meeting, err := entity.GetMeeting(title)
	if err == nil {
		if IsUserAttend(meeting, CurUser) {
			if len(meeting.Participators) == 1 {
				entity.DeleteMeeting(title)
				fmt.Println("quit done(meeting has been cancle beacuse no participators).")
			} else {
				changed_par := Remove(meeting.Participators, CurUser)
				entity.UpdateParticipators(title, changed_par)
				fmt.Println("quit done.")
			}
		} else {
			fmt.Println("you don't attend the meeting.")
		}
	} else {
		fmt.Println("can't find this meeting.")
	}
}

func CancelMeeting(title string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	_, meeting, err := entity.GetMeeting(title)
	if err == nil {
		if meeting.Sponsor == CurUser {
			entity.DeleteMeeting(title)
			fmt.Println("cancel done.")
		} else {
			fmt.Println("this meeting's sponsor is not login user.")
		}
	} else {
		fmt.Println("can't find this meeting.")
	}
}

func ClearMeetings() {
	if CurUser == "" {
		fmt.Println("please login first.")
		return
	}
	var del_title []string
	meetings := entity.GetMeetings()
	for _, meeting := range meetings {
		if meeting.Sponsor == CurUser {
			del_title = append(del_title, meeting.Title)
		}
	}
	for _, title := range del_title {
		entity.DeleteMeeting(title)
	}
	fmt.Println("clear total", len(del_title), "meetings done.")
}

func PrintMeeting(meeting *entity.Meeting) {
	fmt.Printf("%s | %s | %s | %s | %v\n",
		meeting.Title,
		meeting.Sponsor,
		meeting.StartTime,
		meeting.EndTime,
		meeting.Participators)
}

func Remove(users []string, user string) []string {
	var newusers []string
	for _, i := range users {
		if i != user {
			newusers = append(newusers, i)
		}
	}
	return newusers
}

func IsUserAttend(meeting *entity.Meeting, user string) bool {
	for _, i := range meeting.Participators {
		if i == user {
			return true
		}
	}
	return false
}

func IsUserHaveTime(user string, start, end string) bool {
	meetings := entity.GetMeetings()
	ss := util.Time2str(start)
	ee := util.Time2str(end)
	for _, meeting := range meetings {
		s := util.Time2str(meeting.StartTime)
		e := util.Time2str(meeting.EndTime)
		if (s > ss && s < ee) ||
			(e > ss && (e < ee || e == ee)) ||
			(s < ss && e > e) {
			if user == meeting.Sponsor {
				return false
			}
			for _, i := range meeting.Participators {
				if i == user {
					return false
				}
			}
		}
	}
	return true
}
