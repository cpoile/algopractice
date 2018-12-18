package allergies

var allergens = [...]string{"eggs", "peanuts", "shellfish", "strawberries",
	"tomatoes", "chocolate", "pollen", "cats"}

var allergenCodes = map[string]byte{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128}

func Allergies2(input uint) []string {
	ret := make([]string, 0, 8)
	code := byte(input)

	for i := 0; code > 0; i++ {
		if code&1 == 1 {
			ret = append(ret, allergens[i])
		}
		code >>= 1
	}
	return ret
}

func AllergicTo2(code uint, allergen string) bool {
	return byte(code)&allergenCodes[allergen] > 0
}
