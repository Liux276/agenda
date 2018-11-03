package entity

import (
	"strings"
	"testing"

	"github.com/sysu-615/agenda/models"
)

var meetings = []models.Meeting{
	{
		Title:        "first",
		Originator:   "liuyh73",
		Participants: "liuyh74,liuyh75",
		StartTime:    "2018/11/1 10:00:00",
		EndTime:      "2018/11/1 10:30:00",
	},
	{
		Title:        "second",
		Originator:   "liu",
		Participants: "liuyh73,wang",
		StartTime:    "2018/10/13 21:00:00",
		EndTime:      "2018/10/13 21:30:00",
	},
}

func TestReadMeetingFromFile_WriteMeetingToFile(t *testing.T) {
	WriteMeetingToFile(meetings)
	meetingsRead := ReadMeetingFromFile()
	if len(meetingsRead) == 2 && meetings[0] == meetingsRead[0] && meetings[1] == meetingsRead[1] {
		t.Log("ReadMeetingFromFile 和 WriteMeetingToFile 测试通过")
	} else {
		t.Error("ReadMeetingFromFile 或者 WriteMeetingToFile 测试失败")
	}
}

func TestFetchMeetingsByName(t *testing.T) {
	meetingsFetchedByName := FetchMeetingsByName("liuyh73")
	meetingsFetchedByName2 := FetchMeetingsByName("liuyh76")
	if len(meetingsFetchedByName) == 2 && len(meetingsFetchedByName2) == 0 {
		t.Log("FetchMeetingsByName 测试通过")
	} else {
		t.Error("FetchMeetingsByName 测试失败")
	}
}

func TestRemoveParticipantsByName(t *testing.T) {
	meetingRemovedByName := RemoveParticipantsByName("liuyh74", meetings[0])
	meetingRemovedByName2 := RemoveParticipantsByName("liu", meetings[1])
	if (len(strings.Split(meetingRemovedByName.Participants, ",")) == 1 && meetingRemovedByName.Participants == "liuyh75") &&
		(len(strings.Split(meetingRemovedByName2.Participants, ",")) == 2 && meetingRemovedByName2.Participants == "liuyh73,wang") {
		t.Log("RemoveParticipantsByName 测试通过")
	} else {
		t.Error("RemoveParticipantsByName 测试失败")
	}
}
