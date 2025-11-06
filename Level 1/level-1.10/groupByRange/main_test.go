package main

import "testing"

/*
Уникальные 4 группы температур:
Группа -20: [-25.4, -27.0, -21.0]
Группа 10: [13.0, 19.0, 15.5]
Группа 20: [24.5]
Группа 30: [32.5]
*/

func TestGroupByStep(t *testing.T) {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10

	result := GroupTemperaturesByStep(temperatures, step)

	expectedGroups := 4
	if len(result) != expectedGroups {
		t.Errorf("Ожидали %d групп, получили %d", expectedGroups, len(result))
	}
}
