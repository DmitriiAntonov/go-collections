package heap

// Heap implemented the binary heap data structure
type Heap[T any] struct {
	heap []T
	less func(i, j T) bool
}

// New create new instance of the heap
func New[T any](less func(i T, j T) bool) *Heap[T] {
	return &Heap[T]{heap: make([]T, 0), less: less}
}

func (h *Heap[T]) Push(item T) {
	h.heap = append(h.heap, item)
	h.up(len(h.heap) - 1)
}

func (h *Heap[T]) up(index int) {
	for index != 0 && !h.less(h.heap[parent(index)], h.heap[index]) {
		h.heap[index], h.heap[parent(index)] = h.heap[parent(index)], h.heap[index]
		index = parent(index)
	}
}

func (h *Heap[T]) Pop() T {
	if len(h.heap) == 0 {
		panic("The heap is empty")
	}

	item := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.down(0)

	return item
}

func (h *Heap[T]) down(index int) {
	compared := index

	if left(index) < len(h.heap) && !h.less(h.heap[compared], h.heap[left(index)]) {
		compared = left(index)
	}

	if right(index) < len(h.heap) && !h.less(h.heap[compared], h.heap[right(index)]) {
		compared = right(index)
	}

	if index != compared {
		h.heap[index], h.heap[compared] = h.heap[compared], h.heap[index]
		h.down(compared)
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
