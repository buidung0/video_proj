package service

import "gitlab.com/pragmaticrevies/goland-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	video []entity.Video
}

func New() *videoService {
	return &videoService{}
}

func (service videoService) Save(video entity.Video) entity.Video {
	service.video = append(service.video, video)
	return video
}

func (service videoService) FindAll() []entity.Video {
	return nil
}
