extends Node2D

var currentQuestion: int = 0
var questions: Array[QuestionEngine.Question] = []
var responses: Array[QuestionEngine.Response] = []

func _ready() -> void:
	questions = QuestionEngine.load_questions("res://questions/test_questions1.json")
	if questions.is_empty():
		push_error("panic!")
		return
	
	$PanelContainer/TrueButton.pressed.connect( self.true_button_pressed )
	$PanelContainer/FalseButton.pressed.connect( self.false_button_pressed )
	show_question()
	
func true_button_pressed():
	responses = QuestionEngine.answer( responses, true )
	next_round()
	
func false_button_pressed():
	responses = QuestionEngine.answer( responses, false )
	next_round()

func next_round():
	if not QuestionEngine.is_finished(questions, responses):
		self.currentQuestion += 1
		show_question()
	else:
		show_result()

func show_result():
	$PanelContainer/TrueButton.visible = false
	$PanelContainer/FalseButton.visible = false	
	
	var correct_answers: int = 0
	for i in range(0, len(responses)):
		if questions[i].answer == responses[i].answer:
			correct_answers += 1
	
	$PanelContainer/QuoteLabel.text = "finished with " + str(correct_answers) + " correct answers out of " + str(len(responses))

func show_question():
	$PanelContainer/QuoteLabel.text = questions[self.currentQuestion].quote + "\n- " + questions[self.currentQuestion].author
