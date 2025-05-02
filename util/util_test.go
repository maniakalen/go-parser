package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMetaData(t *testing.T) {

	tests := []struct {
		name        string
		body        string
		title       string
		description string
	}{
		{
			name:        "Test title",
			body:        "<html><title>test title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>test h1</h1></html>",
			title:       "test title",
			description: "some description",
		},
		{
			name:        "Test title",
			body:        "<html><title>test title</title><meta name=\"description\" content=\"<p>some description</p>\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>test h1</h1></html>",
			title:       "test title",
			description: "some description",
		},
		{
			name:        "Test title",
			body:        "<html><title>test title</title><meta name=\"description\" content=\"<h3><p>some description</p></\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>test h1</h1></html>",
			title:       "test title",
			description: "some description",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title, description, err := GetPageMetaData(tt.body)
			assert.NoError(t, err)
			assert.Equal(t, tt.title, title, "Expected title to be %v, got %v", tt.title, title)
			assert.Equal(t, tt.description, description, "Expected desc to be %v, got %v", tt.description, description)
		})
	}
}
