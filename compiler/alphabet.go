package compiler

const (
	// alphaSize is the length of alphaTable.
	alphaSize = 129
	// alphaRuleSize is the length of alphaRule.
	alphaRuleSize = 39
)

var (
	// alphaRule is the rule for the alphabet, which is used to build alphaTable.
	alphaRule = [alphaRuleSize]byte{
		0x01, 0x0A, 0x20, '`', '"',
		';', '(', ')', '[', ']',
		'{', '}', '&', '|', '#', '.', ',', '<', '>', '=',
		'!', '*', '$', '@', ':', '+', '-', '/', '\\',
		'0', '1', 'a', '_',
		'~', // todo: add support for ~ in the future
		'^', '%',
		'?', // todo: add support for ? in the future
		0x27,
		0x80,
	}
	// alphaTable is indexed by the ASCII value of a character and returns the index of the character in alphaRule.
	alphaTable [alphaSize]byte
)

func init() {
	buildAlphaTable()
}

// buildAlphaTable builds the alphaTable.
func buildAlphaTable() {
	for ind, ch := range alphaRule {
		r := byte(ind)
		switch ch {
		case 0x20:
			alphaTable[0x09] = r // Horizontal Tab, HT
			alphaTable[0x0d] = r // Carriage Return, CR
			alphaTable[0x20] = r // Space
		case '1':
			for k := '1'; k <= '9'; k++ {
				alphaTable[k] = r
			}
		case 'a':
			for k := 'a'; k <= 'z'; k++ {
				alphaTable[k] = r
				alphaTable[k-32] = r
			}
		case 0x80:
			alphaTable[0x80] = r
		default:
			alphaTable[ch] = r
		}
	}
}

// charToAlpha returns the index of the character in alphaRule.
func charToAlpha(ch rune) byte {
	if ch > 127 {
		return alphaTable[len(alphaTable)-1]
	}
	return alphaTable[ch]
}
