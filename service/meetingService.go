package service

import (
	"github.com/James-Yip/Agenda/entity"
	"github.com/James-Yip/Agenda/util"
	"fmt"
)

func CreateMeeting(title string, participators_str string, startTime string, endTime string) {
	isMeetingValid := true
	meetings := entity.GetMeetings()
	participators := util.Str2slice(participators_str)
	if title == "" || len(participators) <= 0 {
		isMeetingValid = false
	}
	if !util.IsTimeValid(startTime) || !util.IsTimeValid(endTime) {
		isMeetingValid = false
	}
	for _, meeting := range meetings {
		if meeting.Title == title {
			isMeetingValid = false
		}
	}
	if isMeetingValid {
		entity.AddMeeting(title, CurUser, participators, startTime, endTime)
	}
}

func AddParticipators(title string, add_participators_str string) {
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		change_par := meeting.Participators
		add_par := util.Str2slice(add_participators_str)
		for _, par := range add_par {
			isRepeat := false
			for _, par2 := range meeting.Participators {
				if par == par2 {
					isRepeat = true
				}
			}
			if !isRepeat {
				change_par = append(change_par, par)
			}
		}
		entity.UpdateParticipators(title, change_par)
	} else {
		fmt.Println(err.Error())
	}
}

func DeleteParticipators(title string, delete_participators_str string) {
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		change_par := meeting.Participators
		del_par := util.Str2slice(delete_participators_str)
		for _, par := range meeting.Participators {
			isRepeat := false
			for _, par2 := range del_par {
				if par == par2 {
					isRepeat = true
				}
			}
			if !isRepeat {
				change_par = append(change_par, par)
			}
		}
		entity.UpdateParticipators(title, change_par)
	} else {
		fmt.Println(err.Error())
	}
}

func ListMeetings(startTime, endTime string) {
	if !util.IsTimeValid(startTime) || !util.IsTimeValid(endTime) {
		fmt.Println("time is not valid")
		return
	}
	meetings := entity.GetMeetings()
	for _, meeting := range meetings {
		if (startTime >= meeting.StartTime && startTime < meeting.EndTime) ||
			(endTime <= meeting.EndTime && endTime > meeting.StartTime) {
			if meeting.Sponsor == CurUser {
				fmt.Println(meeting.Title)
			}
			for _, str := range meeting.Participators {
				if CurUser == str {
					fmt.Println(meeting.Title)
				}
			}
		}
	}
}

func QuitMeeting(title string) {
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
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
			if isUserPar {
				entity.UpdateParticipators(title, change_par)
			}
		}
	} else {
		fmt.Println(err.Error())
	}
}

func CancelMeeting(title string) {
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		if meeting.Sponsor == CurUser {
			entity.DeleteMeeting(title)
		} else {
			fmt.Println("this meeting's sponsor is not current user")
		}
	} else {
		fmt.Println(err.Error())
	}
}

func ClearMeetings() {
	meetings := entity.GetMeetings()
	for _, meeting := range meetings {
		if meeting.Sponsor == CurUser {
			CancelMeeting(meeting.Title)
		}
	}
}
