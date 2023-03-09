package models

import (
	"errors"
	"strconv"
	"time"

	cache "starsearch/common/redis"

	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type MovieData struct {
	Id           int64  `json:"id"`
	Crawl        string `json:"opening_crawl"`
	Name         string `json:"name"`
	CommentCount int64  `json:"comment_count"`
	ReleasedOn   time.Time
}

type MovieModel struct {
	Id         int64
	Crawl      string
	Name       string
	ReleasedOn time.Time
}

func (m *MovieModel) CommentCount(DB *gorm.DB) int64 {
	c, err := GetComments(DB, m.Id, "count")
	if err != nil {
		c = 0
	}

	return c.(int64)
}

func (m MovieModel) GreaterThan(RHS *MovieModel) bool {
	return m.ReleasedOn.After(RHS.ReleasedOn)
}

func (m MovieModel) LesserThan(RHS *MovieModel) bool {
	return !m.GreaterThan(RHS)
}

func (m *MovieModel) GetData(DB *gorm.DB) MovieData {
	return MovieData{
		Id:           m.Id,
		Name:         m.Name,
		Crawl:        m.Crawl,
		CommentCount: m.CommentCount(DB),
		// ReleasedOn:   m.ReleasedOn, // we'll remove this after test
	}
}

var MovieIdKey = "movie_ids"

// Checks if a MovieID is available in system's cache. It assumes that '/get-movies' route has already been visited before any route that needs to call this function is visited.
// Ideally, if the ids (list) is not found in the cache, we should request movies from swapi and populate cache and then validate again. So this should be our next line of improvement.
//
// Returns an error if the movie(id) is not found.
func ValidateMovieID(RDS redis.Conn, movie_id string) error {
	id, err := strconv.ParseInt(movie_id, 10, 64)
	if err != nil {
		return err
	}

	ids, err := cache.GetData[int64](RDS, MovieIdKey)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			return nil
		}

	}

	return errors.New("Movie does not exist.")
}
