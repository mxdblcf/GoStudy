package main

type requestBody struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userid"`
}

type responseBody struct {
	Code int      `json:"code"`
	Test string   `json:"test"`
	List []string `json:"list"`
	Url  string   `json:"url"`
}

func process(inputChan <-chan string, userid string) {
	for {
		input := <-inputChan
		if input == "EOF" {
			break

		}
	}
	//构建结构体
	//&responseBody{
	//	Key
	//}
}
func main() {

}
