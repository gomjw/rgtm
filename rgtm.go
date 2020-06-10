package rgtm

import "regexp"

// StringToMap returns a map with the regex group names and their values.
func StringToMap(r *regexp.Regexp, s string) (map[string]string, error) {
	names := r.SubexpNames()
	result := r.FindAllStringSubmatch(s, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return m, nil
}
