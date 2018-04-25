package main

type intSet struct {
	set map[int]bool
}

func (set *intSet) add(val int) bool {
	if set.set == nil {
		set.set = make(map[int]bool)
	}
	_, found := set.set[val]
	set.set[val] = true
	return found
}

func (set intSet) has(val int) bool {
	if set.set == nil {
		return false
	}

	_, found := set.set[val]
	return found
}

func (set intSet) size() int {
	if set.set == nil {
		return 0
	}

	return len(set.set)
}
