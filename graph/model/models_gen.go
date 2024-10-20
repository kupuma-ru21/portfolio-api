// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type App struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	LinkType    string `json:"linkType"`
	ImageURL    string `json:"imageUrl"`
}

type Mutation struct {
}

type NewApp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	LinkType    string `json:"linkType"`
	ImageURL    string `json:"imageUrl"`
}

type Query struct {
}
