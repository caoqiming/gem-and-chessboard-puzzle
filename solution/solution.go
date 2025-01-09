package solution

import (
	"fmt"
)

type PuzzleSolution struct {
	BoardLength int
	Data        []int
}

func NewPuzzleSolution(boardLength int) *PuzzleSolution {
	allowed := false
	for _, l := range AllowedBoardLength {
		if boardLength == l {
			allowed = true
			break
		}
	}
	if !allowed {
		panic(fmt.Sprintf("boardLength %d is not allowed", boardLength))
	}

	return &PuzzleSolution{
		BoardLength: boardLength,
		Data:        make([]int, boardLength*boardLength),
	}
}

func (s *PuzzleSolution) SetData(data []int) {
	if len(data) != len(s.Data) {
		panic(fmt.Sprintf("data length not match, want %d got %d", len(s.Data), len(data)))
	}

	copy(s.Data, data)
}

func (s *PuzzleSolution) GetGemLocation() int {
	// n is the number of check codes,
	// that is, the logarithm of the total number of chess pieces to base 2
	n := 0
	dataLength := len(s.Data)
	for dataLength > 1 {
		dataLength >>= 1
		n++
	}

	// Index of the check digit
	checkCodeIndexes := make([]int, n)
	for i := 0; i < n; i++ {
		checkCodeIndexes[i] = 1 << i
	}

	expectValidationCodes := make([]int, n)
	// ignore the first digit
	for i := 1; i < len(s.Data); i++ {
		isCheckBit := false
		for _, one := range checkCodeIndexes {
			if i == one {
				isCheckBit = true
				break
			}
		}

		if isCheckBit || s.Data[i] == 0 {
			continue
		}

		for p := 0; p < n; p++ {
			if i>>p&1 == 1 {
				expectValidationCodes[p] ^= 1
			}
		}
	}

	gotValidationCodes := make([]int, n)
	for i, index := range checkCodeIndexes {
		gotValidationCodes[i] = s.Data[index]
	}

	result := 0
	for i := 0; i < n; i++ {
		result += (expectValidationCodes[i] ^ gotValidationCodes[i]) << i
	}

	return result
}

func (s *PuzzleSolution) GetFlipLocation(gemLocation int) int {
	if gemLocation >= len(s.Data) || gemLocation < 0 {
		panic("invalid gemLocation")
	}
	return s.GetGemLocation() ^ gemLocation
}
