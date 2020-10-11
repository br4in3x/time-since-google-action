package response

type Simple struct {
	Session Session `json:"session"`
	Prompt  Prompt  `json:"prompt"`
	Scene   Scene   `json:"scene"`
}

type Params struct {
}

type Session struct {
	ID     string `json:"id"`
	Params Params `json:"params"`
}

type FirstSimple struct {
	Speech string `json:"speech"`
	Text   string `json:"text"`
}

type Prompt struct {
	Override    bool        `json:"override"`
	FirstSimple FirstSimple `json:"firstSimple"`
}

type Slots struct {
}

type Next struct {
	Name string `json:"name"`
}

type Scene struct {
	Name  string `json:"name"`
	Slots Slots  `json:"slots"`
	Next  Next   `json:"next"`
}

func SimpleResponse(text string) *Simple {
	return &Simple{
		Session: Session{
			ID:     "",
			Params: Params{},
		},
		Prompt: Prompt{
			Override: false,
			FirstSimple: FirstSimple{
				Speech: text,
				Text:   text,
			},
		},
	}
}
