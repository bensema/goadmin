package model

import "errors"

const (
	DefaultNum  = 1
	DefaultSize = 20
)

// SimplePage calculate "from", "to" without total_counts
// "from" index start from 1
func (p *Pagination) SimplePage() (from int, to int) {
	if p.Num == 0 || p.Size == 0 {
		p.Num, p.Size = 1, DefaultSize
	}
	from = (p.Num-1)*p.Size + 1
	to = from + p.Size - 1
	return
}

// Page calculate "from", "to" with total_counts
// index start from 1
func (p *Pagination) Page(total int) (from int, to int) {
	if p.Num == 0 {
		p.Num = 1
	}
	if p.Size == 0 {
		p.Size = DefaultSize
	}

	if total == 0 || total < p.Size*(p.Num-1) {
		return
	}
	if total <= p.Size {
		return 1, total
	}
	from = (p.Num-1)*p.Size + 1
	if (total - from + 1) < p.Size {
		return from, total
	}
	return from, from + p.Size - 1
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
	from, to := p.Page(total)
	if to == 0 || from == 0 {
		return 0, 0
	}
	return from - 1, to - from + 1
}

func (p *Pagination) Verify() error {
	if p.Num < 0 {
		return errors.New("num error")
	} else if p.Num == 0 {
		p.Num = DefaultNum
	}
	if p.Size < 0 {
		return errors.New("size error")
	} else if p.Size == 0 {
		p.Size = DefaultSize
	}
	return nil
}

// Pagination perform page algorithm
type Pagination struct {
	Num  int `json:"num"`
	Size int `json:"size"`
}

// PaginationReply Pagination Response.
type PaginationReply struct {
	Size  int `json:"size"`
	Num   int `json:"num"`
	Total int `json:"total"`
}
