package stripe

import (
	"testing"
)

func TestAccountGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	target, err := c.Account.Get()

	if err != nil {
		t.Error(err)
	}

	if len(target.Id) == 0 {
		t.Errorf("Account is missing id\n")
	}

	if len(target.Country) == 0 {
		t.Errorf("Account is missing country\n")
	}

	if len(target.Currencies) == 0 {
		t.Errorf("Account is missing currencies\n")
	}

	if len(target.DefaultCurrency) == 0 {
		t.Errorf("Account is missing default currency\n")
	}

	if len(target.Name) == 0 {
		t.Errorf("Account is missing name\n")
	}

	if len(target.Email) == 0 {
		t.Errorf("Account is missing email\n")
	}

	if len(target.Timezone) == 0 {
		t.Errorf("Account is missing timezone\n")
	}
}
