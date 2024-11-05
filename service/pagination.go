package service

import (
	"fmt"
	"strings"

	"github.com/kaibling/apiforge/params"
)

type Pagination struct {
	Limit        int
	Filter       string
	Order        string
	WhereClauses []string
	tableName    string
	before       *string
	after        *string
}

func NewPagination(qp params.Pagination, tableName string) Pagination {
	return Pagination{
		Limit:     qp.Limit,
		Order:     qp.Order,
		tableName: tableName,
		before:    qp.Before,
		after:     qp.After,
	}
}

func (p *Pagination) GetCursorSQL() string {
	if p.before != nil {
		p.WhereClauses = append(p.WhereClauses, fmt.Sprintf("id < '%s'", *p.before))
	} else if p.after != nil {
		p.WhereClauses = append(p.WhereClauses, fmt.Sprintf("id > '%s'", *p.after))
	}
	p.WhereClauses = append(p.WhereClauses, fmt.Sprintf(" ORDER BY id LIMIT %d", p.Limit+1))
	return fmt.Sprintf("SELECT id as pagination_id from %s %s ;", p.tableName, p.clauses())
}

func (p *Pagination) clauses() string {
	var clause strings.Builder
	if len(p.WhereClauses) > 0 {
		clause.WriteString("WHERE ")
	}
	for _, c := range p.WhereClauses {
		clause.WriteString(c)
	}
	return clause.String()

}

func (p *Pagination) FinishPagination(ids []string) ([]string, params.Pagination) {

	pag := params.Pagination{
		Limit: p.Limit,
		Order: p.Order,
	}

	if p.before != nil {
		pag.After = &ids[len(ids)-1]
		if len(ids) == p.Limit+1 {
			pag.Before = &ids[1]
		}
	} else {
		if p.after != nil {
			pag.Before = &ids[0]
		}

		if len(ids) == p.Limit+1 {
			pag.After = &ids[len(ids)-2]
		}
	}

	if len(ids) == p.Limit+1 {
		if p.before != nil {
			ids = ids[1:]
		} else {
			ids = ids[:len(ids)-1]
		}
	}

	return ids, pag
}
