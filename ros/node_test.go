package ros

import (
	"testing"
)

func TestLoadJsonFromString(t *testing.T) {
	value, err := loadParamFromString("42")
	if err != nil {
		t.Error(err)
	}
	i, ok := value.(float64)
	if !ok {
		t.Fail()
	}
	if i != 42.0 {
		t.Error(i)
	}
}

func TestNewNodeWithLogger(t *testing.T) {
	logger := NewDefaultLogger()
	node, err := NewNodeWithLogger("mynode", []string{}, logger)

	if err != nil {
		t.Fatal("New node should be created")
	}

	if node.Logger() != logger {
		t.Fatal("Logger should be set from the given argument")
	}
}
