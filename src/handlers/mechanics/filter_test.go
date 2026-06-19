package mechanics

import "testing"


func TestIsDocument(t *testing.T) {
	filename_doc := "test.docs"
	filename_code := "test.js"

	//Check true condition:
	if !isDocument(filename_doc) {
			t.Error("Test is fail! The requested name is a document!")	
	}
	if isDocument(filename_code) {
			t.Error("Test is fail! The requested name is a code!")	
	}
}

func TestIsImage(t *testing.T) {
	filename_image := "test.jpeg"
	filename_false := "test.odt"

	if !isImage(filename_image) {
			t.Error("Test is fail! The requested name is a image!")	
	}

	if isImage(filename_false) {
			t.Error("Test is fail! The requested name is a document!")	
	}
}


func TestIsVideo(t *testing.T) {
	filename_video := "test.mp4"
	filename_false := "test.java"
	filename_false2 := "test"

	if !isVideo(filename_video) {
			t.Error("Test is fail! The requested name is a video!")	
	}

	if isVideo(filename_false) {
			t.Error("Test is fail! The requested name is a code!")	
	}

	if isVideo(filename_false2) {
			t.Error("Test is fail! The requested name is a shit!")	
	}
}

func TestIsAudio(t *testing.T) {
	filename_audio := "test.mp3"
	filename_false := "test.gif"

	if !isAudio(filename_audio) {
			t.Error("Test is fail! The requested name is a audio!")	
	}

	if isAudio(filename_false) {
			t.Error("Test is fail! The requested name is a video!")	
	}
}

func TestIsData(t *testing.T) {
	filename_data := "test.json"
	filename_false := "test.wav"

	if !isData(filename_data) {
			t.Error("Test is fail! The requested name is a data!")	
	}

	if isData(filename_false) {
			t.Error("Test is fail! The requested name is a audio!")	
	}
}

func TestIsCode(t *testing.T) {
	filename_code := "test.go"
	filename_false := "test.odt"

	if !isCode(filename_code) {
			t.Error("Test is fail! The requested name is a code!")	
	}

	if isCode(filename_false) {
			t.Error("Test is fail! The requested name is a document!")	
	}
}


func TestFilter(t *testing.T) {
	data := "test.json"
	code := "test.rs"
	image := "unuu.jpeg"


	if result := Filter(data); result != "data" {
		t.Errorf("Error! The filter returned an %s result, the correct result: data", result)
	}

	if result := Filter(code); result != "code" {
		t.Errorf("Error! The filter returned an %s result, the correct result: code", result)
	}

	if result := Filter(image); result != "image" {
		t.Errorf("Error! The filter returned an %s result, the correct result: image", result)
	}
}
