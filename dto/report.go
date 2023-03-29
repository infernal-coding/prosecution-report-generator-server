package dto

type ReportData struct {
	Header  ReportHeader  `json:"header"`
	Content ReportContent `json:"content"`
	Footer  ReportFooter  `json:"footer"`
}

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

type MatchClass struct {
	Federation bool `json:"federation"`
	Private    bool `json:"private"`
	Other      bool `json:"other"`
}

type ReportRecipients struct {
	GSO bool `json:"gso"`
	KSO bool `json:"kso"`
	BSO bool `json:"bso"`
	VSO bool `json:"vso"`
}
