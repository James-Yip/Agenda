package service

import (
	"fmt"

	"github.com/James-Yip/Agenda/entity"
	"github.com/James-Yip/Agenda/util"
)

func CreateMeeting(title string, participators_str string, startTime string, endTime string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
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
	for _, user := range participators {
		if user == CurUser {
			fmt.Println("login user can't be the particitor.")
			return
		}
		if !IsUserRegistered(user) {
			fmt.Println(user + " is not registered.")
			return
		}
		if !IsUserTimeVaild(user, startTime, endTime) {
			fmt.Println(user + " have not time.")
			return
		}
	}
	entity.AddMeeting(title, CurUser, participators, startTime, endTime)
}

func AddParticipators(title string, add_participators_str string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
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
		if !IsUserTimeVaild(user, meeting.StartTime, meeting.EndTime) {
			fmt.Println(user + " have not time.")
			return
		}
		changed_par = append(changed_par, user)
	}
	//_, meeting, err = entity.GetMeeting(title)
	//for _, i := range meeting.Participators {
	//	fmt.Println(i)
	//}
	entity.UpdateParticipators(title, changed_par)
	//fmt.Println(len(meeting.Participators))
	fmt.Println("change participator done.")
}

func DeleteParticipators(title string, delete_participators_str string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
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
		fmt.Println("delete particitors done(the meeting is del too).")
	} else {
		entity.UpdateParticipators(title, changed_par)
		fmt.Println("delete particitors done.")
	}
}

func ListMeetings(startTime, endTime string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
	}
	if !util.IsTimeValid(startTime) || !util.IsTimeValid(endTime) {
		fmt.Println("time is not valid")
		return
	}
	meetings := entity.GetMeetings()
	ss := util.Time2str(startTime)
	ee := util.Time2str(endTime)
	for _, meeting := range meetings {
		//fmt.Println(len(meeting.Participators))
		s := util.Time2str(meeting.StartTime)
		e := util.Time2str(meeting.EndTime)
		if (s > ss && s < ee) ||
			(e > ss && (e < ee || e == ee)) {
			if meeting.Sponsor == CurUser {
				PrintMeeting(meeting)
			}
			for _, str := range meeting.Participators {
				if CurUser == str {
					PrintMeeting(meeting)
					break
				}
			}
		}
	}
	fmt.Println("list meetings done.")
}

func QuitMeeting(title string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
	}
	_, meeting, err := entity.GetMeeting(title)
	if err == nil {
		if meeting.Sponsor == CurUser {
			entity.DeleteMeeting(title)
		} else {
			var change_par []string
			isUserPar := false
			for _, participator := range meeting.Participators {
				if participator != CurUser {
					change_par = append(change_par, participator)
				} else {
					isUserPar = true
				}
			}
			if len(change_par) == 0 {
				entity.DeleteMeeting(title)
				fmt.Println("quit done(meeting has been cancle).")
				return
			}
			if isUserPar {
				entity.UpdateParticipators(title, change_par)
				fmt.Println("quit done.")
			} else {
				fmt.Println("you don't attend the meeting.")
			}
		}
	} else {
		fmt.Println("can't find this meeting.")
	}
}

func CancelMeeting(title string) {
	if CurUser == "" {
		fmt.Println("please login first.")
		//return
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
		//return
	}
	meetings := entity.GetMeetings()
	for _, meeting := range meetings {
		if meeting.Sponsor == CurUser {
			entity.DeleteMeeting(meeting.Title)
		}
	}
	fmt.Println("clear meetings done.")
}

func PrintMeeting(meeting *entity.Meeting) {
	fmt.Println("title: " + meeting.Title)
	fmt.Println("sponsor: " + meeting.Sponsor)
	fmt.Println("starttime: " + meeting.StartTime)
	fmt.Println("endtime: " + meeting.EndTime)
	fmt.Printf("participators: ")
	for _, par := range meeting.Participators {
		fmt.Printf(par + " ")
	}
	fmt.Println(" ")
	fmt.Println(" ")
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

func IsUserTimeVaild(user string, start, end string) bool {
	meetings := entity.GetMeetings()
	ss := util.Time2str(start)
	ee := util.Time2str(end)
	for _, meeting := range meetings {
		s := util.Time2str(meeting.StartTime)
		e := util.Time2str(meeting.EndTime)
		if (s > ss && s < ee) ||
			(e > ss && (e < ee || e == ee)) {
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

func IsUserRegistered(user string) bool {
	return true
}
