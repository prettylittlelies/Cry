package discord

import (
	"encoding/json"
	"strconv"
)

type Permissions uint64

func (p Permissions) Has(want ...Permission) bool {
	if Permission(p)&PermissionAdministrator != 0 {
		return true
	}
	for _, w := range want {
		if Permission(p)&w != w {
			return false
		}
	}
	return true
}

func (p Permissions) With(add ...Permission) Permissions {
	for _, a := range add {
		p |= Permissions(a)
	}
	return p
}

func (p Permissions) Without(remove ...Permission) Permissions {
	for _, r := range remove {
		p &^= Permissions(r)
	}
	return p
}

func (p Permissions) String() string {
	return strconv.FormatUint(uint64(p), 10)
}

func (p Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Permissions) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		var n uint64
		if err := json.Unmarshal(b, &n); err != nil {
			return err
		}
		*p = Permissions(n)
		return nil
	}
	n, err := strconv.ParseUint(s, 10, 64)
	*p = Permissions(n)
	return err
}
