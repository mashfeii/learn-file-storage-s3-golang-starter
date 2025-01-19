package pkg

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func GetVideoAspectRatio(filePath string) (string, error) {
	type output struct {
		Streams []struct {
			DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		} `json:"streams"`
	}

	command := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	buffer := bytes.NewBuffer(nil)
	command.Stdout = buffer

	err := command.Run()
	if err != nil {
		return "", err
	}

	var out output

	err = json.Unmarshal(buffer.Bytes(), &out)
	if err != nil || len(out.Streams) == 0 {
		return "", err
	}

	switch out.Streams[0].DisplayAspectRatio {
	case "16:9":
		return "landscape", nil
	case "9:16":
		return "portrait", nil
	default:
		return "other", nil
	}
}
