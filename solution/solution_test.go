package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	for _, boardLength := range AllowedBoardLength {
		s := NewPuzzleSolution(boardLength)
		initialData := GenerateRandomSlice(boardLength * boardLength)

		for want := 0; want < boardLength*boardLength; want++ {
			s.SetData(initialData)
			flipLocation := s.GetFlipLocation(want)
			FlipedData := Flip(flipLocation, initialData)
			s.SetData(FlipedData)
			got := s.GetGemLocation()
			assert.Equal(t, want, got)
		}
	}
}
