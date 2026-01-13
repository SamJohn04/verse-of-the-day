package internal

type Result struct {
	Verse Verse `json:"verse"`
}

type Verse struct {
	Details Details `json:"details"`
	Notice  string  `json:"notice"`
}

type Details struct {
	Text      string `json:"text"`
	Reference string `json:"reference"`
	Version   string `json:"version"`
	VerseUrl  string `json:"verseurl"`
}
