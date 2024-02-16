package test

import (
	"context"
	"github.com/Odrabiamy-pl/go-vimeo/pkg/mocks"
	"github.com/Odrabiamy-pl/go-vimeo/pkg/vimeoapi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetVideo(t *testing.T) {
	t.Run("get videos", func(t *testing.T) {
		t.Skip("videos list needs parent object in documentation")
		srv := mocks.Run()
		defer srv.Close()

		client := provideClient(srv.URL)
		ctx := authorize(context.Background())
		videos, httpResp, err := client.VideosEssentialsAPI.
			GetVideos(ctx, 202422358).
			Execute()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		assert.NotNil(t, videos)
	})

	t.Run("get video", func(t *testing.T) {
		srv := mocks.Run()
		defer srv.Close()

		client := provideClient(srv.URL)
		ctx := authorize(context.Background())
		video, httpResp, err := client.VideosEssentialsAPI.
			GetVideo(ctx, 899476146).
			Execute()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpResp.StatusCode)
		assert.NotNil(t, video)
	})
}

func authorize(ctx context.Context) context.Context {
	return context.WithValue(ctx, vimeoapi.ContextAccessToken, os.Getenv("ACCESS_TOKEN"))
}

func provideClient(url string) *vimeoapi.APIClient {
	apiCfg := vimeoapi.NewConfiguration()
	apiCfg.UserAgent = "Learnest/0.3.0"
	apiCfg.Debug = true
	apiCfg.Servers = vimeoapi.ServerConfigurations{
		{URL: url},
	}
	return vimeoapi.NewAPIClient(apiCfg)
}
