package gin_head_check

import (
	"github.com/gin-gonic/gin"
	"sort"
)

const defaultHeadSize = 8

var (
	removeHeadList = []string{
		"Authorization",
	}
)

func AppendRemoveCheckHeader(head ...string) {
	removeHeadList = append(removeHeadList, head...)
	removeHeadList = removeStringDuplicateNotCopy(removeHeadList)
}

func AppendRemoveCheckHeaderList(headList []string) {
	removeHeadList = append(removeHeadList, headList...)
	removeHeadList = removeStringDuplicateNotCopy(removeHeadList)
}

// removeStringDuplicateNotCopy
// Deduplication and sorting of slice elements Destroys slices of the original list to be deduplication
// use in the case of hundreds of content
func removeStringDuplicateNotCopy(list []string) []string {
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

// HeaderCheckAsMap
// remove HEADER by list
// can add by AppendRemoveCheckHeader or AppendRemoveCheckHeaderList
func HeaderCheckAsMap(c *gin.Context) map[string]string {
	requestHeader := make(map[string]string, defaultHeadSize)
	r := c.Request
	for k, v := range r.Header {
		if !stringInSlice(k, removeHeadList) {
			requestHeader[k] = v[0]
		}
	}
	return requestHeader
}

func stringInSlice(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
