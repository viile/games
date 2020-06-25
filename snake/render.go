package snake

func display(x, y int, container map[int][]int) {
	var str string
	str += "\033c"
	for h := y - 1; h >= 0; h-- {
		str += "|"
		for w := 0; w < x; w++ {
			if container[w][h] == PointSnake {
				str += "🐍"
			} else if container[w][h] == PointFruit {
				str += "🍎"
			} else {
				str += "  "
			}
		}
		str += "|\n"
	}
	print(str)
}
