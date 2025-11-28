package foundation

type Node[T any] struct {
	value *T
	next  *Node[T]
	prev  *Node[T]
}

type List[T any] struct {
	begin *Node[T]
	end   *Node[T]
}

func (node *Node[T]) connect(another *Node[T]) {
	node.next = another
	if another != nil {
		another.prev = node
	}
}

func NewList[T any]() *List[T] {
	lst := new(List[T])
	lst.begin = new(Node[T])
	lst.end = new(Node[T])
	lst.begin.connect(lst.end)
	return lst
}

func (lst *List[T]) Append(value T) {
	lst.Insert(lst.end, value)
}

func (lst *List[T]) Prepend(value T) {
	lst.Insert(lst.begin, value)
}

func (lst *List[T]) Insert(node *Node[T], value T) *Node[T] {
	// Handle nil node
	if node == nil {
		return nil
	}

	newNode := new(Node[T])
	newNode.value = new(T)
	*newNode.value = value

	// Case 1: Inserting after begin (at the start of list)
	if node == lst.begin {
		next := lst.begin.next
		if next == nil {
			// Empty list - connect begin -> newNode -> end
			lst.begin.connect(newNode)
			newNode.connect(lst.end)
		} else {
			// List has elements - insert between begin and first element
			lst.begin.connect(newNode)
			newNode.connect(next)
		}
		return newNode
	}

	// Case 2: Inserting before end (at the end of list)
	if node == lst.end {
		prev := lst.end.prev
		if prev == nil || prev == lst.begin {
			// Empty list - connect begin -> newNode -> end
			lst.begin.connect(newNode)
			newNode.connect(lst.end)
		} else {
			// List has elements - insert between last element and end
			prev.connect(newNode)
			newNode.connect(lst.end)
		}
		return newNode
	}

	// Case 3: Inserting after a regular node in the middle
	prev := node.prev
	prev.connect(newNode)
	newNode.connect(node)

	return newNode
}

func (lst *List[T]) Remove(node *Node[T]) *Node[T] {
	if node == nil || node == lst.begin || node == lst.end {
		return nil
	}

	prev, next := node.prev, node.next
	prev.connect(next)

	return node
}

func (lst *List[T]) Find(value T, compare func(T, T) int) *Node[T] {
	for it := lst.begin.next; it != lst.end; it = it.next {
		if compare(*it.value, value) == 0 {
			return it
		}
	}

	return nil
}
