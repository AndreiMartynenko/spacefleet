package main

type SpaceCraft struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Crew     int        `json:"crew"`
	Image    string     `json:"image"`
	Value    float64    `json:"value"`
	Status   string     `json:"status"`
	Armament []Armament `json:"armament"`
}

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type Response struct {
	Payload interface{} `json:"payload"`
	Error   *Error      `json:"error"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
