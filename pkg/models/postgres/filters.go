package postgres

type Filters struct {
	Page          int
	PageSize      int
	SortBy        string
	SortDirection string
}

func NewFilters(page int, sortBy string, sortDirection string) Filters {
	return Filters{
		Page:          page,
		PageSize:      10,
		SortBy:        sortBy,
		SortDirection: sortDirection,
	}
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}
