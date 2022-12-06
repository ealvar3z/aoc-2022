package lib

func Clone[T any](input [][]T) [][]T {
	clone := make([][]T, len(input))
	for i, rows := range input {
		row := make([]T, len(rows))
		for j, value := range row {
			row[j] = value
		}
		clone[i] = row
	}
	return clone
}
