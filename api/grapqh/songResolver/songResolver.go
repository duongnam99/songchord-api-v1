package songResolver

import (
	"context"
	"songchord-api/models"
	"songchord-api/repository/songRepo"

	"github.com/graphql-go/graphql"
)

var commentType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "Comment",
	Description: "Comments for song",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"content": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type:        graphql.NewList(commentType),
				Description: "The list of comment",
			},
		},
	},
)
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"song": &graphql.Field{
				Type:        productType,
				Description: "Get song by title",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					name, ok := p.Args["title"].(string)
					if ok {
						// Find product
						result = songRepo.GetSongByName(context.Background(), name)
					}
					return result, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get song list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = songRepo.GetSongList(context.Background(), limit)
					return result, nil
				},
			},
		},
	})
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new song",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"category": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				song := models.Song{
					Title:    params.Args["title"].(string),
					Content:  params.Args["content"].(string),
					Author:   params.Args["author"].(string),
					Category: params.Args["category"].(string),
				}
				if err := songRepo.InsertSong(context.Background(), song); err != nil {
					return nil, err
				}
				return song, nil
			},
		},
		"update": &graphql.Field{
			Type:        productType,
			Description: "Update song by title",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"category": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"comments": &graphql.ArgumentConfig{
					Type: commentType,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				song := models.Song{}
				if name, nameOk := params.Args["title"].(string); nameOk {
					song.Title = name
				}
				if content, contentOk := params.Args["content"].(string); contentOk {
					song.Content = content
				}
				if author, authorOk := params.Args["author"].(string); authorOk {
					song.Author = author
				}
				if category, categoryOk := params.Args["description"].(string); categoryOk {
					song.Category = category
				}
				if comments, commentsOK := params.Args["comments"].(map[string]interface{}); commentsOK {
					song.Comment = append(song.Comment, models.Comment{
						Name:    comments["name"].(string),
						Email:   comments["email"].(string),
						Content: comments["content"].(string),
					})
				}
				// log.Fatalln(song.Comment)

				if err := songRepo.UpdateSong(context.Background(), song); err != nil {
					return nil, err
				}
				return song, nil
			},
		},
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete song by name",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["title"].(string)
				if err := songRepo.DeleteSong(context.Background(), name); err != nil {
					return nil, err
				}
				return name, nil
			},
		},
	},
})

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

/*
query {
  list(limit: 1) {
    title
    content
    author
    category
  }
}

mutation {
  create(title: "Minhnam", content: "this is content", author: "Harry", category: "nhac vang") {
    title
    content
    author
    category
  }
}

*/
