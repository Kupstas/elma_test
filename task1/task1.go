package task1

import "elma_test/internal"

var (
	openBrackets = map[string]struct{}{
		"[": {},
		"(": {},
		"{": {},
	}
	closeBrackets = map[string]struct{}{
		"]": {},
		")": {},
		"}": {},
	}
	bracketPairs = map[string]string{
		"[": "]",
		"(": ")",
		"{": "}",
	}
)

func CheckBracketSequence(row string) bool {
	if row == "" {
		return false
	}

	seq := internal.NewList()
	for _, ch := range row {
		key := string(ch)
		if _, ok := openBrackets[key]; ok {
			seq.PushBack(key)
			continue
		}

		if _, ok := closeBrackets[key]; ok {
			back := seq.Back()
			// closed brackets are more than opened
			if back == nil {
				return false
			}

			lastOpenBracket := seq.Back().Value
			val := bracketPairs[lastOpenBracket]
			if key == val {
				seq.Remove(seq.Back())
				continue
			}

			return false
		}
	}

	return seq.Len() == 0
}
