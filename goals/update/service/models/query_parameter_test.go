package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	currentAmount := 10
	targetAmount := 12
	goalAchieved := true
	targetDate := "2022-10-10"
	name := "name"
	updatedDate := "2022-10-19"
	body := "{\":c\":" + strconv.Itoa(currentAmount) + ",\":a\":" + strconv.Itoa(targetAmount) + ",\":g\":" + strconv.FormatBool(goalAchieved) + ",\":d\":\"" + targetDate + "\",\":n\":\"" + name + "\",\":u\":\"" + updatedDate + "\"}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.CurrentAmount != float64(currentAmount) {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.CurrentAmount, currentAmount)
		return
	}

	if *queryParameter.TargetAmount != float64(targetAmount) {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.TargetAmount, targetAmount)
		return
	}

	if *queryParameter.GoalAchieved != goalAchieved {
		t.Errorf("QueryParameter: Goal Achieved do not match, got = %v, want = %v", queryParameter.GoalAchieved, goalAchieved)
		return
	}

	if *queryParameter.TargetDate != targetDate {
		t.Errorf("QueryParameter: Target Date do not match, got = %v, want = %v", queryParameter.TargetDate, targetDate)
		return
	}

	if *queryParameter.Name != name {
		t.Errorf("QueryParameter: Name do not match, got = %v, want = %v", *queryParameter.Name, name)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Update Date do not match, got = %v, want = %v", *queryParameter.UpdatedDate, updatedDate)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
