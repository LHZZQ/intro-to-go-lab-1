package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	h, w := p.imageHeight, p.imageWidth
	aliveVal := byte(255) // 活细胞是 255；死细胞是 0

	// 准备下一代
	next := make([][]byte, h)
	for y := range next {
		next[y] = make([]byte, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sum := 0
			if world[(y+h-1)%h][(x+w-1)%w] != 0 {
				sum++
			}
			if world[(y+h-1)%h][(x+w)%w] != 0 {
				sum++
			}
			if world[(y+h-1)%h][(x+w+1)%w] != 0 {
				sum++
			}
			if world[(y+0)%h][(x+w-1)%w] != 0 {
				sum++
			}
			if world[(y+0)%h][(x+w+1)%w] != 0 {
				sum++
			}
			if world[(y+1)%h][(x+w-1)%w] != 0 {
				sum++
			}
			if world[(y+1)%h][(x+w)%w] != 0 {
				sum++
			}
			if world[(y+1)%h][(x+w+1)%w] != 0 {
				sum++
			}

			// 规则（非 0 视为“当前是活”）
			if world[y][x] != 0 {
				if sum == 2 || sum == 3 {
					next[y][x] = aliveVal
				} else {
					next[y][x] = 0
				}
			} else {
				if sum == 3 {
					next[y][x] = aliveVal
				} else {
					next[y][x] = 0
				}
			}
		}
	}

	for y := 0; y < h; y++ {
		copy(world[y], next[y])
	}
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	h, w := p.imageHeight, p.imageWidth
	var alive []cell
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if world[y][x] != 0 { // 255 视为活
				alive = append(alive, cell{x: x, y: y})
			}
		}
	}
	return alive
}
