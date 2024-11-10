package service

import (
	"fmt"
	"strings"

	"github.com/kaibling/apiforge/params"
)

const (
	ASC  = "ASC"
	DESC = "DESC"
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
		Filter:       qp.Filter,
		Limit:        qp.Limit,
		Order:        strings.ToUpper(qp.Order),
		tableName:    tableName,
		before:       qp.Before,
		after:        qp.After,
		WhereClauses: []string{},
	}
}

func (p *Pagination) GetCursorSQL() string {
	innerOrder := p.Order
	if p.before != nil {
		innerOrder = DESC

		p.WhereClauses = append(p.WhereClauses, fmt.Sprintf("id %s '%s'", operator("before", p.Order), *p.before))
	} else if p.after != nil {
		innerOrder = ASC

		p.WhereClauses = append(p.WhereClauses, fmt.Sprintf("id %s '%s'", operator("after", p.Order), *p.after))
	}

	innerSQL := fmt.Sprintf("SELECT id as pagination_id from %s %s ORDER BY id %s LIMIT %d",
		p.tableName, p.clauses(),
		innerOrder, p.Limit+1)

	return fmt.Sprintf("SELECT pagination_id FROM (%s) Order By pagination_id %s;", innerSQL, p.Order)
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

func (p *Pagination) FinishPagination(ids []string) ([]string, params.Pagination) { //nolint: cyclop
	pag := params.Pagination{ //nolint: exhaustruct
		Limit: p.Limit,
		Order: p.Order,
	}

	//nolint: nestif
	if p.before != nil {
		if p.Order == ASC {
			pag.After = &ids[len(ids)-1]
			if len(ids) == p.Limit+1 {
				pag.Before = &ids[1]
			}
		} else {
			pag.After = &ids[1]
			if len(ids) == p.Limit+1 {
				pag.Before = &ids[len(ids)-1]
			}
		}
	} else {
		if p.Order == ASC {
			if p.after != nil {
				pag.Before = &ids[0]
			}

			if len(ids) == p.Limit+1 {
				pag.After = &ids[len(ids)-2]
			}
		} else {
			if p.after != nil {
				pag.Before = &ids[len(ids)-2]
			}

			if len(ids) == p.Limit+1 {
				pag.After = &ids[1]
			}
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

func operator(direction, order string) string {
	var op string

	if direction == "before" { //nolint: nestif
		if order == ASC {
			op = "<"
		} else {
			op = ">"
		}
	} else {
		if order == ASC {
			op = ">"
		} else {
			op = "<"
		}
	}

	return op
}
