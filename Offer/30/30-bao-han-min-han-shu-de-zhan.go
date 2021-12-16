package offer

type MinStack struct {
	minstack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{minstack: []int{}, stack: []int{}}
}

func (s *MinStack) Push(x int) {
	if min := s.Min(); min <= x && len(s.stack) > 0 {
		s.minstack = append(s.minstack, min)
	} else {
		s.minstack = append(s.minstack, x)
	}
	s.stack = append(s.stack, x)
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.minstack = s.minstack[:len(s.minstack)-1]
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.minstack[len(s.minstack)-1]
}
