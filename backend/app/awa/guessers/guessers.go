package guessers

import (
	"context"
	"encoding/json"
	"github.com/carlmjohnson/requests"
	"go.uber.org/zap"
	"strconv"
)

func Init() (NationGuesserRequest, error) {
	var data NationGuesserRequest
	err := requests.URL("/v2/app/conversation").Host("qianfan.baidubce.com").
		ContentType("application/json").
		Header("X-Appbuilder-Authorization",
			"Bearer bce-v3/ALTAK-6fJaqaqDvFOnYvZXBgkf5/c8fef321ddba046c86641c6c40d98ddc51133e65",
		).BodyJSON(map[string]string{
		"app_id": "2bb9f52e-b7db-42f6-82e3-ff7e6d9691c5",
	}).ToJSON(&data).Fetch(context.Background())
	if err != nil {
		zap.L().Error("call nationGuesser failed", zap.Error(err))
		return NationGuesserRequest{}, err
	}
	return data, nil
}

func GuessNation(guesserRequest NationGuesserRequest, query map[string]interface{}) (NationGuesserResult, error) {
	var data NationGuesserResponse
	queryBytes, _ := json.Marshal(query)
	queryStr := strconv.Quote(string(queryBytes))
	err := requests.URL("/v2/app/conversation/runs").Host("qianfan.baidubce.com").
		ContentType("application/json").
		Header("X-Appbuilder-Authorization",
			"Bearer bce-v3/ALTAK-6fJaqaqDvFOnYvZXBgkf5/c8fef321ddba046c86641c6c40d98ddc51133e65").
		BodyJSON(map[string]interface{}{
			"app_id":          "2bb9f52e-b7db-42f6-82e3-ff7e6d9691c5",
			"query":           queryStr,
			"conversation_id": guesserRequest.ConversationId,
			"stream":          false,
		}).ToJSON(&data).Fetch(context.Background())
	if err != nil {
		zap.L().Error("call nationGuesser failed", zap.Error(err))
		return NationGuesserResult{}, err
	}

	// sample =  "```json\n{\n  \"nation\": \"匈牙利\",\n  \"value\": 1.00\n}\n```"
	resultStr := data.Answer[8 : len(data.Answer)-4]

	//normally return
	var result NationGuesserResult
	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		zap.L().Error("call nationGuesser failed", zap.Error(err))
		return NationGuesserResult{}, err
	}
	return result, nil
}

/*

 */
