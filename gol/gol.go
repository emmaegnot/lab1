package main

// takes the current state of the world and completes one evolution
// returns the result
func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := [][]byte{}
	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[i]); j++ {
			//get current cell
			cCell := world[i][j]
			//count adjacent alive cells
			alive := 0
			if world[i-1][j-1] == 255 {
				alive++
			}
			if world[i][j-1] == 255 {
				alive++
			}
			if world[i+1][j-1] == 255 {
				alive++
			}
			if world[i-1][j] == 255 {
				alive++
			}
			if world[i+1][j] == 255 {
				alive++
			}
			if world[i-1][j+1] == 255 {
				alive++
			}
			if world[i][j+1] == 255 {
				alive++
			}
			if world[i+1][j+1] == 255 {
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
			newWorld[i][j] = cCell
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
