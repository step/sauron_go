package parser_test

import (
	"testing"
	"reflect"

	"github.com/step/sauron_go/pkg/parser"
	"github.com/step/saurontypes"
)

func TestGetTasks(t *testing.T) {
	actual := parser.GetTasks("gol")
	
	expected := []saurontypes.Task{
		{Queue: "test", ImageName: "mocha"},
		{Queue: "lint", ImageName: "eslint"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nactual => %s\nexpected => %s", actual, expected)
	}
}