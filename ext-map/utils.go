package ext_map

// check value contain in map
// @param string val
// @param map[string]string m
// @return string the key of value
// @return bool val is contains
func StringMapContains(val string, m map[string]string) (string, bool) {
	if m == nil {
		return "", false
	}

	for k,v := range m {
		if v == val {
			return k, true
		}
	}

	return "", false
}