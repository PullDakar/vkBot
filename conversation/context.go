package conversation

type State int

const (
	DialogNotStarted State = iota
	DialogStarted          // отправил кнопку "Начать"

	IdeaDescribed // описал суть идеи

	ChooseDreamerType // выбрал тип участника

	ChooseAnalystType // выбрал аналитика

	ChooseDevType       //  выбрал разработчика
	ChooseBackDevType   //  выбрал back-end разработчика
	ChooseFrontDevType  // выбрал front-end разработчика
	ChooseMobileDevType // выбрал мобильного разработчика

	ChooseTesterType // выбрал тестировщика

	ChooseDreamerCount // выбрал количество участников
)

type DialogState interface {
	Reply(ctx *DialogStateContext, message string, pathToKeyboard string)
}

type DialogStateContext struct {
	CurrentState DialogState
}

func (ctx *DialogStateContext) execute(message, pathToKeyboard string) {
	ctx.CurrentState.Reply(ctx, message, pathToKeyboard)
}
