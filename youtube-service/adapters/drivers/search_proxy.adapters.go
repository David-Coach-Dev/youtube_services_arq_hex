package adapters

import (
	"github.com/gin-gonic/gin"

	youtube_service "youtube_service_arq_hex/youtube-service"
	models "youtube_service_arq_hex/youtube-service/models"
)

type SearchYouTubeAdapter struct {
	gc             *gin.Context
	youtubeService *youtube_service.YouTubeService
}

func (sya *SearchYouTubeAdapter) GetSearch(gc *gin.Context) (*[]models.DataYT, error) {
	return sya.youtubeService.GetSearch(gc)
}

func NewSearchYouTubeAdapter(gc *gin.Context, youtubeService *youtube_service.YouTubeService) *SearchYouTubeAdapter {
	return &SearchYouTubeAdapter{
		gc:             gc,
		youtubeService: youtubeService,
	}
}
