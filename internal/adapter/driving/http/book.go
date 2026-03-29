package http

import (
	"fmt"
	"log"
	"strings"
	"time"

	"net/http"
	"strconv"

	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/errs"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Pub_date    string `json:"pub_date"`
	Publisher   string `json:"publisher"`
	Genre       string `json:"genre"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func (b *Book) FromDomain(dBook domain.Book) {
	b.ID = dBook.ID
	b.Title = dBook.Title
	b.Pub_date=dBook.Pub_date.Format("2 Jan 2006")
	b.Publisher = dBook.Publisher
	b.Genre = dBook.Genre
	b.Pages = dBook.Pages
	b.Description = dBook.Description
	b.Created_at = dBook.Created_at.Format("2 Jan 2006 15:04:05")
	b.Updated_at = dBook.Updated_at.Format("2 Jan 2006 15:04:05")
}


type CreateBook struct {
	Title         string `json:"title"`
	AuthorName    string `json:"name"`
	AuthorSurname string `json:"surname"`
	PubDate       string `json:"pub_date"`
	Publisher     string `json:"publisher"`
	Genre         string `json:"genre"`
	Pages         int    `json:"pages"`
	Description   string `json:"description"`
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (s *Server) GetAllBooks(c *gin.Context) {
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid userID in context"})
		return
	}

	dBooks, err := s.uc.BookGetter.GetAllBooks(c)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var (
		books []Book
		book  Book
	)
	for _, dBook := range dBooks {
		book.FromDomain(dBook)
		book.Pub_date = dBook.Pub_date.Format("2 Jan 2006")
		book.Created_at = dBook.Created_at.Format("2 Jan 2006 15:04:05")
		book.Updated_at = dBook.Updated_at.Format("2 Jan 2006 15:04:05")
		books = append(books, book)

	}

	c.JSON(http.StatusOK, books)
}

func (s *Server) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidBookID)
		return
	}
	dBook, err := s.uc.BookGetter.GetBookByID(c, id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var book Book
	book.FromDomain(dBook)

	c.JSON(http.StatusOK, book)
}

func (b *CreateBook) ToDomain() *domain.CreateBook {
	strTime:=strings.Split(b.PubDate,"-")
	
	
	yearInt,err:=strconv.Atoi(strTime[0])
	if err!=nil{
		log.Fatal(err)
	}
	monthInt,err:=strconv.Atoi(strTime[1])
	if err!=nil{
		log.Fatal(err)
	}
	dayInt,err:=strconv.Atoi(strTime[2])
	if err!=nil{
		log.Fatal(err)
	}

	t:=time.Date(yearInt,time.Month(monthInt),dayInt,0,0,0,0,time.Now().Location())
	fmt.Printf("t: %v\n", t)
	return &domain.CreateBook{
		Title:         b.Title,
		AuthorName:    b.AuthorName,
		AuthorSurname: b.AuthorSurname,
		Pub_date: t,
		Publisher:     b.Publisher,
		Genre:         b.Genre,
		Pages:         b.Pages,
		Description:   b.Description,
	}
}

func (s *Server) CreateBook(c *gin.Context) {
	createBook := CreateBook{}
	err := c.BindJSON(&createBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{
			Error: fmt.Errorf("%s", "error binding").Error(),
		})
		return
	}
	dBook := createBook.ToDomain()
	log.Println("dBook",dBook)
	err = s.uc.BookCreator.CreateBook(c, *dBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{
			Error: err.Error(),
		})
		return

	}
	if err == nil {
		c.JSON(http.StatusInternalServerError, CommonResponse{
			Message: "book successfully created",
		},
		)
		return
	}
}

type UpdateBookBody struct {
	Title         string `json:"title"`
	AuthorName    string `json:"name"`
	AuthorSurname string `json:"surname"`
	Publisher     string `json:"publisher"`
	Genre         string `json:"genre"`
	Pages         int    `json:"pages"`
	Description   string `json:"description"`
}

func (u *UpdateBookBody) ToDomain() *domain.UpdateBookBody {
	return &domain.UpdateBookBody{
		Title:         u.Title,
		AuthorName:    u.AuthorName,
		AuthorSurname: u.AuthorSurname,
		Publisher:     u.Publisher,
		Genre:         u.Genre,
		Pages:         u.Pages,
		Description:   u.Description,
	}
}

func (s *Server) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidBookID)
		return
	}
	var updBookBody UpdateBookBody
	err = c.BindJSON(&updBookBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommonError{
			Error: fmt.Errorf("error binding").Error(),
		})
		return
	}
	dBookBody := updBookBody.ToDomain()
	if err = s.uc.BookUpdater.UpdateBook(c, *dBookBody, id); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, CommonError{
			Error: errs.ErrSomethingWentWrong.Error(),
		})
		return
	}
	if err == nil {
		c.JSON(http.StatusOK, CommonResponse{
			Message: "book updated",
		})
		return
	}

}

func (s *Server) DeleteBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidBookID)
		return
	}
	err = s.uc.BookDeleter.DeleteBookByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{
			Error: err.Error(),
		})
		return
	}
	if err == nil {
		c.JSON(http.StatusOK, CommonResponse{
			Message: "book is successfully deleted",
		})
		return
	}
}

func (s *Server) SearchByTitle(c *gin.Context) {
	title := c.Query("title")
	book:=Book{}
	dbooks, err := s.uc.BookSearcher.SearchByTitle(c, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{
			Error: err.Error(),
		})
		return
	}
	books:=make([]Book,0)
	for _,v:= range dbooks{
		book.FromDomain(v)
		books=append(books,book)
	}

	c.JSON(http.StatusOK,books)

}


func(s*Server) OrderBookByTitle(c*gin.Context){
	title := c.Query("title")
	username,_:=c.Get(UsernameCtx)
	dbooks,err:=s.uc.BookSearcher.SearchByTitle(c, title)
	if err!=nil{
		c.JSON(http.StatusBadRequest,CommonError{
			Error:err.Error(),
		})
		return
	}
	if dbooks==nil{
		c.JSON(http.StatusBadRequest,CommonError{
			Error: errs.ErrNoBookWithFollowingTitle.Error(),
		})
		return 
	}
	book:=Book{}
	books:=make([]Book,0)
	for _,v:= range dbooks{
		book.FromDomain(v)
		books=append(books,book)
	}
	err=s.uc.BookOrderer.OrderBookByTitle(c,title,username.(string))
	if err!=nil{
		c.JSON(http.StatusBadRequest,CommonError{
			Error: err.Error(),
		})
		return 
	}
}

func (s*Server)GetMyBooks(c*gin.Context){
	username ,_:=c.Get(UsernameCtx)
	
	dBooks,err:=s.uc.BookGetter.GetMyBooks(c,username.(string))
	if err!=nil{
		c.JSON(http.StatusBadRequest,CommonError{
			Error: err.Error(),
		})
		return
	}
	books:=make([]Book,0)
	book:=Book{}

	for _,dBook:=range dBooks{
		book.FromDomain(dBook)
		books=append(books,book)
	}

	c.JSON(http.StatusOK,books)
}