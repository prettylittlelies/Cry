package build

import "time"

type Windows struct {
	Host   Version
	Native NativeNumber
}

type MacOS struct {
	Host     Version
	Electron Version
	Chrome   Version
}

type Release struct {
	Channel Channel
	Number  Number
	Windows Windows
	MacOS   MacOS
}

type Mobile struct {
	Android        MobileVersion
	AndroidMinimum uint32
	IOS            MobileVersion
	IOSMinimumOS   Version
}

type Snapshot struct {
	FetchedAt time.Time
	Releases  map[Channel]Release
	Mobile    Mobile
}

func (s *Snapshot) Release(c Channel) (Release, bool) {
	r, ok := s.Releases[c]
	return r, ok
}
