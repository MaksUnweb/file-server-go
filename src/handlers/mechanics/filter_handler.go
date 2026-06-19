package mechanics

import "strings"


func Filter(filename string) string {
	switch {
	case isDocument(filename): 
		return "document"
	case isImage(filename):
		return "image"
	case isVideo(filename):
		return "video"
	case isAudio(filename):
		return "audio"
	case isData(filename):
		return "data"
	case isCode(filename):
		return "code"
	default:
		return "undefined"
	}
}


//PDF, DOC, DOCX, TXT, MD, RTF
func isDocument(filename string) bool {
	return strings.Contains(filename, ".doc") ||
	strings.Contains(filename, ".pdf") || 
	strings.Contains(filename, ".docs") ||
	strings.Contains(filename, ".txt") ||
	strings.Contains(filename, ".md") || 
	strings.Contains(filename, ".rtf") 
}

//JPG, PNG, GIF, WEBP, SVG, BMP
func isImage(filename string) bool {
	return strings.Contains(filename, ".jpg") ||
		strings.Contains(filename, ".jpeg") ||
		strings.Contains(filename, ".png") || 
		strings.Contains(filename, ".gif") || 
		strings.Contains(filename, ".webp") || 
		strings.Contains(filename, ".svg") || 
		strings.Contains(filename, ".bmp")
}

//MP4, AVI, MKV, MOV, WEBM
func isVideo(filename string) bool {
	return strings.Contains(filename, ".mp4")	||
			strings.Contains(filename, ".avi")	||
			strings.Contains(filename, ".mkv")	||
			strings.Contains(filename, ".mov")	||
			strings.Contains(filename, ".webm")
}

//MP3, WAV, FLAC, OGG, M4A
func isAudio(filename string) bool {
	return  strings.Contains(filename, ".mp3") || 
			strings.Contains(filename, ".wav") ||
			strings.Contains(filename, ".flac") ||
			strings.Contains(filename, ".ogg") ||
			strings.Contains(filename, ".m4a")
}

//JSON, XML, CSV, YAML, TOML
func isData(filename string) bool {
	return strings.Contains(filename, ".json") ||
			strings.Contains(filename, ".xml") ||
			strings.Contains(filename, ".csv") ||
			strings.Contains(filename, ".yaml") ||
			strings.Contains(filename, ".toml")
}

// .go, .py, .js, .java, .rs, .cpp, .html, .css
func isCode(filename string) bool {
	return strings.Contains(filename, ".go") ||
			strings.Contains(filename, ".py") ||
			strings.Contains(filename, ".js") ||
			strings.Contains(filename, ".ts") ||
			strings.Contains(filename, ".java") ||
			strings.Contains(filename, ".rs") ||
			strings.Contains(filename, ".cpp") ||
			strings.Contains(filename, ".c") ||
			strings.Contains(filename, ".cs") ||
			strings.Contains(filename, ".tsx") ||
			strings.Contains(filename, ".jsx") ||
			strings.Contains(filename, ".h") ||
			strings.Contains(filename, ".hpp") ||
			strings.Contains(filename, ".php") ||
			strings.Contains(filename, ".rb") ||
			strings.Contains(filename, ".swift") ||
			strings.Contains(filename, ".ex") ||
			strings.Contains(filename, ".exs") ||
			strings.Contains(filename, ".lua") ||
			strings.Contains(filename, ".vim") ||
			strings.Contains(filename, ".r") ||
			strings.Contains(filename, ".asm") ||
			strings.Contains(filename, ".bash") ||
			strings.Contains(filename, ".psl") 
}


