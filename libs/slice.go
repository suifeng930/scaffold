package libs

import "reflect"

func InSlice(val interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == false {
			continue
		}
		index = i
		exists = true
		return
	}
	return
}

// SliceDuplicationByStr 切片去重
func SliceDuplicationByStr(s []string) []string {
	var (
		newSlice = make([]string, 0)
		sMap     = map[string]struct{}{}
	)
	for _, v := range s {
		if _, ok := sMap[v]; !ok {
			sMap[v] = struct{}{}
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func SliceIntersect(uIds []string, s []string) []string {
	var (
		newSlice = make([]string, 0)
		sMap     = map[string]struct{}{}
	)
	for _, uid := range uIds {
		sMap[uid] = struct{}{}

	}
	for _, v := range s {
		if _, ok := sMap[v]; ok {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func SliceUnion(agentIds []string, invitedIds []string) []string {
	var (
		newSlice = make([]string, 0)
		aMap     = map[string]struct{}{}
	)
	for _, uid := range agentIds {
		aMap[uid] = struct{}{}
		newSlice = append(newSlice, uid)

	}
	for _, v := range invitedIds {
		if _, ok := aMap[v]; !ok {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
