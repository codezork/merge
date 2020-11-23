package handler

import (
	"github.com/merge/pkg/helper"
	"github.com/merge/pkg/interval"
	"sort"
)

// 1. Sort the intervals based on increasing order of low value
// 2. Push the first Interval on to a stack
// 3. For each Interval perform below steps:
// 	3.1. If the current Interval does not overlap with the top of the stack, push it.
// 	3.2. If the current Interval overlaps with top of the stack and high value of current Interval is more
// 		than that of top of stack, update stack top with the high value of current Interval.
// 4. Finally, stack contains the merged intervals.

func (h *InputHandler) Merge() []*interval.Interval{
	defer helper.Elapsed("merge", h.Verbose)()
	var n int
	if n = len(h.Array.Inputs); n == 0 {
		return h.Array.Inputs
	}
 	stack := new(helper.Stack)
	sort.Sort(interval.Array(h.Array))
	stack.Push(h.Array.Inputs[0])

	for i := 1; i < n; i++ {
		top := stack.Top()
		if top.(*interval.Interval).High < h.Array.Inputs[i].Low {
			stack.Push(h.Array.Inputs[i])
		} else if top.(*interval.Interval).High < h.Array.Inputs[i].High {
			top.(*interval.Interval).High = h.Array.Inputs[i].High
			stack.Pop()
			stack.Push(top)
		}
	}
	results := make([]*interval.Interval, stack.Len())
	i := stack.Len()-1
	for stack.Len() > 0 {
		results[i] = stack.Top().(*interval.Interval)
		stack.Pop()
		i --
	}
	return results
}