package catalogservice

type Item struct {
	Url        string `db:"Url"`
	Count      int    `db:"Count"`
	CategoryID int    `db:"CategoryID"`
}
