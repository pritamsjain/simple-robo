package cmd

import (
	"testing"
)

func TestMoveVehicle1(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "E", "FFF")
		if x != 4 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 3, 1, 1)
		}
	}
}
func TestMoveVehicle2(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "N", "FFF")
		if y != 4 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, 4, 1)
		}
	}
}
func TestMoveVehicle3(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "FFF")
		if x != -2 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 2, 1, 1)
		}

	}
}
func TestMoveVehicle4(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "UUU")
		if z != 31 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, 1, 31)
		}
	}
}
func TestMoveVehicle5(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "UD")
		if z != 1 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, 1, 1)
		}
	}
}
func TestMoveVehicle6(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "RFF")
		if y != 3 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, 3, 1)
		}
	}
}
func TestMoveVehicle7(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "LFF")
		if y != -1 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, -1, 1)
		}
	}
}
func TestMoveVehicle8(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "LLLLLFF")
		if y != -1 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, 1, -1, 1)
		}
	}
}
func TestMoveVehicle9(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "LLLLFF")
		if x != -1 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, -1, 1, 1)
		}
	}
}
func TestMoveVehicle10(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "RRRRFF")
		if x != -1 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, -1, 1, 1)
		}
	}
}
func TestMoveVehicle11(t *testing.T) {
	{
		x, y, z := MoveVehicle(1, 1, 1, "W", "FFRRRRFF")
		if x != -3 {
			t.Errorf("robo not moving well %d %d %d expected was %d %d %d ", x, y, z, -3, 1, 1)
		}
	}
}