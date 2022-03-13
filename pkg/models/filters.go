package models

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

const (
	DefaultSortDirection = "ASC"
	DefaultPage          = 1
	PageLowerLimit       = 1
	PageUpperLimit       = 1000
	DefaultPageSize      = 10
)

type Filters struct {
	Page          int
	PageSize      int
	SortBy        string
	SortDirection string
}

func NewFilters(qs url.Values, sortableColumns map[string]struct{}, defaultSortBy string) (*Filters, error) {
	var page int
	var err error
	if qs.Has("page") {
		page, err = strconv.Atoi(qs.Get("page"))
		if err != nil || page < PageLowerLimit || page > PageUpperLimit {
			fmt.Println("hej")
			return nil, ErrInvalidFilters
		}
	} else {
		page = DefaultPage
	}

	sortBy := defaultSortBy
	sortDirection := DefaultSortDirection
	if qs.Has("sort") {
		sortBy = qs.Get("sort")
	}
	if _, ok := sortableColumns[sortBy]; !ok {
		return nil, ErrInvalidFilters
	}
	if strings.HasPrefix(sortBy, "-") {
		sortBy = strings.TrimPrefix(sortBy, "-")
		sortDirection = "DESC"
	}

	filters := &Filters{
		Page:          page,
		PageSize:      DefaultPageSize,
		SortBy:        sortBy,
		SortDirection: sortDirection,
	}

	return filters, nil
}

func (f *Filters) Limit() int {
	return f.PageSize
}

func (f *Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

type Metadata struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	LastPage     int
	TotalRecords int
}

func CalculateMetadata(totalRecords, page, pageSize int) *Metadata {
	if totalRecords == 0 {
		return &Metadata{}
	}

	return &Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
