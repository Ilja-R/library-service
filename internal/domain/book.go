package domain

import "time"


type Book struct{
	ID int 
	Title string 
	Pub_date time.Time
	Publisher string 
	Genre string 
	Pages int 
	Description string 
	Created_at time.Time
	Updated_at time.Time
}

type CreateBook struct{
	Title string 
	AuthorName string 
	AuthorSurname string 
	Pub_date time.Time 
	Publisher string
	Genre string 
	Pages int 
	Description string 
	Created_at time.Time 
	Updated_at time.Time 
}

type UpdateBookBody struct{
	Title string 
	AuthorName string 
	AuthorSurname string 
	Publisher string
	Genre string 
	Pages int 
	Description string 
}