package domain

type SearchResult struct {
	Items []Item
}

type Item struct {
	Title string
	Link  string
}
