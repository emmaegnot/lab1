package main

// takes the current state of the world and completes one evolution
// returns the result
func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := [][]byte{}
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			//get current cell
			cCell := world[y][x]
			//count adjacent alive cells
			//this doesn't work because it goes out of bounds
			// how can top and bottom be connected
			// use mods?
			alive := 0
			if world[y-1][x-1] == 255 {
				alive++
			}
			if world[y][x-1] == 255 {
				alive++
			}
			if world[y+1][x-1] == 255 {
				alive++
			}
			if world[y-1][x] == 255 {
				alive++
			}
			if world[y+1][x] == 255 {
				alive++
			}
			if world[y-1][x+1] == 255 {
				alive++
			}
			if world[y][x+1] == 255 {
				alive++
			}
			if world[y+1][x+1] == 255 {
				alive++
			}
			//check what needs to happen with this many alive cells
			if cCell == 255 {
				if alive < 2 {
					cCell = 0
				} else if alive > 3 {
					cCell = 0
				}
			} else {
				if alive == 3 {
					cCell = 255
				}
			}
			//append to list of cells to change
			newWorld[y][x] = cCell
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
