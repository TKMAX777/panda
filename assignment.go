package panda

import (
	"encoding/json"
)

type Assignment struct {
	Access              string `json:"access"`
	AllPurposeItemText  string `json:"allPurposeItemText"`
	AllowPeerAssessment bool   `json:"allowPeerAssessment"`
	Author              string `json:"author"`
	AuthorLastModified  string `json:"authorLastModified"`
	CloseTime           Time   `json:"closeTime"`
	CloseTimeString     string `json:"closeTimeString"`
	Content             string `json:"content"`
	Context             string `json:"context"`
	Creator             string `json:"creator"`
	DropDeadTime        struct {
	} `json:"dropDeadTime"`
	DropDeadTimeString  string `json:"dropDeadTimeString"`
	DueTime             Time   `json:"dueTime"`
	DueTimeString       string `json:"dueTimeString"`
	GradeScale          string `json:"gradeScale"`
	GradeScaleMaxPoints string `json:"gradeScaleMaxPoints"`
	GradebookItemId     int    `json:"gradebookItemId"`
	GradebookItemName   string `json:"gradebookItemName"`
	ID                  string `json:"id"`
	Instructions        string `json:"instructions"`
	MaxGradePoint       string `json:"maxGradePoint"`
	ModelAnswerText     string `json:"modelAnswerText"`
	OpenTime            Time
	OpenTimeString      string `json:"openTimeString"`
	Position            int    `json:"position"`
	PrivateNoteText     string `json:"privateNoteText"`
	Section             string `json:"section"`
	Status              string `json:"status"`
	SubmissionType      string `json:"submissionType"`
	TimeCreated         Time   `json:"timeCreated"`
	TimeLastModified    Time   `json:"timeLastModified"`
	Title               string `json:"title"`
	AllowResubmission   bool   `json:"allowResubmission"`
	AnonymousGrading    bool   `json:"anonymousGrading"`
	Draft               bool   `json:"draft"`
	EntityReference     string `json:"entityReference"`
	EntityURL           string `json:"entityURL"`
	EntityId            string `json:"entityId"`
	EntityTitle         string `json:"entityTitle"`
}

func (p *Handler) AssignmentGet() (assignments []Assignment, err error) {
	res, err := p.client.Get(BaseURI + "/direct/assignment/my.json")
	if err != nil {
		return
	}
	defer res.Body.Close()

	var data Data

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	assignments = data.AssignmentCollection

	return
}
