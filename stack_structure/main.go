package main

type StackIface interface {
	Empty() bool
	Top() byte
	Pop() bool
	Push(b byte)
}

type Stack []byte

func (s *Stack) Empty() bool {
	return s == nil || len(*s) == 0
}

func (s *Stack) Top() byte {
	if s == nil {
		return 0
	}

	ss := []byte(*s)
	return ss[len(ss)-1]
}

func (s *Stack) Pop() bool {
	if s.Empty() {
		return false
	}

	ss := []byte(*s)
	ss = ss[0 : len(ss)-1]
	*s = Stack(ss)
	return true
}

func (s *Stack) Push(b byte) {
	ss := []byte(*s)
	ss = append(ss, b)
	*s = Stack(ss)
}

// 括号匹配
func isValid(s string) bool {
	brackets := []string{"()", "{}", "[]"}
	counter, opening := make(map[byte]byte), make(map[byte]bool)
	for _, bracket := range brackets {
		counter[bracket[1]] = bracket[0]
		opening[bracket[0]] = true
	}

	var stack Stack
	for _, r := range s {
		c := byte(r)
		if opening[c] {
			stack.Push(c)
		} else {
			if !stack.Empty() && stack.Top() == counter[c] {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	return stack.Empty()
}
