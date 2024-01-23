package data

import "github.com/spobly/greenlight/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func (f Filters) Validate(v *validator.Validator) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be less than 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be less than 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}
