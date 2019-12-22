package main

import "testing"

func TestFuelRequired(t *testing.T) {
	mass := 12
	if fuel, expected := fuelRequired(mass), 2; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}

	mass = 14
	if fuel, expected := fuelRequired(mass), 2; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}

	mass = 1969
	if fuel, expected := fuelRequired(mass), 654; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}

	mass = 100756
	if fuel, expected := fuelRequired(mass), 33583; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}
}

func TestFuelRequiredWithFuelMass(t *testing.T) {

	mass := 14
	if fuel, expected := fuelRequiredWithFuelMass(mass), 2; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}

	mass = 1969
	if fuel, expected := fuelRequiredWithFuelMass(mass), 966; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}

	mass = 100756
	if fuel, expected := fuelRequiredWithFuelMass(mass), 50346; fuel != expected {
		t.Errorf("Expected mass=%d to be fuel=%d, but got %d", mass, expected, fuel)
	}
}
