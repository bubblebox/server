package model_test

import (
	"testing"

	"github.com/ariejan/firedragon/server/model"
)

func TestURLGetID(t *testing.T) {
	expected := "xyz"

	item := &model.Item{
		ID:      expected,
		Content: "http://example.com",
	}

	actual := item.GetID()

	if actual != expected {
		t.Errorf("GetID() returned '%s' expected '%s' instead", actual, expected)
	}
}

func TestURLSetID(t *testing.T) {
	expected := "xyz"

	item := &model.Item{
		ID:      "abc",
		Content: "http://example.com",
	}

	item.SetID(expected)
	actual := item.GetID()

	if actual != expected {
		t.Errorf("SetID() should change ID to '%s' but got '%s' instead", expected, actual)
	}
}
