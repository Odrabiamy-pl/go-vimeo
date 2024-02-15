package test

import (
	"context"
	"github.com/Odrabiamy-pl/go-vimeo/pkg/vimeoapi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetVideo(t *testing.T) {
	t.Run("get videos", func(t *testing.T) {
		client := provideClient()

		ctx := authorize(context.Background())
		videos, httpResp, err := client.VideosEssentialsAPI.
			GetVideos(ctx, 202422358).
			Page(1).
			Execute()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		assert.NotNil(t, videos)
	})

	t.Run("get video", func(t *testing.T) {
		client := provideClient()

		ctx := authorize(context.Background())
		video, httpResp, err := client.VideosEssentialsAPI.
			GetVideo(ctx, 899476146). //910433176
			Execute()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		assert.NotNil(t, video)
	})
}

func authorize(ctx context.Context) context.Context {
	//ctx = context.WithValue(ctx, vimeoapi.ContextOAuth2, oauth2.StaticTokenSource(oauth2.Token{
	//	AccessToken:  "",
	//	TokenType:    "",
	//	RefreshToken: "",
	//	Expiry:       time.Time{},
	//}))
	return context.WithValue(ctx, vimeoapi.ContextAccessToken, "")
}

func provideClient() *vimeoapi.APIClient {
	apiCfg := vimeoapi.NewConfiguration()
	apiCfg.UserAgent = "Learnest/0.3.0"
	apiCfg.Debug = true
	return vimeoapi.NewAPIClient(apiCfg)
}
