package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	pk := "Wallet#2022-05-12T14:40:29.156Z"
	sk := "Transaction#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	targetAmount := 1000
	currentAmount := 500
	targetDate := "2022-05-12T20:25:19Z"
	goalName := "goalName"
	goalAchieved := true
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"creation_date\":\"" + creationDate + "\",\"target_amount\":" + strconv.Itoa(targetAmount) + ",\"current_amount\":" + strconv.Itoa(currentAmount) + ",\"target_date\":\"" + targetDate + "\",\"updated_date\":\"" + updatedDate + "\",\"goal_name\":\"" + goalName + "\",\"goal_achieved\": " + strconv.FormatBool(goalAchieved) + "}"
	queryParameter := QueryParameter{}
	err := json.Unmarshal([]byte(body), &queryParameter)

	if *queryParameter.Pk != pk {
		t.Errorf("QueryParameter: PK do not match, got = %v, want = %v", *queryParameter.Pk, pk)
		return
	}

	if queryParameter.Sk != sk {
		t.Errorf("QueryParameter: SK do not match, got = %v, want = %v", queryParameter.Sk, sk)
		return
	}

	if *queryParameter.CreationDate != creationDate {
		t.Errorf("QueryParameter: Creation Date do not match, got = %v, want = %v", queryParameter.CreationDate, creationDate)
		return
	}

	if *queryParameter.UpdatedDate != updatedDate {
		t.Errorf("QueryParameter: Updated Date do not match, got = %v, want = %v", queryParameter.UpdatedDate, updatedDate)
		return
	}

	if *queryParameter.TargetAmount != float64(targetAmount) {
		t.Errorf("QueryParameter: TargetAmount do not match, got = %v, want = %v", *queryParameter.TargetAmount, targetAmount)
		return
	}

	if *queryParameter.CurrentAmount != float64(currentAmount) {
		t.Errorf("QueryParameter: Current Amount do not match, got = %v, want = %v", *queryParameter.CurrentAmount, currentAmount)
		return
	}

	if *queryParameter.TargetDate != targetDate {
		t.Errorf("QueryParameter: Target Date do not match, got = %v, want = %v", queryParameter.TargetDate, targetDate)
		return
	}

	if *queryParameter.Name != goalName {
		t.Errorf("QueryParameter: Target Date do not match, got = %v, want = %v", queryParameter.TargetDate, targetDate)
		return
	}

	if *queryParameter.GoalAchieved != goalAchieved {
		t.Errorf("QueryParameter: Goal Achieved do not match, got = %v, want = %v", queryParameter.GoalAchieved, goalAchieved)
		return
	}

	if err != nil {
		t.Errorf("QueryParameter Struct has an error = %v", err)
		return
	}
}
