package main

// takes the current state of the world and completes one evolution
// returns the result
func calculateNextState(p golParams, world [][]byte) [][]byte {
	imht := p.imageHeight
	imwd := p.imageWidth
	// new (outer) slice for the new world
	newWorld := make([][]byte, imht)
	for y := 0; y < imht; y++ {
		// creating a slice in each element of the outer slice
		newWorld[y] = make([]byte, imwd)
		for x := 0; x < imwd; x++ {
			//get value of current cell
			cVal := world[y][x]
			//count adjacent alive cells
			//this could also be done by summing and dividing by 255 instead of having loads of if stmts
			alive := 0
			if world[(y+imht-1)%imht][(x+imwd-1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht)%imht][(x+imwd-1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht+1)%imht][(x+imwd-1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht+1)%imht][(x+imwd)%imwd] == 255 {
				alive++
			}
			if world[(y+imht-1)%imht][(x+imwd+1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht)%imht][(x+imwd+1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht+1)%imht][(x+imwd+1)%imwd] == 255 {
				alive++
			}
			if world[(y+imht-1)%imht][(x+imwd)%imwd] == 255 {
				alive++
			}
			//check what needs to happen with this many alive cells
			if cVal == 255 {
				if alive < 2 {
					cVal = 0
				} else if alive > 3 {
					cVal = 0
				}
			} else {
				if alive == 3 {
					cVal = 255
				}
			}
			//append to list of cells to change
			newWorld[y][x] = cVal
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	imht := p.imageHeight
	imwd := p.imageWidth
	var aliveCells []cell
	for y := 0; y < imht; y++ {
		for x := 0; x < imwd; x++ {
			cVal := world[y][x]
			if cVal == 255 {
				cCell := cell{x, y}
				aliveCells = append(aliveCells, cCell)
			}
		}
	}
	return aliveCells
}
