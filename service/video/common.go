package video

var storePath = "/home/ubuntu/go/src/tiktok/public/"

func getFilePath(filename string) string {
	switch parseFileType(filename) {
	case "mp4":
		return storePath + "/video/" + filename
	case "png":
		return storePath + "/picture/" + filename
	default:
		return "empty_path"
	}
}

func parseFileName(filepath string) string {
	s := filepath
	i := len(s) - 1
	for i >= 0 && s[i] != '/' {
		i--
	}
	return s[i+1:]
}

func parseFileType(filename string) string {
	s := filename
	i := len(s) - 1
	for i >= 0 && s[i] != '.' {
		i--
	}
	return s[i+1:]
}
