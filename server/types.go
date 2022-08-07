package server

type requestJson struct {
	Message       string `json:"message"`
	To            string `json:"to"`
	From          string `json:"from"`
	TimeToLifeSec int    `json:"timeToLifeSec"`
}

type responseJson struct {
	Message string `json:"message"`
}
