func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
			case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack[len(stack)-2] = a + b
			stack = stack[:len(stack)-1]
			case "-":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack[len(stack)-2] = b - a
			stack = stack[:len(stack)-1]
			case "*":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack[len(stack)-2] = b * a
			stack = stack[:len(stack)-1]
			case "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack[len(stack)-2] = b / a
			stack = stack[:len(stack)-1]
			default:
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}
		
	}
	return stack[0]
}
