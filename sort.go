package choose

type uints []uint

func (self uints) Len() int {
	return len(self)
}

func (self uints) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self uints) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
