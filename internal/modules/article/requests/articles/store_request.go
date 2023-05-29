package articles

type StoreRequest struct {
	Title   string `form:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" binding:"required,min=8,max=60000"`
}
