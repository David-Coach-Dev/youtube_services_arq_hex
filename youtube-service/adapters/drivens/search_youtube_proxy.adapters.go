package adapters

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	models "youtube_service_arq_hex/youtube-service/models"

)

type youtubeApi struct {
	apiKey     string
	channelIDs []string
	service    *youtube.Service
}

type YoutubeSearchAdapter struct {
	gc *gin.Context
}

// The connectToYoutube method is used in the NewYoutubeSearchAdapter function.
func (ysa *youtubeApi) connectToYoutube() (*youtubeApi, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("YOUTUBE_API_KEY environment variable is not set")
	}

	channelIDs := os.Getenv("YOUTUBE_CHANNEL_IDS")
	if channelIDs == "" {
		return nil, errors.New("YOUTUBE_CHANNEL_IDS environment variable is not set")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	ysa.apiKey = apiKey
	ysa.service = service
	ysa.channelIDs = strings.Split(channelIDs, ",")

	return ysa, nil
}

func (ysa *YoutubeSearchAdapter) GetSearch(gc *gin.Context) (*[]models.DataYT, error) {

	searchQuery := gc.Query("search")
	if searchQuery == "" {
		return nil, errors.New("the search param is required")
	}

	maxResultsStr := gc.Query("maxResult")
	var err error
	var maxResultsParams int64
	if maxResultsStr == "" {
		maxResultsParams=10
	}else{
	maxResultsParams, err = strconv.ParseInt(maxResultsStr, 10, 64)
	if err != nil {
		return nil, err
	}
	}
  /*
		channel
		playlist
		video
	*/
	typeSearch := gc.Query("type")
	if typeSearch == "" {
		typeSearch = "video"
	}

	pageTokenSearch := gc.Query("pageToken")

	ytApi := youtubeApi{}

	ytc, err := ytApi.connectToYoutube()
	if err != nil {
		return nil, err
	}

	var resultSearch []models.DataYT
	for _, channelID := range ytc.channelIDs {
		call := ytc.service.Search.List([]string{"id", "snippet"}).
			Q(searchQuery).
			ChannelId(channelID).
			MaxResults(maxResultsParams).
			Type(typeSearch).
			PageToken(pageTokenSearch)
		response, err := call.Do()
		if err != nil {
			return nil, err
		}

		if len(response.Items) == 0 {
			continue
		}


		idTemp := ""
		urlTemp := ""
		resultItems := []models.Items{}
		for _, item := range response.Items {
			if typeSearch == "playlist" {
				idTemp = item.Id.PlaylistId
				urlTemp = "https://www.youtube.com/playlist?list=" + idTemp
			}else{
				idTemp = item.Id.VideoId
				urlTemp = "https://www.youtube.com/watch?v=" + idTemp
			}
			resultItems = append(resultItems, models.Items{
				ID:           idTemp,
				Title:        item.Snippet.Title,
				Description:  item.Snippet.Description,
				Channel:      item.Snippet.ChannelTitle,
				Live:         item.Snippet.LiveBroadcastContent,
				URL:          urlTemp,
				Thumbnails:   item.Snippet.Thumbnails.Default.Url,
			})
		}

		PageInfo := []models.PageInfo{}

		PageInfo = append(PageInfo, models.PageInfo{
			TotalResults:     int(response.PageInfo.TotalResults),
			ResultsPerPage:   int(response.PageInfo.ResultsPerPage),
		})

		resultSearch = append(resultSearch, models.DataYT{
				Kind: "youtube#searchListResponse",
				Etag: response.Etag,
				NextPageToken: response.NextPageToken,
				PrevPageToken: response.PrevPageToken,
				PageInfo: PageInfo,
				Items: resultItems,
			})
	}

	return &resultSearch, nil
}

func NewYoutubeSearchAdapter(gc *gin.Context) *YoutubeSearchAdapter {
	return &YoutubeSearchAdapter{
		gc: gc,
	}
}
