package entity

type Url struct {
	Id    int `db:"id"`
	Url   int `db:"url"`
	Alias int `db:"alias"`
}
