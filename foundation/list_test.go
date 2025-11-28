package foundation

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestInsert_AtNilNode_ReturnsNil(t *testing.T) {
	sut := NewList[int]()

	assert.Nil(t, sut.Insert(nil, 3))
}

func TestInsert_Succeeds(t *testing.T) {
	sut := NewList[int]()
	position := sut.begin
	expectedList := []int{5, 4, 3, 2, 1}
	actualList := make([]int, 0, 5)

	for i := 1; i <= 5; i++ {
		assert.NotNil(t, position)
		nextPosition := sut.Insert(position, i)
		position = nextPosition
	}

	for it := sut.begin.next; it != sut.end; it = it.next {
		actualList = append(actualList, *it.value)
	}

	for i := 0; i < 5; i++ {
		assert.Equalf(t, expectedList[i], actualList[i], "Items at %d aren't equal", i)
	}
}
