package controllers

import (
	"errors"
	"starsearch/common/response"
	"starsearch/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	Movie   int64  `json:"movie_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *handler) AddComment(c *gin.Context) {
	resp := response.New()
	data := CommentRequest{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		resp.SetError(406, err)
		c.JSON(200, resp)
		return
	}

	id := strconv.Itoa(int(data.Movie))
	err = models.ValidateMovieID(h.RDS, id)
	if err != nil {
		resp.SetError(404, err)
		c.JSON(200, resp)
		return
	}

	if len(data.Content) > 500 {
		resp.SetError(417, errors.New("Maximum length of characters exceeded. Please ensure your comment is not more than 500 characters."))
		c.JSON(200, resp)
		return
	}

	comment := models.Comment{
		Content: data.Content,
		Movie:   data.Movie,
		Author:  c.RemoteIP(),
	}

	err = comment.Save(h.DB)
	if err != nil {
		resp.SetError(500, err)
	} else {
		resp.SetData(comment.GetData())
	}

	c.JSON(200, resp)
}

func (h *handler) GetMovieComments(c *gin.Context) {
	resp := response.New()
	movie := c.Param("movie_id")
	err := models.ValidateMovieID(h.RDS, movie)
	if err != nil {
		resp.SetError(404, err)
		c.JSON(200, resp)
		return
	}

	movie_id, err := strconv.ParseInt(movie, 10, 64)
	if err != nil {
		resp.SetError(417, errors.New("Unable to parse movie id"))
		c.JSON(200, resp)
		return
	}

	comments, err := models.GetComments(h.DB, movie_id, "data")
	if err != nil {
		resp.SetError(500, err)
	} else {
		resp.SetData(comments)
	}

	c.JSON(200, resp)
}
