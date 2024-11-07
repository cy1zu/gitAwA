package guessers

import (
	"backend/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/carlmjohnson/requests"
	"go.uber.org/zap"
	"strconv"
)

func Init() (NationGuesserRequest, error) {
	var data NationGuesserRequest
	err := requests.URL("/v2/app/conversation").Host("qianfan.baidubce.com").
		ContentType("application/json").
		Header("X-Appbuilder-Authorization",
			fmt.Sprintf("Bearer %s", config.Conf.LLMAccessToken)).
		BodyJSON(map[string]string{
			"app_id": config.Conf.LLMAppId,
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
			fmt.Sprintf("Bearer %s", config.Conf.LLMAccessToken)).
		BodyJSON(map[string]interface{}{
			"app_id":          config.Conf.LLMAppId,
			"query":           queryStr,
			"conversation_id": guesserRequest.ConversationId,
			"stream":          false,
		}).ToJSON(&data).Fetch(context.Background())
	if err != nil {
		zap.L().Error("call nationGuesser failed", zap.Error(err))
		return NationGuesserResult{}, err
	}

	resultStr := data.Answer[8 : len(data.Answer)-4]
	println(resultStr)
	//normally return
	var result NationGuesserResult
	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		zap.L().Error("call nationGuesser failed", zap.Error(err))
		return NationGuesserResult{
			Nation: "N/A",
			Value:  0.00,
		}, err
	}
	return result, nil
}

/*

 */
