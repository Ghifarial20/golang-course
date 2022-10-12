package datastruct

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

// type ResponseShow struct {
// 	Total  int
// 	Detail []Album
// }

type ResponseData struct {
	Status      int
	Description Album
}
