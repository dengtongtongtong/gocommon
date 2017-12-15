package datastructutils

import "sort"

// 按key升序排列
func SortMapByStringKey(unsortedmap map[string]interface{}) (sortedmap []interface{}, err error) {
	if len(unsortedmap) < 1 {
		sortedmap = []interface{}{}
		return sortedmap, nil
	}
	var keys []string
	for k, _ := range unsortedmap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sortedmap = append(sortedmap, k)
	}
	return sortedmap, nil
}

// 按key降序排列
func ReverseSortMapByStringKey(unsortedmap map[string]interface{}) (sortedmap []interface{}, err error) {
	if len(unsortedmap) < 1 {
		sortedmap = []interface{}{}
		return sortedmap, nil
	}
	var keys []string
	for k, _ := range unsortedmap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, k := range keys {
		sortedmap = append(sortedmap, k)
	}
	return sortedmap, nil
}
