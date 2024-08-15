package common

type SortingDirection string

const (
	Asc  = "asc"
	Desc = "desc"
)

type Sorting struct {
	Direction SortingDirection
}
