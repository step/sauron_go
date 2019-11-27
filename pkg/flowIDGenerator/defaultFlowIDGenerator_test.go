package flowidgenerator_test

import (
	"testing"
	
	f "github.com/step/sauron_go/pkg/flowidgenerator"
)

func TestDefaultIDGenerator(t *testing.T)  {
	generator := f.NewDefaultFlowIDGenerator()

	expected := "ABCD"
	actual := generator.New()

	if string(expected) != string(actual[:]) {
		t.Errorf("\nexpected => %s\nactual => %s\n", expected, actual)
	}
}