package sum

func Sum(num [5]int) int {

	result := 0
	for _, i := range num {
		result += i
	}
	return result

}
