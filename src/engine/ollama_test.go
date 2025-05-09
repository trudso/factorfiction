package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	ollama := NewOllama("qwen:32b")
	reply, err := ollama.Generate("Generate a question for a game where the objective is to determine if a quote is true or fake news")
	
	assert.NoError(t, err)
	assert.NotEmpty(t, reply)
}
