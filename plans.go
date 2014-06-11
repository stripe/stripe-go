package stripe

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type PlanInternval string

const (
	Day   PlanInternval = "day"
	Week  PlanInternval = "week"
	Month PlanInternval = "month"
	Year  PlanInternval = "year"
)

type PlanParams struct {
	Id, Name                   string
	Currency                   Currency
	Amount                     uint64
	Interval                   PlanInternval
	IntervalCount, TrialPeriod uint64
	Meta                       map[string]string
	Statement                  string
}

type Plan struct {
	Id            string            `json:"id"`
	Live          bool              `json:"livemode"`
	Amount        uint64            `json:"amount"`
	Created       int64             `json:"created"`
	Currency      Currency          `json:"currency"`
	Interval      PlanInternval     `json:"interval"`
	IntervalCount uint64            `json:"interval_count"`
	Name          string            `json:"name"`
	Meta          map[string]string `json:"metadata,omitempty"`
	TrialPeriod   uint64            `json:"trial_period_days"`
	Statement     string            `json:"statement_description,omitempty"`
}

type PlanClient struct {
	api   Api
	token string
}

func (c *PlanClient) Create(params *PlanParams) (*Plan, error) {
	body := &url.Values{
		"id":       {params.Id},
		"name":     {params.Name},
		"amount":   {strconv.FormatUint(params.Amount, 10)},
		"currency": {string(params.Currency)},
		"interval": {string(params.Interval)},
	}

	if params.IntervalCount > 0 {
		body.Add("interval_count", strconv.FormatUint(params.IntervalCount, 10))
	}

	if params.TrialPeriod > 0 {
		body.Add("trial_period_days", strconv.FormatUint(params.TrialPeriod, 10))
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", "/plans", c.token, body)

	if err != nil {
		return nil, err
	}

	var plan Plan
	err = json.Unmarshal(res, &plan)
	return &plan, err
}

func (c *PlanClient) Get(id string) (*Plan, error) {
	res, err := c.api.Call("GET", "/plans/"+id, c.token, nil)

	if err != nil {
		return nil, err
	}

	var plan Plan
	err = json.Unmarshal(res, &plan)
	return &plan, err
}

func (c *PlanClient) Update(id string, params *PlanParams) (*Plan, error) {
	body := &url.Values{}

	if len(params.Name) > 0 {
		body.Add("name", params.Name)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	res, err := c.api.Call("POST", "/plans/"+id, c.token, body)

	if err != nil {
		return nil, err
	}

	var plan Plan
	err = json.Unmarshal(res, &plan)
	return &plan, err
}

func (c *PlanClient) Delete(id string) error {
	_, err := c.api.Call("DELETE", "/plans/"+id, c.token, nil)
	return err
}
