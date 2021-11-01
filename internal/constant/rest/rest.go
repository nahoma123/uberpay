package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Links struct {
	Self     string `json:"self"`
	First    string `json:"first"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}
type Pagination struct {
	Page       uint   `form:"page" json:"page"`
	Limit      int    `form:"limit" json:"limit"`
	Sort       string `form:"sort" json:"sort"`
	OrderBy    string `form:"order_by" json:"order_by"`
	FilterBy   string `form:"filter_by" json:"filter_by"`
	Filterkey  string `form:"filter_key" json:"filter_key"`
	TotalCount uint   `json:"total_count"`
	Links      Links  `json:"links,omitempty"`
}

func ParseLinks(c *gin.Context, pagination *Pagination) *Pagination {
	queries := c.Request.URL.Query()
	filters := ""
	for key, value := range queries {
		if key == "limit" || key == "order_by" || key == "sort" || key == "page" {
			continue
		}
		filters = fmt.Sprintf("%s&%s=%s", filters, key, value[0])
		fmt.Println(filters)

	}

	// calculate the last page
	lastPage := uint(int(pagination.TotalCount) / pagination.Limit)
	if lastPage == 0 {
		lastPage = 1
	}

	links := Links{
		Self:     fmt.Sprintf("%v?limit=%v&page=%v&order_by=%v&sort=%v%v", c.Request.URL.Path, pagination.Limit, pagination.Page, pagination.OrderBy, pagination.Sort, filters),
		First:    fmt.Sprintf("%v?limit=%v&page=%v&order_by=%v&sort=%v%v", c.Request.URL.Path, pagination.Limit, 1, pagination.OrderBy, pagination.Sort, filters),
		Previous: fmt.Sprintf("%v?limit=%v&page=%v&order_by=%v&sort=%v%v", c.Request.URL.Path, pagination.Limit, pagination.Page-1, pagination.OrderBy, pagination.Sort, filters),
		Next:     fmt.Sprintf("%v?limit=%v&page=%v&order_by=%v&sort=%v%v", c.Request.URL.Path, pagination.Limit, pagination.Page+1, pagination.OrderBy, pagination.Sort, filters),
		Last:     fmt.Sprintf("%v?limit=%v&page=%v&order_by=%v&sort=%v%v", c.Request.URL.Path, pagination.Limit, lastPage, pagination.OrderBy, pagination.Sort, filters),
	}
	pagination.Links = links

	return pagination
}

type Response struct {
	MetaData interface{} `json:"meta_data,omitempty"`
	Data     interface{} `json:"data"`
}

//ResponseJson creates new json object
func ErrorResponseJson(c *gin.Context, responseData interface{}, statusCode int) {
	c.JSON(statusCode, responseData)
	return
}

//ResponseJson creates new json object
func SuccessResponseJson(c *gin.Context, metaData interface{}, responseData interface{}, statusCode int) {
	c.JSON(statusCode, Response{
		MetaData: metaData,
		Data:     responseData,
	})
	return
}
