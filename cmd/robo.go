package cmd

import "fmt"
//directions map to decide current direction
var directions = map[string]int{
	"N": 0,
	"E": 1,
	"S": 2,
	"W": 3,
}
//directions array to decide next dirction on left or right
var dir = []string{"N", "E", "S", "W"}

func MoveVehicle(x, y, z int, facing, commands string) (xnew,ynew,znew int){
	curFacingDir := facing
	xnew,ynew,znew = x,y,z
	for _, curInstruction := range commands {
		switch curInstruction {
		case 'R':
			curd := directions[curFacingDir]
			curFacingDir = dir[(curd+1)%4]
		case 'L':
			curd := directions[curFacingDir]
			nextDirection:=(curd-1)%4
			//in golang -1%4 gives -1 where as -1%4 should be 3
			if nextDirection<0{
				nextDirection= 4+nextDirection
			}
			curFacingDir = dir[nextDirection]
		case 'F':
			if curFacingDir == "E" {
				xnew++
			}
			if curFacingDir == "W" {
				xnew--
			}
			if curFacingDir == "S" {
				ynew--
			}
			if curFacingDir == "N" {
				ynew++
			}

		case 'U':
			znew = znew + 10
		case 'D':
			znew = znew - 10
		}
	}
	return
}
