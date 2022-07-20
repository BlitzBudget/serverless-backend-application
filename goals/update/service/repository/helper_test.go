package repository

import (
	"strconv"
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "pk"
	sk := "sk"
	targetAmount := 1000
	goalAchieved := true
	goalName := "goalName"
	currentAmount := 2000
	targetDate := "2022-09-15"
	body := "{\"pk\":\"" + pk + "\",\"sk\":\"" + sk + "\",\"target_amount\":" + strconv.Itoa(targetAmount) + ",\"goal_achieved\":" + strconv.FormatBool(goalAchieved) + ",\"goal_name\":\"" + goalName + "\",\"current_amount\":" + strconv.Itoa(currentAmount) + ",\"target_date\":\"" + targetDate + "\"}"

	requestModel := AttributeBuilder(&body)

	if requestModel == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *requestModel.Pk != pk {
		t.Errorf("pk convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Pk, pk)
		return
	}

	if *requestModel.Sk != sk {
		t.Errorf("SK convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Sk, sk)
		return
	}

	if *requestModel.TargetAmount != float64(targetAmount) {
		t.Errorf("TargetAmount convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.TargetAmount, targetAmount)
		return
	}

	if *requestModel.CurrentAmount != float64(currentAmount) {
		t.Errorf("CurrentAmount convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.CurrentAmount, currentAmount)
		return
	}

	if *requestModel.GoalAchieved != goalAchieved {
		t.Errorf("GoalAchieved convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.GoalAchieved, goalAchieved)
		return
	}

	if *requestModel.Name != goalName {
		t.Errorf("Name convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.Name, goalName)
		return
	}

	if *requestModel.TargetDate != targetDate {
		t.Errorf("TargetDate convertion to DynamoDB attribute not correct, got = %v, want = %v", *requestModel.TargetDate, targetDate)
		return
	}
}
