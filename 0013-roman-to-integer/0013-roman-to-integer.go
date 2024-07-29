func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	res := 0
	maxIdx := len(s) + 1
	for i := 0; i < len(s); i++ {
		if i+2 < maxIdx {
			roman := s[i : i+2]
			if v, ok := m[roman]; ok {
				res += v
				i++
				continue
			}
		}

		roman := string(s[i])
		res += m[roman]
	}

	return res
}