package fetchers

import (
	"backend/app/models"
	"context"
	"github.com/carlmjohnson/requests"
	"go.uber.org/zap"
)

func GetDeveloperInfo(githubId string, githubToken *string) (*DeveloperFull, error) {
	var data *DeveloperFull
	err := requests.
		URL("https://api.github.com/users/"+githubId).
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		zap.L().Error("Error fetching developer info", zap.Error(err))
		return nil, err
	}
	data.AllRepos, err = GetDeveloperPublicRepos(githubId, data.PublicRepos, githubToken)
	if err != nil {
		zap.L().Error("GetDeveloperPublicRepos failed", zap.Error(err))
		zap.L().Debug("GetDeveloperPublicRepos failed", zap.Error(err),
			zap.String("githubId", githubId))
		allRepos := make([]ReposFull, 0)
		data.AllRepos = allRepos
		return data, err
	}
	return data, nil
}

func GetDeveloperPublicRepos(githubId string, lens int, githubToken *string) ([]ReposFull, error) {
	data := make([]ReposFull, 0, lens)
	err := requests.URL("https://api.github.com/users/"+githubId+"/repos").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetReposDetail(reposFullName string, githubToken *string) (*ReposDetailsFull, error) {
	data := new(ReposDetailsFull)
	err := requests.URL("https://api.github.com/repos/"+reposFullName).
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposLanguages(reposFullName string, githubToken *string) (map[string]int64, error) {
	var data map[string]int64
	err := requests.URL("https://api.github.com/repos/"+reposFullName+"/languages").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetReposContributors(reposFullName string, githubToken *string) (*[]models.MiniDeveloper, error) {
	var data []models.MiniDeveloper
	err := requests.URL("https://api.github.com/repos/"+reposFullName+"/contributors").
		Header("Authorization", "Bearer "+*githubToken).
		// Header("X-GitHub-Api-Version: 2022-11-28").
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return &data, err
	}
	return &data, nil
}

func GetDeveloperComments(githubLogin string, githubToken *string) (*[]string, error) {
	var data CommentItems
	err := requests.URL("https://api.github.com/search/issues?q=commenter:"+githubLogin).
		Header("Authorization", "Bearer "+*githubToken).
		ToJSON(&data).
		Fetch(context.Background())
	if err != nil {
		return nil, err
	}
	if data.Comments == nil {
		return nil, nil
	}

	devComments := make([]string, 0, MaxCommentLines)
	for _, comment := range data.Comments {
		var lines []CommentLines
		err = requests.URL(comment.TimelineUrl).
			Header("Authorization", "Bearer "+*githubToken).
			ToJSON(&lines).
			Fetch(context.Background())
		if err != nil {
			zap.L().Debug("GetDeveloperComments failed", zap.Error(err),
				zap.String("url", comment.TimelineUrl))
			continue
		}
		for _, line := range lines {
			if line.Event == "commented" && line.User.Login == githubLogin {
				str := line.Body[:min(MaxLineLength, len(line.Body))]
				devComments = append(devComments, str)
			}
			if len(devComments) >= MaxCommentLines {
				break
			}
		}
		if len(devComments) >= MaxCommentLines {
			break
		}
	}
	return &devComments, nil
}

/*
接下来我会给出一个用户的github历史评论记录，请你通过评论使用的语言，评论的语法正确性猜测该用户的在住地区，并将结果以json格式{nation:结果}返回。历史评论记录以json格式给出。
[
	"编码格式UTF-8",
	"在admin表中插入数据即可，没有对密码进行加密",
	"连数据需要现安装mysql 建议5.6/5.7版本，创建数据库，执行webContent/db/下的sql文件。具体度娘有很多",
	"在admin表中插入数据即可，没有进行密码加密",
	"qq 296943773 和我联系吧",
	"在webconten/db目录下，觉得好记得打个star\r\n\r\n\r\n\r\n\r\n------------------\u0026nbsp;原始邮件\u0026nbsp;------------------\r\n发件人:                                                                                                                        \"y"
]
*/
