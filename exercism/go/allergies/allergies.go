package allergies

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var strToConst = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

var order = []string{"eggs", "peanuts", "shellfish", "strawberries",
	"tomatoes", "chocolate", "pollen", "cats"}

func Allergies(score uint) []string {
	var ret []string
	for _, v := range order {
		if AllergicTo(score, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func AllergicTo(score uint, substance string) bool {
	return strToConst[substance]&score != 0
}
