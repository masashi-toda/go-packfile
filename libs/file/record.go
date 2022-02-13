package file

import "strings"

type Record struct {
	headers []string
	values  []string
}

func (r *Record) Headers() []string {
	return r.headers
}

func (r *Record) Values() []string {
	return r.values
}

func (r *Record) TargetValues(targetHeaders ...string) []string {
	values := make([]string, 0)
	for pos, header := range r.headers {
		for _, target := range targetHeaders {
			if strings.EqualFold(header, target) {
				values = append(values, r.values[pos])
			}
		}
	}
	return values
}
