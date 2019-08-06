package app

import (
	"testing"
)

func TestProduceTitleChile(t *testing.T) {
	expected := "Trabajadores al poder!"
	res, err := ProduceTitle("chile")
	if err != nil {
		t.Error(err)
	}
	if res != expected {
		t.Errorf("Expected %s, go %s", expected, res)
	}
}

func TestProduceTitleBrazil(t *testing.T) {
	expected := "Oi sim sim sim, oi não não não"
	res, err := ProduceTitle("brazil")
	if err != nil {
		t.Error(err)
	}
	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
