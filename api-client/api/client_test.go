package api_test

import (
	"api-client/api"
	mock_api "api-client/api/mock-httpclient"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestClient(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHttpClient := mock_api.NewMockHttpCilent(ctrl)

	testBaseUrl := "https://www.bilibili.com"
	exceptedRequest, err := http.NewRequest(http.MethodGet, testBaseUrl+"/category/monsters?game=totk", nil)

	require.NoError(t, err)

	exceptionResponse := &api.GetMonsterResponse{
		Data: []api.Monster{
			{
				Category:        "monsters",
				CommonLocations: []string{"Gerudo Highlands", "Gerudo Desert"},
				Description:     "These particularly clever monsters bury themselves in deep sand or snow and disguise themselves as treasure chests. Anyone who approaches the chests is attacked. The treasures chests are not magnetic, which proves that they are actually a part of these monsters' bodies.",
				Dlc:             false,
				Drops: []string{"octorok tentacle",
					"octo balloon",
					"octorok eyeball",
					"green rupee",
					"blue rupee",
					"red rupee",
					"purple rupee",
					"silver rupee"},
				ID:    96,
				Image: "https://botw-compendium.herokuapp.com/api/v3/compendium/entry/treasure_octorok/image",
				Name:  "treasure octorok",
			},
		},
		Message: "",
		Status:  200,
	}

	jsonBytes, err := json.Marshal(exceptionResponse)
	require.NoError(t, err)

	mockHttpClient.EXPECT().Do(exceptedRequest).Return(&http.Response{
		Body: io.NopCloser(bytes.NewBuffer(jsonBytes)),
	}, nil)

	client := api.NewClient(testBaseUrl, mockHttpClient)
	response, err := client.GetMonster()
	require.NoError(t, err)
	require.Equal(t, exceptionResponse, response)
}
