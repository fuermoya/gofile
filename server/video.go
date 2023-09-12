package server

import (
	response "github.com/fuermoya/gofile/common"
	"github.com/fuermoya/gofile/utils"
	astisub "github.com/fuermoya/gofile/utils/subtitles"
	"github.com/gin-gonic/gin"
	"os"
)

var exts = [...]string{".ass", ".srt", ".ssa", ".stl", ".ttml"}

const vtt = ".vtt"

// LookVideo TODO  不支持的视频需要转码为mp4
func (s *Server) LookVideo(c *gin.Context) {
	pathAny, _ := c.Get("path")
	filePath := pathAny.(string)
	c.File(filePath)
}

// GetSubtitle 获取字幕，与视频同名
func (s *Server) GetSubtitle(c *gin.Context) {
	pathAny, _ := c.Get("path")
	filePath := pathAny.(string)

	extSubtitles := utils.JoinExt(filePath, vtt)
	_, err := os.Stat(extSubtitles)
	if !(err == nil || os.IsExist(err)) {
		//没有字幕，尝试转换为vtt
		for _, v := range exts {
			ext := utils.JoinExt(filePath, v)
			_, err = os.Stat(ext)
			if err == nil || os.IsExist(err) {
				err = convertSubtitles(ext, extSubtitles)
				if err != nil {
					continue
				}
				response.OkWithData(gin.H{"path": extSubtitles, "ext": vtt}, c)
				return
			}
		}
		response.OkWithData(gin.H{}, c)
		return
	}

	response.OkWithData(gin.H{"path": extSubtitles, "ext": vtt}, c)
}

func convertSubtitles(front, after string) error {
	// Open subtitles
	s1, _ := astisub.OpenFile(front)
	// Write subtitles
	err := s1.Write(after)
	return err
}
