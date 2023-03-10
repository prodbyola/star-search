package controllers

import (
	"errors"
	"fmt"

	cache "starsearch/common/redis"
	"starsearch/common/response"
	"starsearch/common/swapi"
	"starsearch/common/utils"
	"starsearch/models"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// loads the app's movies
func (h *handler) GetMovies(c *gin.Context) {
	resp := response.New()
	cache_key := "star_movies"

	// get movie list from redis cache...
	movies, err := cache.GetData[models.MovieModel](h.RDS, cache_key)
	// err = cache.ClearData(h.RDS, cache_key)
	if err != nil {
		fmt.Print("Cache error ", err)
		// ... if movie list is not found in cache, get it from swapi and add it to cache
		if errors.Is(redis.ErrNil, err) {
			fmt.Print("Not in cache")
			if movies, ids, err := swapi.GetMovies(h.RDS); err == nil {
				err = cache.SetData(h.RDS, cache_key, movies)
				err = cache.SetData(h.RDS, models.MovieIdKey, ids)

				if err != nil {
					resp.SetError(500, err)
					c.JSON(200, resp)
					return
				}
			}
		}
	}

	movies = utils.Sort(movies, 0, len(movies)-1, false)

	// serialize the data
	data := []models.MovieData{}
	for i := 0; i < len(movies); i++ {
		movie := movies[i]
		data = append(data, movie.GetData(h.DB))
	}

	resp.SetData(data)
	c.JSON(200, resp)
}

// loads characters of a movie by a given parameter 'movie_id'.
func (h *handler) GetCharacters(c *gin.Context) {
	movie := c.Param("movie_id") // validate movie id
	resp := response.New()

	err := models.ValidateMovieID(h.RDS, movie)
	if err != nil {
		resp.SetError(404, err)
		c.JSON(200, resp)
		return
	}

	cache_key := "chars_" + movie
	characters, err := cache.GetData[map[string]interface{}](h.RDS, cache_key)
	if err != nil {
		if errors.Is(redis.ErrNil, err) {
			if characters, err = swapi.GetCharacters[map[string]interface{}](h.RDS, movie); err == nil {
				cache.SetData(h.RDS, cache_key, characters)
			}
		}
	}

	sort_by := c.Query("sort_by")
	sort_order := c.Query("sort_order")
	filter_by := c.Query("filter_by")
	chars := models.CharacterList{}
	chars.Parse(characters, sort_by, sort_order, filter_by)

	if err != nil {
		resp.SetError(500, err)
		c.JSON(200, resp)
		return
	}

	resp.SetData(chars)
	c.JSON(200, resp)
}
