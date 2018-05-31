package server

import "testing"

type TestData struct {
	IsError bool
	Config  *ServerConfig
}

func TestValidate(t *testing.T) {
	testData := []*TestData{
		{
			IsError: false,
			Config: &ServerConfig{
				Address: "127.0.0.1",
				Port:    12345,
			},
		},
		{
			IsError: true,
			Config: &ServerConfig{
				Address: "127.0.a.1",
				Port:    12345,
			},
		},
		{
			IsError: true,
			Config: &ServerConfig{
				Address: "127.0.0.1",
				Port:    12345000,
			},
		},
	}

	for _, td := range testData {
		err := td.Config.Validate()
		if (err != nil && !td.IsError) ||
			(err == nil && td.IsError) {
			t.Error(
				"For adress:", td.Config.Address, "port:", td.Config.Port,
				"validation error expectation", td.IsError,
				"got", err != nil,
			)
		}
	}
}
