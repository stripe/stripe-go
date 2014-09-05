package plan

import (
	"fmt"
	"testing"

	. "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestPlanCreate(t *testing.T) {
	planParams := &PlanParams{
		Id:            "test_plan",
		Name:          "Test Plan",
		Amount:        99,
		Currency:      USD,
		Interval:      Month,
		IntervalCount: 3,
		TrialPeriod:   30,
		Statement:     "Test Plan",
	}

	target, err := Create(planParams)

	if err != nil {
		t.Error(err)
	}

	if target.Id != planParams.Id {
		t.Errorf("Id %q does not match expected id %q\n", target.Id, planParams.Id)
	}

	if target.Name != planParams.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, planParams.Name)
	}

	if target.Amount != planParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, planParams.Amount)
	}

	if target.Currency != planParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, planParams.Currency)
	}

	if target.Interval != planParams.Interval {
		t.Errorf("Interval %q does not match expected interval %q\n", target.Interval, planParams.Interval)
	}

	if target.IntervalCount != planParams.IntervalCount {
		t.Errorf("Interval count %v does not match expected interval count %v\n", target.IntervalCount, planParams.IntervalCount)
	}

	if target.TrialPeriod != planParams.TrialPeriod {
		t.Errorf("Trial period %v does not match expected trial period %v\n", target.TrialPeriod, planParams.TrialPeriod)
	}

	if target.Statement != planParams.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, planParams.Statement)
	}

	Delete(planParams.Id)
}

func TestPlanGet(t *testing.T) {
	planParams := &PlanParams{
		Id:       "test_plan",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	Create(planParams)
	target, err := Get(planParams.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if target.Id != planParams.Id {
		t.Errorf("Plan id %q does not match expected id %q\n", target.Id, planParams.Id)
	}

	Delete(planParams.Id)
}

func TestPlanUpdate(t *testing.T) {
	planParams := &PlanParams{
		Id:            "test_plan",
		Name:          "Original Name",
		Amount:        99,
		Currency:      USD,
		Interval:      Month,
		IntervalCount: 3,
		TrialPeriod:   30,
		Statement:     "Original Plan",
	}

	Create(planParams)

	updatedPlan := &PlanParams{
		Name:      "Updated Name",
		Statement: "Updated Plan",
	}

	target, err := Update(planParams.Id, updatedPlan)

	if err != nil {
		t.Error(err)
	}

	if target.Name != updatedPlan.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, updatedPlan.Name)
	}

	if target.Statement != updatedPlan.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, updatedPlan.Statement)
	}

	Delete(planParams.Id)
}

func TestPlanList(t *testing.T) {
	for i := 0; i < 5; i++ {
		planParams := &PlanParams{
			Id:       fmt.Sprintf("test_%v", i),
			Name:     fmt.Sprintf("test_%v", i),
			Amount:   99,
			Currency: USD,
			Interval: Month,
		}

		Create(planParams)
	}

	target, err := List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for i := 0; i < 5; i++ {
		Delete(fmt.Sprintf("test_%v", i))
	}
}
