package swapi

import (
	"encoding/json"
	"net/http"
	"starsearch/models"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var base_url = "https://swapi.dev/api/films"

func swapiClient[T any](url string) (T, error) {
	var data T
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return data, err
	}

	json.NewDecoder(res.Body).Decode(&data)

	return data, nil
}

func GetMovies(RDS redis.Conn) ([]models.MovieModel, []int64, error) {
	var movies []models.MovieModel
	var ids []int64
	data, err := swapiClient[map[string]interface{}](base_url)
	if err != nil {
		return movies, ids, err
	}

	results := data["results"].([]interface{})
	for i := 0; i < len(results); i++ {
		result := results[i].(map[string]interface{})
		url := result["url"].(string)
		id, err := extractID(url)
		if err != nil {
			return []models.MovieModel{}, []int64{}, err
		}

		ids = append(ids, id)

		crawl := result["opening_crawl"].(string)
		date := result["release_date"].(string)
		title := result["title"].(string)
		time, _ := time.Parse("2006-01-02", date)
		movie := models.MovieModel{
			Id:         id,
			Crawl:      crawl,
			Name:       title,
			ReleasedOn: time,
		}

		movies = append(movies, movie)
	}

	if err != nil {
		return []models.MovieModel{}, []int64{}, err
	}

	return movies, ids, nil
}

func GetCharacters[T any](RDS redis.Conn, movie_id string) ([]T, error) {
	movie, err := swapiClient[map[string]interface{}](base_url + "/" + movie_id)
	var characters []T
	if err != nil {
		return characters, err
	}

	// // extract characters as list of urls
	links := movie["characters"].([]interface{})

	for i := 0; i < len(links); i++ {
		link := links[i].(string)

		ch, err := swapiClient[T](link)
		if err != nil {
			characters = []T{}
			return characters, err
		}

		characters = append(characters, ch)
	}

	return characters, nil
}

func extractID(url string) (int64, error) {
	base_len := len(base_url) + 1 // length of base url + trailing slash
	id_str := url[base_len:]
	id_str = id_str[:len(id_str)-1]

	return strconv.ParseInt(id_str, 10, 64)
}
