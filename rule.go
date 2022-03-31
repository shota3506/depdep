package depdep

import (
	"regexp"
)

type rule struct {
	mapping map[*regexp.Regexp][]*regexp.Regexp
}

func (r *rule) add(from string, tos ...string) error {
	rfrom, err := regexp.Compile(from)
	if err != nil {
		return err
	}
	rtos := make([]*regexp.Regexp, len(tos))
	for i, to := range tos {
		var err error
		rtos[i], err = regexp.Compile(to)
		if err != nil {
			return err
		}
	}

	r.mapping[rfrom] = rtos
	return nil
}

func (r *rule) check(from, to string) bool {
	for r, rtos := range r.mapping {
		if !r.MatchString(from) {
			continue
		}
		for _, rto := range rtos {
			if rto.MatchString(to) {
				return true
			}
		}
	}
	return false
}
