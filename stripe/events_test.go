package stripe

import "testing"

func TestEvent(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	targetList, err := c.Events.List(&EventListParams{Type: "charge.*"})

	if err != nil {
		t.Error(err)
	}

	if len(targetList.Values) == 0 {
		t.Fatalf("No events returned\n")
	}

	e := targetList.Values[0]

	if len(e.Id) == 0 {
		t.Errorf("Id is not set\n")
	}

	if e.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if len(e.Type) == 0 {
		t.Errorf("Type is not set\n")
	}

	if len(e.Req) == 0 {
		t.Errorf("Request is not set\n")
	}

	if e.Data == nil {
		t.Errorf("Data is not set\n")
	}

	if len(e.Data.Obj) == 0 {
		t.Errorf("Object is empty\n")
	}

	target, err := c.Events.Get(e.Id)

	if err != nil {
		t.Error(err)
	}

	if e.Id != target.Id {
		t.Errorf("Id %q does not match expected id %q\n", e.Id, target.Id)
	}

	targetVal := e.GetObjValue("card", "last4")
	val := target.Data.Obj["card"].(map[string]interface{})["last4"]

	if targetVal != val {
		t.Errorf("Value %q does not match expected value %q\n", targetVal, val)
	}
}
