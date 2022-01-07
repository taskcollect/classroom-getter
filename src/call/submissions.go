package call

import "google.golang.org/api/classroom/v1"

func ListAllActiveSubmissions(srv *classroom.Service, course *classroom.Course) ([]*classroom.StudentSubmission, error) {
	resp, err := srv.Courses.CourseWork.StudentSubmissions.List(course.Id, "-").States("NEW", "RECLAIMED_BY_STUDENT", "TURNED_IN").Do()

	if err != nil {
		return nil, err
	}

	return resp.StudentSubmissions, nil
}
