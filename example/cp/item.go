package cp

type Item struct {
	Ptr     *int
	Text    string
	Number  int
	Numbers []int
}

func (i *Item) Copy() *Item {
	snap := make([]int, len(i.Numbers))
	copy(snap, i.Numbers)
	val := *i.Ptr
	return &Item{
		Ptr:     &val,
		Text:    i.Text,
		Number:  i.Number,
		Numbers: snap,
	}
}
