package uniqueemailaddresses

import "testing"

func Test_splitEmail(t *testing.T) {
	tests := []struct {
		args           string
		wantLocalName  string
		wantDomainName string
	}{
		{
			args:           "test.email+alex@leetcode.com",
			wantLocalName:  "testemail",
			wantDomainName: "leetcode.com",
		},
		{
			args:           "test.email.leet+alex@code.com",
			wantLocalName:  "testemailleet",
			wantDomainName: "code.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			gotLocalName, gotDomainName := splitEmail(tt.args)
			if gotLocalName != tt.wantLocalName {
				t.Errorf("splitEmail() gotLocalName = %v, want %v", gotLocalName, tt.wantLocalName)
			}
			if gotDomainName != tt.wantDomainName {
				t.Errorf("splitEmail() gotDomainName = %v, want %v", gotDomainName, tt.wantDomainName)
			}
		})
	}
}
