package fingerprint

import (
	"strconv"
	"strings"
)

var (
	greaseChars    = []byte{' ', '(', ':', '-', '.', '/', ')', ';', '=', '?', '_'}
	greaseVersions = []string{"8", "99", "24"}
	brandOrders    = map[int][][]int{
		2: {{0, 1}, {1, 0}},
		3: {{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}},
	}
)

type Brand struct {
	Name    string
	Version string
}

func greaseBrand(major uint32) Brand {
	n := int(major)
	name := "Not" + string(greaseChars[n%len(greaseChars)]) + "A" + string(greaseChars[(n+1)%len(greaseChars)]) + "Brand"
	return Brand{Name: name, Version: greaseVersions[n%len(greaseVersions)]}
}

func Brands(major uint32, product string) []Brand {
	v := strconv.FormatUint(uint64(major), 10)
	list := []Brand{greaseBrand(major), {Name: "Chromium", Version: v}}
	if product != "" {
		list = append(list, Brand{Name: product, Version: v})
	}
	orders := brandOrders[len(list)]
	order := orders[int(major)%len(orders)]
	if len(list) == 2 {
		order = []int{int(major) % 2, (int(major) + 1) % 2}
	}
	out := make([]Brand, len(list))
	for i, idx := range order {
		out[idx] = list[i]
	}
	return out
}

func SecCHUA(brands []Brand) string {
	parts := make([]string, len(brands))
	for i, b := range brands {
		parts[i] = strconv.Quote(b.Name) + ";v=" + strconv.Quote(b.Version)
	}
	return strings.Join(parts, ", ")
}
