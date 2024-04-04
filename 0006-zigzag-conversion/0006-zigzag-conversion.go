func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
    
    rows := make([][]string, numRows)
	var i, direction int
	for _, r := range s {
		if i == 0 {
			direction = 1
		} else if i == numRows-1 {
			direction = -1
		}

		rows[i] = append(rows[i], string(r))
		i += direction
	}

	str := ""
	for _, row := range rows {
		str += strings.Join(row, "")
	}
	return str
}