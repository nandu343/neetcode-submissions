type MinStack struct {
	stack []int
	mono []int
}

func Constructor() MinStack {
	return MinStack{make([]int,0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.mono) == 0 {
		this.mono = append(this.mono, val)
	} else if this.mono[len(this.mono)-1] < val {
		this.mono = append(this.mono, this.mono[len(this.mono)-1])
	} else {
		this.mono = append(this.mono, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mono = this.mono[:len(this.mono)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mono[len(this.mono)-1]
}
