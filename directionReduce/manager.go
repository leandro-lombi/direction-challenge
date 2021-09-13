package directionreduce

func DirReduc(direction []string) []string {

	var from, to string
	result := []string{}
	clean := true

	for i := 0; i < len(direction); i++ {
		to = direction[i]

		if (from == "NORTH" && to == "SOUTH") ||
			(from == "SOUTH" && to == "NORTH") ||
			(from == "EAST" && to == "WEST") ||
			(from == "WEST" && to == "EAST") {
			direction[i] = ""
			clean = false
		} else {
			from = to
			if clean {
				result = append(result, from)
			} else {
				result = append([]string{}, from)
			}
		}
	}

	if len(result) == 1 && result[0] == direction[0] {
		result = append([]string{}, "")
	}

	return result
}