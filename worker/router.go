package worker

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func RouterGetData(c *gin.Context) {
	switch c.Query("req") {
	case "music":
		RouterMusic(c)
	case "level-noTimetag":
		RouterLevel(c)
	case "timetag":
		RouterNote(c)
	case "score":
		RouterScore(c)
	case "comment":
		RouterComment(c)
	case "username":
		RouterUser(c)
	default:

	}
}

func RouterPostData(c *gin.Context) {
	switch c.Query("req") {
	case "userName-add":
		RouterUserAdd(c)
	case "score-add":
		RouterScoreAdd(c)
	case "playcount-add":
		RouterPlayCountAdd(c)
	case "comment-add":
		RouterCommentAdd(c)
	default:

	}
}

/**
Get Data
*/
func RouterMusic(c *gin.Context) {
	//if t, err := strconv.ParseInt(c.Query("time"), 10, 64); err != nil {
	//	print(t)
	//} else if t == 0 {
	//	MusicDataGetInitMusic(c)
	//
	//} else {
	//	MusicDataGetLastUpdateMusic(c, t)
	//}
	if t, err := strconv.ParseInt(c.Query("time"), 10, 64); err != nil {
		print(t)
	} else {
		MusicDataGetInitMusic(c)

	}
}

func RouterLevel(c *gin.Context) {
	if t, err := strconv.ParseInt(c.Query("time"), 10, 64); err != nil {
		print(t)
	} else if t == 0 {
		LevelDataGetInitLevel(c)
	}
}

func RouterNote(c *gin.Context) {
	if id, err := strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
		print(id)
	} else if id != 0 {
		LevelDataGetNotes(c, id)
	}
}

func RouterScore(c *gin.Context) {
	var l = c.Query("levelID")
	if t, err := strconv.ParseInt(c.Query("time"), 10, 64); err != nil {
		print(l)
		print(t)
	} else if l == "ALL" {

	} else {

	}

}

func RouterComment(c *gin.Context) {
	var l = c.Query("levelID")
	if t, err := strconv.ParseInt(c.Query("time"), 10, 64); err != nil {
		print(l)
		print(t)
	} else if l == "ALL" {

	} else {

	}
}

func RouterUser(c *gin.Context) {
	t, err := strconv.ParseInt(c.Query("time"), 10, 64)
	if err == nil {
		print(t)
	}
}

/**
Post Data
*/

func RouterUserAdd(c *gin.Context) {

}

func RouterScoreAdd(c *gin.Context) {

}

func RouterPlayCountAdd(c *gin.Context) {

}

func RouterCommentAdd(c *gin.Context) {

}
