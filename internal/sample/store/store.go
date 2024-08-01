package store

type Book struct {
	Id      int64    `json:"id"`      // 图书ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

type BookStore interface {
	Save(book *Book)
	GetById(id int64) *Book
}
