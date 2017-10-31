package entity

type Meeting struct {
	Title         string
	Sponsor       string
	StartTime     string
	EndTime       string
	Participators []string
}

type Meetings []*Meeting

func AddMeeting(title string, participators_str string, startTime string, endTime string) {

}

func DeleteMeeting(title string) {

}

func GetMeeting(title string) {

}

func UpdateParticipators(participators []string) {

}
