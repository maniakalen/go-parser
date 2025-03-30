package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericParser(t *testing.T) {

	tests := []struct {
		name     string
		body     string
		keywords []string
		result   bool
	}{
		{
			name:     "Test title",
			body:     "<html><title>test title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>test h1</h1></html>",
			keywords: []string{"test", "title"},
			result:   true,
		},
		{
			name:     "Some mitle",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>some h1</h1></html>",
			keywords: []string{"mitle"},
			result:   false,
		},
		{
			name:     "Some title",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>some h1</h1></html>",
			keywords: []string{"test", "title"},
			result:   true,
		},
		{
			name:     "Some description",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>some h1</h1></html>",
			keywords: []string{"description"},
			result:   true,
		},
		{
			name:     "Some keywords",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>some h1</h1></html>",
			keywords: []string{"keywords"},
			result:   true,
		},
		{
			name:     "Some h1",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1>some hashone</h1></html>",
			keywords: []string{"hashone"},
			result:   true,
		},
		{
			name:     "None h1",
			body:     "<html><title>some title</title><meta name=\"description\" content=\"some description\" /><meta name=\"keywords\" content=\"some keywords\" /><h1></h1></html>",
			keywords: []string{"hashone"},
			result:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := ParseHtml(tt.body, tt.keywords)
			assert.Equal(t, tt.result, result, "Expected result to be %v, got %v", tt.result, result)
		})
	}
}
