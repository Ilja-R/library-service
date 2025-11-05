package http


func (s *Server) endpoints() {
	s.router.GET("/ping", s.Ping)

	

	apiG := s.router.Group("/api", s.checkUserAuthentication)
	{
		bookG:=apiG.Group("/books")
		{
			bookG.GET("/", s.GetAllBooks)
			bookG.GET("/:id", s.GetBookByID)
			bookG.POST("/", s.checkIsAdmin, s.CreateBook)
			bookG.PUT("/:id", s.checkIsAdmin, s.UpdateBook)
			bookG.DELETE("/:id", s.checkIsAdmin, s.DeleteBookByID)
		}
		
	}

}