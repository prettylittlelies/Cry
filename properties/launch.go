package properties

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ClientMod uint8

const (
	JQuery        ClientMod = 119
	BetterDiscord ClientMod = 108
	Rambox        ClientMod = 100
	Vendetta      ClientMod = 91
	Vencord       ClientMod = 84
	Replugged     ClientMod = 75
	Legcord       ClientMod = 61
	Dorion        ClientMod = 55
	GoofCord      ClientMod = 48
	OpenAsar      ClientMod = 38
	Shelter       ClientMod = 24
	Moonlight     ClientMod = 11
)

var clientMods = []ClientMod{JQuery, BetterDiscord, Rambox, Vendetta, Vencord, Replugged, Legcord, Dorion, GoofCord, OpenAsar, Shelter, Moonlight}

func (m ClientMod) locate() (int, byte) {
	pos := 127 - int(m)
	return pos / 8, 1 << (7 - pos%8)
}

func (m ClientMod) In(signature string) bool {
	id, err := uuid.Parse(signature)
	if err != nil {
		return false
	}
	i, bit := m.locate()
	return id[i]&bit != 0
}

func LaunchSignature() string {
	id := uuid.New()
	for _, m := range clientMods {
		i, bit := m.locate()
		id[i] &^= bit
	}
	return id.String()
}

func MobileLaunchSignature(now time.Time) string {
	return strconv.FormatInt(now.UnixNano(), 10)
}
