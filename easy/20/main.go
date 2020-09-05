package main

func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	stack := &stack{[]rune{}}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.Push(r)
		case ')':
			if stack.Pop() != '(' {
				return false
			}
		case '}':
			if stack.Pop() != '{' {
				return false
			}
		case ']':
			if stack.Pop() != '[' {
				return false
			}
		}
	}

	return len(stack.data) == 0
}

type stack struct {
	data []rune
}

func (s *stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *stack) Pop() rune {
	if len(s.data) == 0 {
		return 'e'
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}
