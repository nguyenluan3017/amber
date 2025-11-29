package foundation

import (
	"strings"
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

func TestRemove_NilNode_ReturnsNil(t *testing.T) {
	sut := NewList[int]()
	sut.Append(1)

	result := sut.Remove(nil)

	assert.Nil(t, result)
}

func TestRemove_BeginNode_ReturnsNil(t *testing.T) {
	sut := NewList[int]()
	sut.Append(1)

	result := sut.Remove(sut.begin)

	assert.Nil(t, result)
}

func TestRemove_EndNode_ReturnsNil(t *testing.T) {
	sut := NewList[int]()
	sut.Append(1)

	result := sut.Remove(sut.end)

	assert.Nil(t, result)
}

func TestRemove_OnlyElement_RemovesSuccessfully(t *testing.T) {
	sut := NewList[int]()
	node := sut.Insert(sut.begin, 42)

	result := sut.Remove(node)

	assert.NotNil(t, result)
	assert.Equal(t, 42, *result.value)
	assert.Equal(t, sut.end, sut.begin.next)
	assert.Equal(t, sut.begin, sut.end.prev)
}

func TestRemove_FirstElement_RemovesCorrectly(t *testing.T) {
	sut := NewList[int]()
	first := sut.Insert(sut.begin, 10)
	second := sut.Insert(sut.end, 20)
	sut.Insert(sut.end, 30)

	result := sut.Remove(first)

	assert.NotNil(t, result)
	assert.Equal(t, 10, *result.value)
	assert.Equal(t, second, sut.begin.next)
	assert.Equal(t, sut.begin, second.prev)
}

func TestRemove_LastElement_RemovesCorrectly(t *testing.T) {
	sut := NewList[int]()
	sut.Insert(sut.end, 10)
	second := sut.Insert(sut.end, 20)
	third := sut.Insert(sut.end, 30)

	result := sut.Remove(third)

	assert.NotNil(t, result)
	assert.Equal(t, 30, *result.value)
	assert.Equal(t, sut.end, second.next)
	assert.Equal(t, second, sut.end.prev)
}

func TestRemove_MiddleElement_RemovesCorrectly(t *testing.T) {
	sut := NewList[int]()
	first := sut.Insert(sut.end, 10)
	second := sut.Insert(sut.end, 20)
	third := sut.Insert(sut.end, 30)

	result := sut.Remove(second)

	assert.NotNil(t, result)
	assert.Equal(t, 20, *result.value)
	assert.Equal(t, third, first.next)
	assert.Equal(t, first, third.prev)
}

func TestRemove_MultipleElements_MaintainsListIntegrity(t *testing.T) {
	sut := NewList[int]()
	node1 := sut.Insert(sut.end, 1)
	node2 := sut.Insert(sut.end, 2)
	node3 := sut.Insert(sut.end, 3)
	node4 := sut.Insert(sut.end, 4)

	sut.Remove(node2)
	sut.Remove(node4)

	values := make([]int, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, []int{1, 3}, values)
	assert.Equal(t, node3, node1.next)
	assert.Equal(t, node1, node3.prev)
}

func TestFind_EmptyList_ReturnsNil(t *testing.T) {
	sut := NewList[int]()

	result := sut.Find(42, func(a *int, b *int) int {
		return *a - *b
	})

	assert.Nil(t, result)
}

func TestFind_NonExistingValue_ReturnsNil(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	result := sut.Find(-1, func(a *int, b *int) int {
		return *a - *b
	})

	assert.Nil(t, result)
}

func TestFind_FirstElement_ReturnsCorrectNode(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	result := sut.Find(0, func(a *int, b *int) int {
		return *a - *b
	})

	assert.NotNil(t, result)
	assert.Equal(t, 0, *result.value)
	assert.Equal(t, sut.begin, result.prev)
}

func TestFind_LastElement_ReturnsCorrectNode(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	result := sut.Find(4, func(a *int, b *int) int {
		return *a - *b
	})

	assert.NotNil(t, result)
	assert.Equal(t, 4, *result.value)
	assert.Equal(t, sut.end, result.next)
}

func TestFind_MiddleElement_ReturnsCorrectNode(t *testing.T) {
	sut := NewList[int]()

	for value := range 5 {
		sut.Append(value)
	}

	result := sut.Find(3, func(a *int, b *int) int {
		return *a - *b
	})

	assert.NotNil(t, result)
	assert.Equal(t, 3, *result.value)
	assert.NotEqual(t, sut.begin, result.prev)
	assert.NotEqual(t, sut.end, result.next)
}

func TestFind_WithStrings_UsesCustomComparator(t *testing.T) {
	sut := NewList[string]()
	sut.Append("apple")
	sut.Append("banana")
	sut.Append("cherry")

	result := sut.Find("banana", func(a *string, b *string) int {
		if *a == *b {
			return 0
		}

		return 1
	})

	assert.NotNil(t, result)
	assert.Equal(t, "banana", *result.value)
}

func TestFind_WithCustomComparator_CaseInsensitive(t *testing.T) {
	sut := NewList[string]()
	sut.Append("Hello")
	sut.Append("World")

	result := sut.Find("HELLO", func(a *string, b *string) int {
		aLower := strings.ToLower(*a)
		bLower := strings.ToLower(*b)
		if aLower == bLower || (len(*a) == len(*b) && a != b) {
			return 0
		}
		if aLower < bLower {
			return -1
		}
		return 1
	})

	assert.NotNil(t, result)
	assert.Equal(t, "Hello", *result.value)
}

func TestFind_DuplicateValues_ReturnsFirstMatch(t *testing.T) {
	sut := NewList[int]()
	first := sut.Insert(sut.end, 5)
	sut.Insert(sut.end, 10)
	sut.Insert(sut.end, 5)

	result := sut.Find(5, func(a *int, b *int) int {
		return *a - *b
	})

	assert.Equal(t, first, result)
}

// Test Value()
func TestValue_ReturnsNodeValue(t *testing.T) {
	sut := NewList[int]()
	node := sut.Insert(sut.begin, 42)

	result := node.Value()

	assert.NotNil(t, result)
	assert.Equal(t, 42, *result)
}

func TestValue_NilValueNode_ReturnsNil(t *testing.T) {
	node := &Node[int]{}

	result := node.Value()

	assert.Nil(t, result)
}

func TestValue_WithString_ReturnsCorrectValue(t *testing.T) {
	sut := NewList[string]()
	node := sut.Insert(sut.begin, "hello")

	result := node.Value()

	assert.NotNil(t, result)
	assert.Equal(t, "hello", *result)
}

// Test NewListOf
func TestNewListOf_EmptyVariadic_CreatesEmptyList(t *testing.T) {
	sut := NewListOf[int]()

	assert.NotNil(t, sut)
	assert.Equal(t, sut.end, sut.begin.next)
	assert.Equal(t, sut.begin, sut.end.prev)
}

func TestNewListOf_SingleValue_CreatesListWithOneElement(t *testing.T) {
	sut := NewListOf(42)

	values := make([]int, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, []int{42}, values)
}

func TestNewListOf_MultipleValues_CreatesListInCorrectOrder(t *testing.T) {
	sut := NewListOf(1, 2, 3, 4, 5)

	values := make([]int, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, values)
}

func TestNewListOf_WithStrings_CreatesCorrectList(t *testing.T) {
	sut := NewListOf("apple", "banana", "cherry")

	values := make([]string, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, []string{"apple", "banana", "cherry"}, values)
}

func TestNewListOf_WithStructs_CreatesCorrectList(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	sut := NewListOf(
		Person{"Alice", 30},
		Person{"Bob", 25},
	)

	values := make([]Person, 0)
	for it := sut.begin.next; it != sut.end; it = it.next {
		values = append(values, *it.value)
	}

	assert.Equal(t, 2, len(values))
	assert.Equal(t, "Alice", values[0].Name)
	assert.Equal(t, "Bob", values[1].Name)
}

// Test Find[T comparable]
func TestFindComparable_EmptyList_ReturnsNil(t *testing.T) {
	sut := NewList[int]()

	result := Find(sut, 42)

	assert.Nil(t, result)
}

func TestFindComparable_NonExistingValue_ReturnsNil(t *testing.T) {
	sut := NewListOf(1, 2, 3, 4, 5)

	result := Find(sut, 10)

	assert.Nil(t, result)
}

func TestFindComparable_ExistingValue_ReturnsCorrectNode(t *testing.T) {
	sut := NewListOf(10, 20, 30, 40, 50)

	result := Find(sut, 30)

	assert.NotNil(t, result)
	assert.Equal(t, 30, *result.value)
}

func TestFindComparable_FirstElement_ReturnsCorrectNode(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	result := Find(sut, 1)

	assert.NotNil(t, result)
	assert.Equal(t, 1, *result.value)
	assert.Equal(t, sut.begin, result.prev)
}

func TestFindComparable_LastElement_ReturnsCorrectNode(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	result := Find(sut, 3)

	assert.NotNil(t, result)
	assert.Equal(t, 3, *result.value)
	assert.Equal(t, sut.end, result.next)
}

func TestFindComparable_WithStrings_FindsCorrectValue(t *testing.T) {
	sut := NewListOf("apple", "banana", "cherry")

	result := Find(sut, "banana")

	assert.NotNil(t, result)
	assert.Equal(t, "banana", *result.value)
}

func TestFindComparable_DuplicateValues_ReturnsFirstMatch(t *testing.T) {
	sut := NewList[int]()
	first := sut.Insert(sut.end, 5)
	sut.Insert(sut.end, 10)
	sut.Insert(sut.end, 5)

	result := Find(sut, 5)

	assert.Equal(t, first, result)
}

func TestFindComparable_WithFloat_WorksCorrectly(t *testing.T) {
	sut := NewListOf(1.5, 2.7, 3.14)

	result := Find(sut, 2.7)

	assert.NotNil(t, result)
	assert.Equal(t, 2.7, *result.value)
}

// Test Begin()
func TestBegin_EmptyList_ReturnsBeginSentinel(t *testing.T) {
	sut := NewList[int]()

	result := sut.Begin()

	assert.NotNil(t, result)
	assert.Nil(t, result.value)
	assert.Equal(t, sut.end, result.next)
}

func TestBegin_NonEmptyList_ReturnsBeginSentinel(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	result := sut.Begin()

	assert.NotNil(t, result)
	assert.Nil(t, result.value)
	assert.Equal(t, 1, *result.next.value)
}

func TestBegin_ConsistentReference_ReturnsSameNode(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	first := sut.Begin()
	second := sut.Begin()

	assert.Equal(t, first, second)
}

// Test End()
func TestEnd_EmptyList_ReturnsEndSentinel(t *testing.T) {
	sut := NewList[int]()

	result := sut.End()

	assert.NotNil(t, result)
	assert.Nil(t, result.value)
	assert.Equal(t, sut.begin, result.prev)
}

func TestEnd_NonEmptyList_ReturnsEndSentinel(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	result := sut.End()

	assert.NotNil(t, result)
	assert.Nil(t, result.value)
	assert.Equal(t, 3, *result.prev.value)
}

func TestEnd_ConsistentReference_ReturnsSameNode(t *testing.T) {
	sut := NewListOf(1, 2, 3)

	first := sut.End()
	second := sut.End()

	assert.Equal(t, first, second)
}

func TestBeginEnd_DifferentNodes(t *testing.T) {
	sut := NewList[int]()

	assert.NotEqual(t, sut.Begin(), sut.End())
}
