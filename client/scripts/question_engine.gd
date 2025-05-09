extends Node

class Question:
	var quote: String
	var author: String
	var answer: bool
	var url: String
	
class Response:
	var answer: bool

func is_finished(questions: Array[Question], responses: Array[Response]) -> bool:
	return len(questions) == len(responses)

func next_question(questions: Array[Question], responses: Array[Response]) -> Question:
	return questions[len(responses)]

func answer(responses: Array[Response], newAnswer: bool) -> Array[Response]:
	var r = Response.new()
	r.answer = newAnswer
	responses.append( r )
	return responses

func load_questions(path: String) -> Array[Question]:
	if not FileAccess.file_exists(path):
		push_error("file does not exist")
		return []
	
	var file = FileAccess.open( path, FileAccess.READ )
	var jsonResult = JSON.parse_string(file.get_as_text())
	
	if typeof(jsonResult) != TYPE_ARRAY:
		push_error("file does not contain a list")
		return []
	
	var result: Array[Question] = []
	for q in jsonResult:
		var question = Question.new()
		question.quote = q["quote"]
		question.author = q["author"]
		question.answer = q["answer"]
		question.url = q["url"]
		result.append(question)

	return result
