package main_test

import (
	"github.com/golang/mock/gomock"
	"testing"
	"workshop/apis/02httclient"
	mock_main "workshop/apis/02httclient/mocks"
)

func TestHttpGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_main.NewMockHttpRequester(ctrl)

	m.EXPECT().HttpGet("/get").Return(200)

	testService := main.HttpService{"test_url"}

	testService.GetUsers()



}
