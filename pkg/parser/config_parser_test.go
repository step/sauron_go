package parser_test

import (
	"bytes"
	"testing"
	"reflect"

	"github.com/step/saurontypes"

	"github.com/spf13/viper"
	"github.com/step/sauron_go/pkg/parser"
)

func TestGetTasks(t *testing.T) {

	viperInst := viper.New()
	viperInst.SetConfigType("toml")

	var sampleConfig = []byte(
		`[[assignments]]
			name = "gol"
			description = "something"
			prefix = "gol-"

				[[assignments.tasks]]
				name = "test"
				queue = "test"
				image = "steptw/test"
				data = "/github/somewhere"`,
	)

	expectedSauronConfig := saurontypes.SauronConfig{
		Assignments: []saurontypes.Assignment{
			{
				Description: "something",
				Prefix: "gol-",
				Name:"gol",
				Tasks: []saurontypes.Task{
					{
						Queue:"test",
						ImageName:"steptw/test",
						Name:"test",
						Data:"/github/somewhere",
					},
				},
			},
		},
	}

	viperInst.ReadConfig(bytes.NewBuffer(sampleConfig))

	sauronConfig := parser.ParseConfig(*viperInst)

	if !reflect.DeepEqual(sauronConfig, expectedSauronConfig){
		t.Errorf("\n actual => %s \n expected => %s\n", sauronConfig, expectedSauronConfig)
	}
}
