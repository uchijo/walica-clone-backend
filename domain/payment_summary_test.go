package domain

import "testing"

var userA = User{
	Name: "A",
	Id:   "A",
}

var userB = User{
	Name: "B",
	Id:   "B",
}

func TestTmpSummaryResolve(t *testing.T) {
	tests := []struct {
		name         string
		summaryA     tmpSummary
		summaryB     tmpSummary
		postSummaryA tmpSummary
		postSummaryB tmpSummary
	}{
		{
			"aが1000円、bが-300円の場合",
			tmpSummary{
				user:  &userA,
				total: 1000,
			},
			tmpSummary{
				user:  &userB,
				total: -300,
			},
			tmpSummary{
				user:  &userA,
				total: 700,
			},
			tmpSummary{
				user:  &userB,
				total: 0,
			},
		},
		{
			"aが300円、bが-1000円の場合",
			tmpSummary{
				user:  &userA,
				total: 300,
			},
			tmpSummary{
				user:  &userB,
				total: -1000,
			},
			tmpSummary{
				user:  &userA,
				total: 0,
			},
			tmpSummary{
				user:  &userB,
				total: -700,
			},
		},
		{
			"aが-1000円、bが300円の場合",
			tmpSummary{
				user:  &userA,
				total: -1000,
			},
			tmpSummary{
				user:  &userB,
				total: 300,
			},
			tmpSummary{
				user:  &userA,
				total: -700,
			},
			tmpSummary{
				user:  &userB,
				total: 0,
			},
		},
		{
			"aが-300円、bが1000円の場合",
			tmpSummary{
				user:  &userA,
				total: -300,
			},
			tmpSummary{
				user:  &userB,
				total: 1000,
			},
			tmpSummary{
				user:  &userA,
				total: 0,
			},
			tmpSummary{
				user:  &userB,
				total: 700,
			},
		},
		{
			"aが1000円、bが-1000円の場合",
			tmpSummary{
				user:  &userA,
				total: 1000,
			},
			tmpSummary{
				user:  &userB,
				total: -1000,
			},
			tmpSummary{
				user:  &userA,
				total: 0,
			},
			tmpSummary{
				user:  &userB,
				total: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.summaryA.resolve(&tt.summaryB)
			successA := tt.postSummaryA.Alike(tt.summaryA)
			successB := tt.postSummaryB.Alike(tt.summaryB)
			if !successA {
				t.Errorf("summaryA's total was expected to be [%v], but got [%v]", tt.postSummaryA.total, tt.summaryA.total)
			}
			if !successB {
				t.Errorf("summaryB's total was expected to be [%v], but got [%v]", tt.postSummaryB.total, tt.summaryB.total)
			}
		})
	}
}

func (ts tmpSummary) Alike(subject tmpSummary) bool {
	return ts.total == subject.total && ts.user.Alike(*subject.user)
}
