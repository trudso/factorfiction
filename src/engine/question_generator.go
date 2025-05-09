package engine

func GenerateQuestion() (string, error) {
ollama := NewOllama("deepseek-r1:14b")
	return ollama.Generate(`Generate 2 quotes for a game where the objective is to determine if a quote is real or fake news. 
		The result should be returned as in a json format with the following fields:
		quote: the quote to determine if it is real or fake,
		author: the author of the quote if the quote is real. If the quote is fake choose a random famouse person for the question instead,
		o
		url: the url-source of the quote if the quote is real or null if it's fake,
		answer: the answer to the question: Is the quote real or fake news?`)
}
