package domain

type SearchResp struct {
	Items []Item
}

type Item struct {
	Title string
	Link  string
}
