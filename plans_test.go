package stripe

import (
	"fmt"
	"testing"
)

func TestPlanCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	plan := &PlanParams{
		Id:            "test_plan",
		Name:          "Test Plan",
		Amount:        99,
		Currency:      USD,
		Interval:      Month,
		IntervalCount: 3,
		TrialPeriod:   30,
		Statement:     "Test Plan",
	}

	target, err := c.Plans.Create(plan)

	if err != nil {
		t.Error(err)
	}

	if target.Id != plan.Id {
		t.Errorf("Id %q does not match expected id %q\n", target.Id, plan.Id)
	}

	if target.Name != plan.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, plan.Name)
	}

	if target.Amount != plan.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, plan.Amount)
	}

	if target.Currency != plan.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, plan.Currency)
	}

	if target.Interval != plan.Interval {
		t.Errorf("Interval %q does not match expected interval %q\n", target.Interval, plan.Interval)
	}

	if target.IntervalCount != plan.IntervalCount {
		t.Errorf("Interval count %v does not match expected interval count %v\n", target.IntervalCount, plan.IntervalCount)
	}

	if target.TrialPeriod != plan.TrialPeriod {
		t.Errorf("Trial period %v does not match expected trial period %v\n", target.TrialPeriod, plan.TrialPeriod)
	}

	if target.Statement != plan.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, plan.Statement)
	}

	c.Plans.Delete(plan.Id)
}

func TestPlanGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	plan := &PlanParams{
		Id:       "test_plan",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	c.Plans.Create(plan)
	target, err := c.Plans.Get(plan.Id)

	if err != nil {
		t.Error(err)
	}

	if target.Id != plan.Id {
		t.Errorf("Plan id %q does not match expected id %q\n", target.Id, plan.Id)
	}

	c.Plans.Delete(plan.Id)
}

func TestPlanUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	plan := &PlanParams{
		Id:            "test_plan",
		Name:          "Original Name",
		Amount:        99,
		Currency:      USD,
		Interval:      Month,
		IntervalCount: 3,
		TrialPeriod:   30,
		Statement:     "Original Plan",
	}

	c.Plans.Create(plan)

	updatedPlan := &PlanParams{
		Name:      "Updated Name",
		Statement: "Updated Plan",
	}

	target, err := c.Plans.Update(plan.Id, updatedPlan)

	if err != nil {
		t.Error(err)
	}

	if target.Name != updatedPlan.Name {
		t.Errorf("Name %q does not match expected name %q\n", target.Name, updatedPlan.Name)
	}

	if target.Statement != updatedPlan.Statement {
		t.Errorf("Statement %q does not match expected statement %q\n", target.Statement, updatedPlan.Statement)
	}

	c.Plans.Delete(plan.Id)
}

func TestPlanList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	for i := 0; i < 5; i++ {
		plan := &PlanParams{
			Id:       fmt.Sprintf("test_%v", i),
			Name:     fmt.Sprintf("test_%v", i),
			Amount:   99,
			Currency: USD,
			Interval: Month,
		}

		c.Plans.Create(plan)
	}

	target, err := c.Plans.List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for i := 0; i < 5; i++ {
		c.Plans.Delete(fmt.Sprintf("test_%v", i))
	}
}
