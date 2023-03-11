package models

import (
	"starsearch/common/utils"
	"strconv"
)

type SortKey int64

const (
	Name SortKey = iota
	Gender
	Height
)

var sort_key_map = map[string]SortKey{
	"name":   Name,
	"gender": Gender,
	"height": Height,
}

func (s SortKey) Str() string {
	key := "name"

	for k, v := range sort_key_map {
		if v == s {
			key = k
		}
	}

	return key
}

type Character struct {
	Data    map[string]interface{}
	SortKey SortKey
}

func (c *Character) CreateSortKey(key string) SortKey {
	if key == "" {
		key = "name"
	}

	sk := sort_key_map[key]
	c.SortKey = sk
	return c.SortKey
}

func (c Character) GreaterThan(r *Character) bool {
	right := *r

	right_data := r.Data
	right_sk := right.SortKey.Str()
	right_sort_value := right_data[right_sk].(string)

	left_data := c.Data
	left_sk := c.SortKey.Str()
	left_sort_value := left_data[left_sk].(string)

	// when comparing heights, we need to convert values to int64
	if r.SortKey == Height {
		rv, err := strconv.ParseInt(right_sort_value, 10, 64)
		lv, err := strconv.ParseInt(left_sort_value, 10, 64)

		if err != nil {
			return false
		}

		return lv > rv
	}

	return left_sort_value > right_sort_value
}

type CharacterList struct {
	CharCount   int                      `json:"total_character_count"`
	TotalHeight TotalHeight              `json:"total_character_height"`
	Characters  []map[string]interface{} `json:"characters"`
	Items       []Character              `json:"-"`
	Desc        bool                     `json:"-"`
	FilterBy    string                   `json:"-"`
}

func (cl *CharacterList) Parse(items []map[string]interface{}, sort_key string, sort_order string, filter_by string) {
	cl.FilterBy = filter_by
	cl.CharCount = len(items)
	cl.Desc = false

	if sort_order == "desc" {
		cl.Desc = true
	}

	for i := 0; i < len(items); i++ {
		item := items[i]
		ch := Character{
			Data: item,
		}

		ch.CreateSortKey(sort_key)
		cl.Items = append(cl.Items, ch)
	}

	cl.Sort()
	cl.Filter()

	cl.CharCount = len(cl.Characters)
	th := cl.GetTotalHeight()
	cl.TotalHeight.Cm = th
	cl.TotalHeight.Ft = utils.CmToFt(th)
}

func (cl *CharacterList) GetTotalHeight() int64 {
	chs := cl.Characters
	var th int64

	for i := 0; i < len(chs); i++ {
		item := chs[i]
		height := item["height"].(string)
		hi_int, _ := strconv.ParseInt(height, 10, 64)
		th += hi_int
	}

	return th
}

func (cl *CharacterList) Sort() {
	cl.Items = utils.Sort(cl.Items, 0, cl.CharCount-1, cl.Desc)
	cl.Characters = []map[string]interface{}{}

	for i := 0; i < cl.CharCount; i++ {
		item := cl.Items[i]
		cl.Characters = append(cl.Characters, item.Data)
	}
}

func (cl *CharacterList) Filter() {
	if cl.FilterBy == "" {
		return
	}

	chs := cl.Characters
	new_chs := []map[string]interface{}{}
	for j := 0; j < len(chs); j++ {
		c := chs[j]
		if c["gender"] == cl.FilterBy {
			new_chs = append(new_chs, c)
		}
	}

	cl.Characters = new_chs
}

type TotalHeight struct {
	Ft string `json:"ft"`
	Cm int64  `json:"cm"`
}
