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
		assert.Equalf(
			t,
			expectedList[i],
			actualList[i],
			"Items at %d aren't equal", i)
	}
}

func TestInsert_AtBegin_EmptyList_InsertsElement(t *testing.T) {
	sut := NewList[int]()

	newNode := sut.Insert(sut.begin, 42)

	assert.NotNil(t, newNode)
	assert.Equal(t, 42, *newNode.value)
	assert.Equal(t, sut.begin, newNode.prev)
	assert.Equal(t, sut.end, newNode.next)
	assert.Equal(t, newNode, sut.begin.next)
	assert.Equal(t, newNode, sut.end.prev)
}

func TestInsert_AtEnd_EmptyList_InsertsElement(t *testing.T) {
	sut := NewList[int]()

	newNode := sut.Insert(sut.end, 42)

	assert.NotNil(t, newNode)
	assert.Equal(t, 42, *newNode.value)
	assert.Equal(t, sut.begin, newNode.prev)
	assert.Equal(t, sut.end, newNode.next)
	assert.Equal(t, newNode, sut.begin.next)
	assert.Equal(t, newNode, sut.end.prev)
}

func TestInsert_AtBegin_NonEmptyList_PrependsElement(t *testing.T) {
	sut := NewList[int]()
	sut.Append(10)
	sut.Append(20)

	newNode := sut.Insert(sut.begin, 5)

	assert.NotNil(t, newNode)
	assert.Equal(t, 5, *newNode.value)
	assert.Equal(t, sut.begin, newNode.prev)
	assert.Equal(t, newNode, sut.begin.next)
	assert.Equal(t, 10, *newNode.next.value)
}

func TestInsert_AtEnd_NonEmptyList_AppendsElement(t *testing.T) {
	sut := NewList[int]()
	sut.Append(10)
	sut.Append(20)

	newNode := sut.Insert(sut.end, 30)

	assert.NotNil(t, newNode)
	assert.Equal(t, 30, *newNode.value)
	assert.Equal(t, sut.end, newNode.next)
	assert.Equal(t, newNode, sut.end.prev)
	assert.Equal(t, 20, *newNode.prev.value)
}

func TestInsert_InMiddle_InsertsAtCorrectPosition(t *testing.T) {
	sut := NewList[int]()
	first := sut.Insert(sut.begin, 10)
	third := sut.Insert(sut.end, 30)

	second := sut.Insert(third, 20)

	assert.NotNil(t, second)
	assert.Equal(t, 20, *second.value)
	assert.Equal(t, first, second.prev)
	assert.Equal(t, third, second.next)
	assert.Equal(t, second, first.next)
	assert.Equal(t, second, third.prev)
}

func TestInsert_MultipleElements_MaintainsCorrectOrder(t *testing.T) {
	sut := NewList[int]()

	sut.Insert(sut.end, 1)
	sut.Insert(sut.end, 2)
	sut.Insert(sut.end, 3)

	values := make([]int, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestInsert_WithStrings_WorksCorrectly(t *testing.T) {
	sut := NewList[string]()

	node1 := sut.Insert(sut.begin, "hello")
	node2 := sut.Insert(sut.end, "world")

	assert.Equal(t, "hello", *node1.value)
	assert.Equal(t, "world", *node2.value)
	assert.Equal(t, node2, node1.next)
	assert.Equal(t, node1, node2.prev)
}

func TestInsert_ReturnsInsertedNode(t *testing.T) {
	sut := NewList[int]()

	node := sut.Insert(sut.begin, 100)

	assert.NotNil(t, node)
	assert.NotNil(t, node.value)
	assert.Equal(t, 100, *node.value)
}

func TestFind_NonExistingValue_ReturnsNil(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	assert.Nil(t, Find(sut, -1))
}

func TestFind_Succeeds(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	assert.NotNil(t, Find(sut, 3))
}
