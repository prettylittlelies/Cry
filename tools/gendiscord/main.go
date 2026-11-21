package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var overrides = map[string]string{
	"*.permissions":     "Permissions",
	"*.allow":           "Permissions",
	"*.deny":            "Permissions",
	"*.permissions_new": "Permissions",
}

var enumNames = map[string]string{
	"Bitwise Permission Flags": "Permission",
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: gendiscord <userdoccers/pages> <models dir> <api dir>")
		os.Exit(2)
	}
	spec, err := ParseDir(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	n := &namer{taken: map[string]string{"Permissions": "manual", "ImageData": "manual"}}
	count := map[string]int{}
	for _, s := range spec.Structures {
		count[GoName(s.Title)]++
	}
	var enums []*namedEnum
	for _, e := range spec.Enums {
		if k := enumKind(e); k != "" {
			count[GoName(e.Title)]++
			enums = append(enums, &namedEnum{enum: e, kind: k})
		}
	}
	sort.Slice(spec.Structures, func(i, j int) bool {
		a, b := spec.Structures[i], spec.Structures[j]
		if pa, pb := priority(a.Doc), priority(b.Doc); pa != pb {
			return pa < pb
		}
		if a.Doc != b.Doc {
			return a.Doc < b.Doc
		}
		return a.Title < b.Title
	})
	names := map[*Structure]string{}
	for _, s := range spec.Structures {
		names[s] = n.assign(s.Doc, s.Title, count)
	}
	sort.Slice(enums, func(i, j int) bool {
		if enums[i].enum.Doc != enums[j].enum.Doc {
			return enums[i].enum.Doc < enums[j].enum.Doc
		}
		return enums[i].enum.Title < enums[j].enum.Title
	})
	for _, e := range enums {
		if fixed, ok := enumNames[e.enum.Title]; ok {
			e.name = fixed
			n.taken[fixed] = e.enum.Title
			continue
		}
		e.name = n.assign(e.enum.Doc, e.enum.Title, count)
	}
	sort.Slice(enums, func(i, j int) bool { return enums[i].name < enums[j].name })
	r := newResolver(spec, names, enums)
	bindUnions(enums, r)
	write := func(dir, name string, data []byte) {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	structs, fields := emitStructs(spec, r, overrides)
	write(os.Args[2], "structures.go", structs)
	write(os.Args[2], "enums.go", emitEnums(enums))
	write(os.Args[3], "endpoints.go", emitAPI(spec, r, fields))
	fmt.Printf("generated %d structures, %d enums\n", len(spec.Structures), len(enums))
}
