package build

type Channel string

const (
	Stable Channel = "stable"
	PTB    Channel = "ptb"
	Canary Channel = "canary"
)

func Channels() []Channel {
	return []Channel{Stable, PTB, Canary}
}

func (c Channel) Valid() bool {
	switch c {
	case Stable, PTB, Canary:
		return true
	}
	return false
}

func (c Channel) Subdomain() string {
	if c == Stable {
		return ""
	}
	return string(c) + "."
}
