package util

import "google.golang.org/api/classroom/v1"

type Material struct {
	Title string
	Link  string
}

func ParseClassroomMaterial(mat *classroom.Material) *Material {
	var (
		title string
		link  string
	)

	if mat.DriveFile != nil {
		// google drive file
		title = mat.DriveFile.DriveFile.Title
		link = mat.DriveFile.DriveFile.AlternateLink
	} else if mat.YoutubeVideo != nil {
		// video
		title = mat.YoutubeVideo.Title
		link = mat.YoutubeVideo.AlternateLink
	} else if mat.Link != nil {
		// direct link
		title = mat.Link.Title
		link = mat.Link.Url
	} else if mat.Form != nil {
		// form
		title = mat.Form.Title
		link = mat.Form.FormUrl
	} else {
		// unknown material type
		return nil
	}

	return &Material{
		Title: title,
		Link:  link,
	}
}
