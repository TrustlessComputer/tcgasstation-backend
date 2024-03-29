package entity

import (
	"encoding/json"
	"net/http"
	"strings"
	"tcgasstation-backend/utils"
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IEntity interface {
	CollectionName() string
	ToBson() (*bson.D, error)
	SetCreatedAt()
	SetUpdatedAt()
	SetDeletedAt()
	SetID()
	GetID() string
	Decode(from *primitive.D) error
}

func (b *BaseEntity) SetID() {
	b.ID = primitive.NewObjectID()
	b.UUID = b.ID.Hex()
}
func (b *BaseEntity) GetID() string {
	return b.UUID
}

func (b *BaseEntity) Decode(from *primitive.D) error {
	err := helpers.Transform(from, b)
	if err != nil {
		return err
	}
	return nil
}

//

type SortType int

const (
	SORT_ASC  SortType = 1
	SORT_DESC SortType = -1
)

type BaseFilters struct {
	Page   int64
	Limit  int64
	SortBy string
	Sort   SortType
}

type Pagination struct {
	Result    interface{} `json:"result" query:"-"`
	Page      int64       `json:"page" query:"page"`
	PageSize  int64       `json:"pageSize" query:"limit"`
	Total     int64       `json:"total" query:"-"`
	TotalPage int64       `json:"totalPage" query:"-"`
	Cursor    string      `json:"cursor" query:"cursor"`
	Sort      []string    `json:"sort" query:"sort"`
	Sorts     []*Sort     `json:"sorts" query:"-"`
}

func (m Pagination) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Pagination) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *Pagination) SetTotalPage() {
	if m.Total%m.PageSize == 0 {
		m.TotalPage = m.Total / m.PageSize
		return
	}
	m.TotalPage = (m.Total / m.PageSize) + 1
}

type Sort struct {
	Field string   `json:"field"`
	Type  SortType `json:"type"`
}

func NewDefaultPagination(opts ...int64) *Pagination {
	page := int64(1)
	limit := int64(10)
	if len(opts) > 0 && opts[0] > 0 {
		page = opts[0]
	}
	if len(opts) > 1 && opts[1] > 0 {
		limit = opts[1]
	}
	return &Pagination{
		PageSize: limit,
		Page:     page,
	}
}

// Input: page=1&limit=20&sort[]=id,asc&sort[]=created_at,desc
// Or: page=1&limit=20&sort=id,asc&sort=created_at,desc
//
//	Out: &Pagination{
//				Page:     1,
//				PageSize: 20,
//				Sorts: []*Sort{
//					{"id", 1},
//					{"created_at", -1},
//				},
//			}
func GetPagination(r *http.Request) *Pagination {
	pag := NewDefaultPagination()
	if err := utils.QueryParser(r, pag); err == nil && len(pag.Sort) > 0 {
		sortMap := make(map[string]string, 0)
		for i, sort := range pag.Sort {
			if i%2 != 0 {
				sortMap[pag.Sort[i-1]] = sort
			}
		}
		pag.Sorts = make([]*Sort, 0, len(sortMap))
		for key, val := range sortMap {
			pag.Sorts = append(pag.Sorts, &Sort{
				Field: key,
				Type:  getSortType(val),
			})
		}
	}
	if pag.Page <= 0 {
		pag.Page = 1
	}
	if pag.PageSize <= 0 {
		pag.PageSize = 10
	}
	return pag
}

func getSortType(sortType string) SortType {
	switch strings.ToLower(sortType) {
	case "desc":
		return SORT_DESC
	case "asc":
		return SORT_ASC
	default:
		return SORT_DESC
	}
}
