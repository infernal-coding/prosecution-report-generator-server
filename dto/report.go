package dto

import "errors"

// ReportData contains the data for a report.
type ReportData struct {
	Header  ReportHeader  `json:"header"`
	Content ReportContent `json:"content"`
	Footer  ReportFooter  `json:"footer"`
}

// ReportHeader contains the data for the report header.
// This specifies the match the report is associated to.
type ReportHeader struct {
	Team            string     `json:"team"`
	Home            string     `json:"home"`
	Away            string     `json:"away"`
	Date            string     `json:"date"`
	Place           string     `json:"place"`
	HomeScoreHalf   uint8      `json:"homeScoreHalf"`
	AwayScoreHalf   uint8      `json:"awayScoreHalf"`
	HomeScore       uint8      `json:"homeScore"`
	AwayScore       uint8      `json:"awayScore"`
	MatchIdentifier string     `json:"matchIdentifier"`
	MatchClass      MatchClass `json:"matchClass"`
}

// ReportContent contains the main content of the report.
// This specifies the incident that happened at the match.
type ReportContent struct {
	Subject       string `json:"subject"`
	PassportId    string `json:"passportId"`
	Who           string `json:"who"`
	When          string `json:"when"`
	How           string `json:"how"`
	What          string `json:"what"`
	Against       string `json:"against"`
	Where         string `json:"where"`
	WhereBall     string `json:"whereBall"`
	WhereReferee  string `json:"whereReferee"`
	AlreadyWarned string `json:"alreadyWarned"`
	WasProvoked   string `json:"wasProvoked"`
	CouldPlayOn   string `json:"couldPlayOn"`
	HowContinued  string `json:"howContinued"`
	WhatEffect    string `json:"whatEffect"`
	BehaviorAfter string `json:"behaviorAfter"`
	Other         string `json:"other"`
}

// ReportFooter contains the data for the footer of the report.
// This specifies the authoring referee and the recipients of the report.
type ReportFooter struct {
	Name       string           `json:"name"`
	Street     string           `json:"street"`
	Zip        string           `json:"zip"`
	Place      string           `json:"place"`
	Group      string           `json:"group"`
	Phone      string           `json:"phone"`
	Mail       string           `json:"mail"`
	Recipients ReportRecipients `json:"recipients"`
}

// MatchClass contains the available options of the match class.
// A match is either classified as federation, private or other.
type MatchClass struct {
	Federation bool `json:"federation"`
	Private    bool `json:"private"`
	Other      bool `json:"other"`
}

// Validate validates that exactly one option for the match class is set to true.
func (matchClass *MatchClass) Validate() error {
	federation := matchClass.Federation
	private := matchClass.Private
	other := matchClass.Other

	if (federation && !private && !other) || (!federation && private && !other) || (!federation && !private && other) {
		return nil
	}
	return errors.New("exactly one match class must be chosen")
}

// ReportRecipients contains the possible recipients of a report.
// The report can be sent to one or more recipients.
type ReportRecipients struct {
	GSO bool `json:"gso"`
	KSO bool `json:"kso"`
	BSO bool `json:"bso"`
	VSO bool `json:"vso"`
}

// Validate validates that at least one recipient is set to true.
func (reportRecipients *ReportRecipients) Validate() error {
	gso := reportRecipients.GSO
	kso := reportRecipients.KSO
	bso := reportRecipients.BSO
	vso := reportRecipients.VSO

	if !gso && !kso && !bso && !vso {
		return errors.New("at least one recipient must be chosen")
	}
	return nil
}
