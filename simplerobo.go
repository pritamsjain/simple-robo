package main

import (
	"fmt"
	"strings"
	"strconv"
	 "simple-robo/cmd"
)

func main() {
	var input,strcmd, xyz,facingDirection string
	var x, y, z int
	fmt.Println("Enter current position of robo comma separated like 2,3,1 ")
	fmt.Scanln(&xyz)
	fmt.Println(xyz)
	coordinates := strings.Split(xyz, ",")
	var coordinatesint []int
	for i, a := range coordinates {
		in, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
			return
		}
		coordinatesint = append(coordinatesint, in)
		switch i {
		case 0:
			x = in
		case 1:
			y = in
		case 2:
			z = in

		}
	}

	fmt.Println("Enter commands for robo eg 'ground RFFFFLLFFFFRFFFF'")
	fmt.Scanln(&input,&strcmd)
	fmt.Println(input,strcmd)

	fmt.Println("Enter robo facing direction eg. E or W or N or S")
	fmt.Scanln(&facingDirection)
	fmt.Println(facingDirection)

	newx,newy,newz := cmd.MoveVehicle(x, y, z, facingDirection, strcmd)
	fmt.Println(newx,newy,newz)
}