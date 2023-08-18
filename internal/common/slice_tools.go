package common

import "sort"

// RemoveStringDuplicateNotCopy
// Deduplication and sorting of slice elements Destroys slices of the original list to be deduplication
// use in the case of hundreds of content
func RemoveStringDuplicateNotCopy(list []string) []string {
	if list == nil {
		return nil
	}
	sort.Strings(list)
	uniq := list[:0]
	for _, x := range list {
		if len(uniq) == 0 || uniq[len(uniq)-1] != x {
			uniq = append(uniq, x)
		}
	}
	return uniq
}

// RemoveStringDuplicateUseMap
// list slices to be duplicated
// use after length exceeds 10000
func RemoveStringDuplicateUseMap(list []string) []string {
	var data []string
	rd := map[string]struct{}{}
	for _, v := range list {
		if _, ok := rd[v]; !ok { // Add elements in the corresponding slice by determining whether there is a corresponding key value in the map.
			rd[v] = struct{}{}
			data = append(data, v)
		}
	}
	return data
}
