func isValid(s string) bool {
    stack := make([]rune, 0)
    for _, v := range s {
        if rune(v) == '(' || rune(v) == '{' || rune(v) == '[' {
            stack = append(stack, rune(v))
        } else {
            if len(stack) == 0 {
                return false
            }
            check := stack[len(stack)-1]
            if rune(v) == ')' && check == '(' {
                stack = stack[:len(stack)-1]
                continue
            } else if rune(v) == ')' && check != '('{
                return false
            }
            if rune(v) == ']' && check == '[' {
                stack = stack[:len(stack)-1]
                continue
            } else if rune(v) == ']' && check != '[' {
                return false
            }
            if rune(v) == '}' && check == '{' {
                stack = stack[:len(stack)-1]
                continue
            } else if rune(v) == '}' && check != '{' {
                return false
            }
        }
    }

    if len(stack) != 0 {
        return false
    }
    return true
}
