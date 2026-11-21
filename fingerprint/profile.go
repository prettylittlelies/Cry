package fingerprint

type Profile struct {
	Name    string
	Hello   ClientHello
	H2      H2
	Headers HeaderOrder
}

func Electron44() Profile {
	return Profile{
		Name:    "electron44",
		Hello:   Chromium152,
		H2:      Chromium152H2(),
		Headers: DiscordDesktopOrder(),
	}
}
