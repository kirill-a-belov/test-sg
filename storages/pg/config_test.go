package pg

import "testing"

type TestData struct {
	IsError bool
	Config  *PostgresConfig
}

func TestValidate(t *testing.T) {
	testData := []*TestData{
		{
			IsError: false,
			Config: &PostgresConfig{
				Host:     "localhost",
				Port:     12345,
				User:     "usr",
				Password: "passwd",
				DBName:   "some_db",
			},
		},
		{
			IsError: true,
			Config: &PostgresConfig{
				Host:     "localhost",
				Port:     12345,
				User:     "usr",
				Password: "",
				DBName:   "some_db",
			},
		},
	}

	for _, td := range testData {
		err := td.Config.Validate()
		if (err != nil && !td.IsError) ||
			(err == nil && td.IsError) {
			t.Error(
				"For config:", td.Config,
				"validation error expectation", td.IsError,
				"got", err != nil,
			)
		}
	}
}
