package sort

type Sortable interface {
	Less(i, j int) bool
	Len() int
	Swap(i, j int)
}

type Indexable interface {
	Get(i int) int
	Set(i int, v int)
}

type SortAndIndexAble interface {
	Sortable
	Indexable
}
