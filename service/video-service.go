package service

import (
	"PR_gin_g/entity"
	"PR_gin_g/repository"
)

type VideoService interface {
	Save(video entity.Video) error
	Update(video entity.Video) error
	Delete(video entity.Video) error
	Detail(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	VideoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		VideoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) error {
	service.VideoRepository.Save(video)
	return nil
}

func (service *videoService) Update(video entity.Video) error {
	service.VideoRepository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.VideoRepository.Delete(video)
	return nil
}

func (service *videoService) Detail(video entity.Video) entity.Video {
	videoDetail := service.VideoRepository.Detail(video)
	return videoDetail

}

func (service *videoService) FindAll() []entity.Video {
	return service.VideoRepository.FindAll()
}
