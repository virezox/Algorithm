package main

func main() {

}

func isValid(s string) bool {
	m := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		")": 4,
		"]": 5,
		"}": 6,
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		if m[s[i:i+1]] < 4 {
			stack = append(stack, s[i:i+1])
		}
		if m[s[i:i+1]] > 3 {
			if len(stack) == 0 || m[stack[len(stack)-1]] != m[s[i:i+1]]-3 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
