package internal

import "testing"

func TestCreateTask(t *testing.T) {
	err := CreateTask("TESTTEST")
	if err != nil {
		t.Errorf("resultado '%s', esperado nil", err.Error())
	}
}
