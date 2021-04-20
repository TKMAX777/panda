package panda

import (
	"encoding/json"
	"time"
)

type Assignment struct {
	Access              string `json:"access"`
	AllPurposeItemText  string `json:"allPurposeItemText"`
	AllowPeerAssessment bool   `json:"allowPeerAssessment"`
	Author              string `json:"author"`
	AuthorLastModified  string `json:"authorLastModified"`
	CloseTimeString     string `json:"closeTimeString"`
	Content             string `json:"content"`
	Context             string `json:"context"`
	Creator             string `json:"creator"`
	CloseTime           time.Time
	DropDeadTime        time.Time
	TimeLastModified    time.Time
	OpenTime            time.Time
	TimeCreated         time.Time
	DueTime             time.Time
	DropDeadTimeString  string `json:"dropDeadTimeString"`
	DueTimeString       string `json:"dueTimeString"`
	GradeScale          string `json:"gradeScale"`
	GradeScaleMaxPoints string `json:"gradeScaleMaxPoints"`
	GradebookItemId     int    `json:"gradebookItemId"`
	GradebookItemName   string `json:"gradebookItemName"`
	ID                  string `json:"id"`
	Instructions        string `json:"instructions"`
	MaxGradePoint       string `json:"maxGradePoint"`
	ModelAnswerText     string `json:"modelAnswerText"`
	OpenTimeString      string `json:"openTimeString"`
	Position            int    `json:"position"`
	PrivateNoteText     string `json:"privateNoteText"`
	Section             string `json:"section"`
	Status              string `json:"status"`
	SubmissionType      string `json:"submissionType"`
	Title               string `json:"title"`
	AllowResubmission   bool   `json:"allowResubmission"`
	AnonymousGrading    bool   `json:"anonymousGrading"`
	Draft               bool   `json:"draft"`
	EntityReference     string `json:"entityReference"`
	EntityURL           string `json:"entityURL"`
	EntityId            string `json:"entityId"`
	EntityTitle         string `json:"entityTitle"`
}

func (a *Assignment) UnmarshalJSON(b []byte) error {
	type Alias Assignment
	var aux = &struct {
		DueTime          unixTime `json:"dueTime"`
		CloseTime        unixTime `json:"closeTime"`
		DropDeadTime     unixTime `json:"dropDeadTime"`
		TimeLastModified unixTime `json:"timeLastModified"`
		OpenTime         unixTime `json:"openTime"`
		TimeCreated      unixTime `json:"timeCreated"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	a.DueTime = time.Unix(aux.DueTime.EpochSecond, aux.DueTime.Nano)
	a.CloseTime = time.Unix(aux.CloseTime.EpochSecond, aux.CloseTime.Nano)
	a.DropDeadTime = time.Unix(aux.DropDeadTime.EpochSecond, aux.DropDeadTime.Nano)
	a.TimeLastModified = time.Unix(aux.TimeLastModified.EpochSecond, aux.TimeLastModified.Nano)
	a.OpenTime = time.Unix(aux.OpenTime.EpochSecond, aux.OpenTime.Nano)
	a.TimeCreated = time.Unix(aux.TimeCreated.EpochSecond, aux.TimeCreated.Nano)

	return nil
}

func (p *Handler) GetSiteAssignment(siteID string) (assignments []Assignment, err error) {
	p.sustainAuth()

	res, err := p.client.Get(BaseURI + "/direct/assignment/site/" + siteID + ".json")
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

func (p *Handler) GetAssignmentDetail(assignmentID string) (assignment Assignment, err error) {
	p.sustainAuth()

	res, err := p.client.Get(BaseURI + "/direct/assignment/item/" + assignmentID + ".json")
	if err != nil {
		return
	}
	defer res.Body.Close()

	var data Assignment

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	assignment = data

	return
}

func (p *Handler) GetAssignment() (assignments []Assignment, err error) {
	p.sustainAuth()

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
