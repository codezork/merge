package interval

func (a Array) Len() int {
	return len(a.Inputs)
}

func (a Array) Less(i, j int) bool {
	return a.Inputs[i].Low < a.Inputs[j].Low
}

func (a Array) Swap(i, j int) {
	a.Inputs[i], a.Inputs[j] = a.Inputs[j], a.Inputs[i]
}
