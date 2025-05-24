package shareds

type Value struct {
	Value string `json:"value"`
}

type IntValue struct {
	Value int `json:"value"`
}

type FloatValue struct {
	Value float32 `json:"value"`
}
type ValueColor struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type IntValueColor struct {
	Value int    `json:"value"`
	Color string `json:"color"`
}
