package task1

import "github.com/Kupstas/elma_test/internal/list"

var bracketPairs = map[string]string{
	"[": "]",
	"(": ")",
	"{": "}",
}

func CheckBracketSequence(row string) bool {
	if row == "" {
		return false
	}

	seq := list.New[string]()
	for _, ch := range row {
		key := string(ch)
		if _, ok := bracketPairs[key]; ok {
			seq.Add(key)
			continue
		}

		if el, ok := seq.Get(); ok {
			if key == bracketPairs[*el] {
				seq.Remove()
				continue
			}
		}

		return false
	}

	return seq.Len() == 0
}
