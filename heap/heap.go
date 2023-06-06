package heap

type Heap[T any] struct {
	heap       []T
	comparator func(i, j T) bool
}

func New[T any](comparator func(i T, j T) bool) *Heap[T] {
	return &Heap[T]{heap: make([]T, 0), comparator: comparator}
}

func (h *Heap[T]) Push(item T) {
	h.heap = append(h.heap, item)

	index := len(h.heap) - 1

	for !h.comparator(h.heap[parent(index)], h.heap[index]) {
		h.heap[index], h.heap[parent(index)] = h.heap[parent(index)], h.heap[index]
		index = parent(index)
	}
}

func (h *Heap[T]) Pop() T {
	if len(h.heap) == 0 {
		panic("The heap is empty")
	}

	item := h.heap[0]
	h.heap = h.heap[1:]

	return item
}

func (h *Heap[T]) heapify(index int) {
	leftChildIndex, rightChildIndex, compared := left(index), right(index), index

	if leftChildIndex < len(h.heap) && !h.comparator(h.heap[compared], h.heap[leftChildIndex]) {
		compared = leftChildIndex
	}

	if rightChildIndex < len(h.heap) && !h.comparator(h.heap[compared], h.heap[rightChildIndex]) {
		compared = rightChildIndex
	}

	if index != compared {
		h.heap[index], h.heap[compared] = h.heap[compared], h.heap[index]
		h.heapify(compared)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}
