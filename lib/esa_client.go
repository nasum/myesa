package lib

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/upamune/go-esa/esa"
)

type EsaClient struct {
	client *esa.Client
	team   string
}

// Init is initialize
func (e *EsaClient) Init() {
	viper.SetConfigName("myesarc")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")

	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read config file: %v", err)
		os.Exit(1)
	}
	accessToken := viper.GetString("ACCESS_TOKEN")
	team := viper.GetString("TEAM")
	client := esa.NewClient(accessToken)

	e.client = client
	e.team = team
}

// Search is return articles
func (e *EsaClient) Search(query url.Values) []esa.PostResponse {
	response, err := e.client.Post.GetPosts(e.team, query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return response.Posts
}

// Create is create article
func (e *EsaClient) Create(name string, body string) string {

	post := esa.Post{
		Name:   name,
		BodyMd: body,
	}
	response, err := e.client.Post.Create(e.team, post)

	if err != nil {
		log.Fatal(err)
		return "Can`t create article."
	}

	return response.URL
}

// Update is update article
func (e *EsaClient) Update(postNumber int, name string, body string) string {
	response, err := e.client.Post.Update(e.team, postNumber, esa.Post{
		Name:   name,
		BodyMd: body,
	})

	if err != nil {
		log.Fatal(err)
		return "Can`t update article."
	}

	return response.URL
}

// GetArticlesByName is get article by name
func (e *EsaClient) GetArticlesByName(articleFullName string) []esa.PostResponse {
	queryStrArray := strings.Split(articleFullName, "/")
	articleName := queryStrArray[len(queryStrArray)-1]
	categoryName := strings.Join(queryStrArray[0:len(queryStrArray)-1], "/")
	query := url.Values{}
	query.Add("name", articleName)
	query.Add("in", categoryName)

	return e.Search(query)
}
