package utils

var storePath = "/files/"

func GetFilePath(filename string) string {
	switch ParseFileType(filename) {
	case "mp4":
		return storePath + "videos/" + filename
	case "mov":
		return storePath + "videos/" + filename
	case "png":
		return storePath + "images/" + filename
	case "jpg":
		return storePath + "images/" + filename
	case "jpeg":
		return storePath + "images/" + filename
	default:
		return "empty_path"
	}
}

func ParseFileName(filepath string) string {
	s := filepath
	i := len(s) - 1
	for i >= 0 && s[i] != '/' {
		i--
	}
	return s[i+1:]
}

func ParseFileType(filename string) string {
	s := filename
	i := len(s) - 1
	for i >= 0 && s[i] != '.' {
		i--
	}
	return s[i+1:]
}
