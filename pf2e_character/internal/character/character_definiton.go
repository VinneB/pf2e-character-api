package character

type Character struct {
	Name       string `json:"Name"`
	Alignment  string
	Level      uint8
	XP         uint16
	HeroPoints uint8
	Race       string
	Class      string
	Size       string
	Gender     string
	Age        string
	Str        int8
	Dex        int8
	Con        int8
	Wis        int8
	Cha        int8
	HP         uint8
}
