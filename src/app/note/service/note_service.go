package service

import "github.com/gin-gonic/gin"

// NoteService /**
type NoteService interface {
	Add(r *gin.Context)
	Edit(r *gin.Context)
	MarkFinish(r *gin.Context)
	Get(r *gin.Context)
}

type Note struct{}

func (Note) Add(r *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (Note) Edit(r *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (Note) MarkFinish(r *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (Note) Get(r *gin.Context) {
	//TODO implement me
	panic("implement me")
}
