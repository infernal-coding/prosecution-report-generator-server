package dto

import "testing"

func TestMatchClass_Validate(t *testing.T) {
	tests := []struct {
		name    string
		tested  MatchClass
		wantErr bool
	}{
		{name: "only federation", tested: MatchClass{Federation: true, Private: false, Other: false}, wantErr: false},
		{name: "only private", tested: MatchClass{Federation: false, Private: true, Other: false}, wantErr: false},
		{name: "only other", tested: MatchClass{Federation: false, Private: false, Other: true}, wantErr: false},
		{name: "none", tested: MatchClass{Federation: false, Private: false, Other: false}, wantErr: true},
		{name: "federation and private", tested: MatchClass{Federation: true, Private: true, Other: false}, wantErr: true},
		{name: "federation and other", tested: MatchClass{Federation: true, Private: false, Other: true}, wantErr: true},
		{name: "private and other", tested: MatchClass{Federation: false, Private: true, Other: true}, wantErr: true},
		{name: "all", tested: MatchClass{Federation: true, Private: true, Other: true}, wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.tested.Validate(); (err != nil) != test.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestReportRecipients_Validate(t *testing.T) {
	tests := []struct {
		name    string
		tested  ReportRecipients
		wantErr bool
	}{
		{name: "none", tested: ReportRecipients{GSO: false, KSO: false, BSO: false, VSO: false}, wantErr: true},
		{name: "only gso", tested: ReportRecipients{GSO: true, KSO: false, BSO: false, VSO: false}, wantErr: false},
		{name: "only kso", tested: ReportRecipients{GSO: false, KSO: true, BSO: false, VSO: false}, wantErr: false},
		{name: "only bso", tested: ReportRecipients{GSO: false, KSO: false, BSO: true, VSO: false}, wantErr: false},
		{name: "only vso", tested: ReportRecipients{GSO: false, KSO: false, BSO: false, VSO: true}, wantErr: false},
		{name: "gso and kso", tested: ReportRecipients{GSO: true, KSO: true, BSO: false, VSO: false}, wantErr: false},
		{name: "gso and bso", tested: ReportRecipients{GSO: true, KSO: false, BSO: true, VSO: false}, wantErr: false},
		{name: "gso and vso", tested: ReportRecipients{GSO: true, KSO: false, BSO: false, VSO: true}, wantErr: false},
		{name: "kso and bso", tested: ReportRecipients{GSO: false, KSO: true, BSO: true, VSO: false}, wantErr: false},
		{name: "kso and vso", tested: ReportRecipients{GSO: false, KSO: true, BSO: false, VSO: true}, wantErr: false},
		{name: "bso and vso", tested: ReportRecipients{GSO: false, KSO: false, BSO: true, VSO: true}, wantErr: false},
		{name: "gso and kso and bso", tested: ReportRecipients{GSO: true, KSO: true, BSO: true, VSO: false}, wantErr: false},
		{name: "gso and kso and vso", tested: ReportRecipients{GSO: true, KSO: true, BSO: false, VSO: true}, wantErr: false},
		{name: "gso and bso and vso", tested: ReportRecipients{GSO: true, KSO: false, BSO: true, VSO: true}, wantErr: false},
		{name: "kso and bso and vso", tested: ReportRecipients{GSO: false, KSO: true, BSO: true, VSO: true}, wantErr: false},
		{name: "all", tested: ReportRecipients{GSO: true, KSO: true, BSO: true, VSO: true}, wantErr: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.tested.Validate(); (err != nil) != test.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
