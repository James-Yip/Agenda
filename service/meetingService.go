package service

import (
	"agenda/entity"
	"agenda/utils"
	"fmt"
)

func CreateMeeting(title string, participators_str string, startTime string, endTime string) {
	var isMeetingValid bool
	var meetings entity.Meetings
	var participators []string
	isMeetingValid := true
	meetings := entity.GetMeetings()
	participators := utils.str2slice(participators_str)
	if title == "" || len(participators) <= 0 {
		isMeetingValid := false
	}
	if !utils.isTimeValid(startTime) || !utils.isTimeValid(endTime) {
		isMeetingValid := false
	}
	for i, meeting := range meetings {
		if meeting.Title == title {
			isMeetingValid := false
		}
	}
	if isMeetingValid {
		entity.AddMeeting(title.participators_str, startTime, endTime)
	}
}

func AddParticipators(title string, add_participators_str string) {
	var meeting entity.Meeting
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		var change_par []string
		var add_par []string
		var isRepeat bool
		change_par := meeting.Participators
		add_par := utils.str2slice(add_participators_str)
		for i, par := range add_par {
			isRepeat := false
			for j, par2 := range meeting.Participators {
				if par == par2 {
					isRepeat := true
				}
			}
			if !isRepeat {
				change_par := append(change_par, par)
			}
		}
		entity.UpdateParticipators(title, change_par)
	} else {
		fmt.Println(err.Error())
	}
}

func DeleteParticipators(title string, delete_participators_str string) {
	var meeting entity.Meeting
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		var change_par []string
		var del_par []string
		var isRepeat bool
		change_par := meeting.Participators
		del_par := utils.str2slice(delete_participators_str)
		for i, par := range meeting.Participators {
			isRepeat := false
			for j, par2 := range del_par {
				if par == par2 {
					isRepeat := true
				}
			}
			if !isRepeat {
				change_par := append(change_par, par)
			}
		}
		entity.UpdateParticipators(title, change_par)
	} else {
		fmt.Println(err.Error())
	}
}

func ListMeetings(startTime, endTime string) {
	if !utils.isTimeValid(startTime) || !utils.isTimeValid(endTime) {
		fmt.Println("time is not valid")
		return
	}
	var meetings entity.Meetings
	meetings := entity.GetMeetings()
	for i, meeting := range meetings {
		if (startTime >= meeting.StartTime && startTime < meeting.EndTime) ||
			(endTime <= meeting.EndTime && endTime > meeting.StartTime) {
			if meeting.Sponsor == username {
				fmt.Println(meeting.Title)
			}
			for j, str := range meeting.Participators {
				if username == str {
					fmt.Println(meeting.Title)
				}
			}
		}
	}
}

func QuitMeeting(title string) {
	var meeting entity.Meeting
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		if meeting.Sponsor == username {
			entity.DeleteMeeting(title)
		} else {
			var change_par []string
			var isUserPar bool
			isUserPar := false
			for i, participator := range meeting.Participators {
				if participator != username {
					change_par := append(change_par, participator)
				} else {
					isUserPar := true
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
	var meeting entity.Meeting
	_, meeting, err := entity.GetMeeting(title)
	if err != nil {
		if meeting.Sponsor == username {
			entity.DeleteMeeting(title)
		} else {
			//sponsor is not current user
			fmt.Println("this meeting's sponsor is not current user")
		}
	} else {
		fmt.Println(err.Error())
	}
}

func ClearMeetings() {
	var meetings entity.Meetings
	meetings := entity.GetMeetings()
	for index, meeting := range meetings {
		if meeting.Sponsor == username {
			CancelMeeting(meeting.Title)
		}
	}
}
