# simple-robo
program is related to remotely controlled robots. The robots are very simplistic, but there are two of them, a ground vehicle, and an airborne vehicle.

The ground vehicle can do the following:
- Turn right 90 degrees. - Turn left 90 degrees.
- Move forward 1 meter.

The airborne can do the following:
- Turn right 90 degrees.
- Turn left 90 degrees.
- Move forward 1 meter.
- Move higher by 10 meters. - Move lower by 10 meters.

This means the robot can move around on a grid, where each grid is 1m x 1m. We can assigns codes to the above commands:
- R: turn right.
- L: turn left.
- F: move forward. - U: move higher. - D: move lower.

This means we can send the robot the following string of commands: 
```
ground RFFFFLLFFFFRFFFF
```
or
```
air URFFFLLLUUUDDD
```
This moves the robots around the grid.

Finally, we can request that the robot tell us the direction it is facing, and which part of the grid it is on. It will tell us the x,y,z position of the coordinate, and whether it is facing N,E,S,W (north, east, south, or west)

For example, if we position the robot at the starting point on the bottom left hand corner of the grid (0,0) facing North, and issue the following command FFR, then the robot will report that its position is: 0,2,0 and its direction is E.

when you run simplerobo.go you will be prompted as follow
```
Enter current position of robo comma separated like 2,3,1 
1,1,1
Enter commands for robo eg 'ground RFFFFLLFFFFRFFFF'
air UUFF
Enter robo facing direction eg. E or W or N or S
N
finally robo is facing to  N
final location of robo  1 3 21
```
