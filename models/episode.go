package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type Episode struct {
	Id            int `json:"id" orm:"pk"`
	MovieId       string
	EpisodeNumber string
	EpisodeName   string
	Download      string
	EpisodeUrl    string
	EpisodeImg    string
	EpisodeInfo   string
	DeEpisodeInfo string
}

//type M3u8s []*M3u8
type M3u8 struct {
	Format      string `json:"format"`
	AudioLang   string `json:"audio_lang"`
	HardsubLang string `json:"hardsub_lang"`
	Url         string `json:"url"`
}

func GetEpisodeList(movieId string) []Episode {
	o := orm.NewOrm()
	var list []Episode
	_, _ = o.Raw("select * from  episode where movie_id = ?", movieId).QueryRows(&list)
	return list
}

func GetEpisodeUrl(episodeId string) (string, string) {
	o := orm.NewOrm()
	var episode Episode
	var list []*M3u8
	o.Raw("select episode_url,movie_id from episode where id = ?", episodeId).QueryRow(&episode)
	if err := json.Unmarshal([]byte(episode.EpisodeUrl), &list); err == nil {

	} else {

	}
	if list == nil {
		return "", episode.MovieId
	}
	var playerPath string
	for _, x := range list {
		if x.HardsubLang == "esES" {
			playerPath = x.Url
			break
		}

	}
	if playerPath == "" {
		playerPath = list[0].Url
	}
	return playerPath, episode.MovieId
}
