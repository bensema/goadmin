package model

import "errors"

const (
	DefaultNum  = 1
	DefaultSize = 20
)

// SimplePage calculate "from", "to" without total_counts
// "from" index start from 1
func (p *Pagination) SimplePage() (from int, to int) {
	if p.Page == 0 || p.PageSize == 0 {
		p.Page, p.PageSize = 1, DefaultSize
	}
	from = (p.Page-1)*p.PageSize + 1
	to = from + p.PageSize - 1
	return
}

// CalPage calculate "from", "to" with total_counts
// index start from 1
func (p *Pagination) CalPage(total int) (from int, to int) {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultSize
	}

	if total == 0 || total < p.PageSize*(p.Page-1) {
		return
	}
	if total <= p.PageSize {
		return 1, total
	}
	from = (p.Page-1)*p.PageSize + 1
	if (total - from + 1) < p.PageSize {
		return from, total
	}
	return from, from + p.PageSize - 1
}

// VagueOffsetLimit calculate "offset", "limit" without total_counts
func (p *Pagination) VagueOffsetLimit() (offset int, limit int) {
	from, to := p.SimplePage()
	if to == 0 || from == 0 {
		return 0, 0
	}
	return from - 1, to - from + 1
}

// OffsetLimit calculate "offset" and "start" with total_counts
func (p *Pagination) OffsetLimit(total int) (offset int, limit int) {
	from, to := p.CalPage(total)
	if to == 0 || from == 0 {
		return 0, 0
	}
	return from - 1, to - from + 1
}

func (p *Pagination) Verify() error {
	if p.Page < 0 {
		return errors.New("num error")
	} else if p.Page == 0 {
		p.Page = DefaultNum
	}
	if p.PageSize < 0 {
		return errors.New("size error")
	} else if p.PageSize == 0 {
		p.PageSize = DefaultSize
	}
	return nil
}

// Pagination perform page algorithm
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// PaginationReply Pagination Response.
type PaginationReply struct {
	PageSize  int `json:"page_size"`
	Page      int `json:"page"`
	RowsTotal int `json:"rows_total"`
}
