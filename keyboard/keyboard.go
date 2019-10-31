package keyboard

var buttonsRow []button

type color string

const (
	Positive  color = "positive"
	Primary         = "primary"
	Secondary       = "secondary"
	Negative        = "negative"
)

type payload struct {
	Value uint8 `json:"button"`
}

type action struct {
	Type    string  `json:"type"`
	Payload payload `json:"payload"`
	Label   string  `json:"label"`
}

type button struct {
	Action action `json:"action"`
	Color  color  `json:"color"`
}

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]button `json:"buttons"`
}

func (k *Keyboard) AddTextButton(buttonText string, color color) {
	b := button{
		Action: action{
			Type:    "text",
			Payload: payload{gen()},
			Label:   buttonText,
		},
		Color: color,
	}

	buttonsRow = append(buttonsRow, b)
}

func (k *Keyboard) FormLine() {
	k.Buttons = append(k.Buttons, buttonsRow)
	buttonsRow = buttonsRow[:0]
}

func Empty() Keyboard {
	return Keyboard{
		OneTime: true,
	}
}

func New(oneTime bool) *Keyboard {
	gen = newPayloadValue()
	return &Keyboard{
		OneTime: oneTime,
	}
}

var gen func() uint8

func newPayloadValue() func() uint8 {
	var n uint8 = 0
	return func() uint8 {
		n += 1
		return n
	}
}
