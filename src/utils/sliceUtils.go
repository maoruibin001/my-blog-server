package utils

func RemoveStringElement(arr []string, cb func(ele string, index int) bool) []string {
	ret := []string{}
	for i := 0; i < len(arr); i ++ {
		if cb(arr[i], i) == false {
			ret = append(ret, arr[i])
		}
	}
	return ret
}