package stringSet

type StringSet struct {
	value []string
}

func Of(slice []string) *StringSet {
	mapSlice := make(map[string]bool)
	var uniqueSlice []string
	for _, element := range slice {
		if !mapSlice[element] {
			mapSlice[element] = true
			uniqueSlice = append(uniqueSlice, element)
		}
	}
	return &StringSet{value: uniqueSlice}
}

func (t StringSet) ToSlice() []string {
	return t.value
}
