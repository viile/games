package tetris

func display(x, y int, container map[int][]int) {
	var str string
	str += "\033c"
	//str +="|--------------------------------|\n"
	//str +="|- Next  %s|Current %d -|\n"
	//str +="|- Block %s|History %d -|\n"
	//str +="|-----------------------g-|\n"
	for h := y - 1; h >= 0; h-- {
		str += "|"
		for w := 0; w < x; w++ {
			if container[w][h] == 1 {
				str += "▓▓"
			} else {
				str += "  "
			}
		}
		str += "|\n"
	}
	//str +="|------------------------|\n"
	print(str)
}
