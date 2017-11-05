package entity

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Meeting struct {
	Title         string
	Sponsor       string
	StartTime     string
	EndTime       string
	Participators []string
}

type Meetings []*Meeting

var meetings Meetings

func AddMeeting(title string, sponsor string, participators []string, startTime string, endTime string) {
	meeting := &Meeting{
		Title:         title,
		Sponsor:       sponsor,
		StartTime:     startTime,
		EndTime:       endTime,
		Participators: participators}
	meetings = append(meetings, meeting)
	meetingWriteToFile()
}

func DeleteMeeting(title string) error {
	index, _, err := GetMeeting(title)
	if err != nil {
		return errors.New("delete meeting fail: " + err.Error())
	} else {
		meetings = append(meetings[:index], meetings[index+1:]...)
		meetingWriteToFile()
	}
	return nil
}

func GetMeeting(title string) (int, *Meeting, error) {
	for index, meeting := range meetings {
		if meeting.Title == title {
			return index, meeting, nil
		}
	}
	return -1, nil, errors.New("can not get meeting by " + title)
}

func GetMeetings() Meetings {
	return meetings
}

func UpdateParticipators(title string, participators []string) error {
	index, _, err := GetMeeting(title)
	if err != nil {
		return errors.New("update participators fail: " + err.Error())
	} else {
		meetings[index].Participators = participators
	}
	meetingWriteToFile()
	return nil
}

func meetingEncode(meeting *Meeting) ([]byte, error) {
	encodedMeeting, err := json.Marshal(meeting)
	if err != nil {
		return nil, errors.New("meeting encode fail: " + err.Error())
	}
	return encodedMeeting, nil
}

func meetingDecode(encodedMeeting []byte) (*Meeting, error) {
	var meeting Meeting
	err := json.Unmarshal(encodedMeeting, &meeting)
	if err != nil {
		return nil, errors.New("meeting decode fail: " + err.Error())
	}
	return &meeting, nil
}

func meetingWriteToFile() error {
	fout, err := os.Create("data/meetingInfo")
	if err != nil {
		return errors.New("meeting write to file fail: \n->" + err.Error())
	}
	defer fout.Close()

	for _, meeting := range meetings {
		encodedMeeting, err := meetingEncode(meeting)
		if err != nil {
			return errors.New("meeting write to file fail: \n->" + err.Error())
		}
		fmt.Fprintf(fout, "%s\n", encodedMeeting)
	}

	return nil
}

func meetingReadFromFile() error {
	fin, err := os.Open("data/meetingInfo")
	if err != nil {
		return errors.New("meeting read from file fail: \n->" + err.Error())
	}
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	for scanner.Scan() {
		token := scanner.Text()
		// fmt.Printf("%v\n", scanner.Text())
		meeting, err := meetingDecode([]byte(token))
		if err != nil {
			return errors.New("meeting read from file fail: \n->" + err.Error())
		}
		meetings = append(meetings, meeting)
	}

	return nil
}

func init() {
	err := meetingReadFromFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
