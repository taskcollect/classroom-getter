package call

import "google.golang.org/api/classroom/v1"

func ListCourses(srv *classroom.Service) ([]*classroom.Course, error) {
	resp, err := srv.Courses.List().CourseStates("ACTIVE").Do()

	if err != nil {
		return nil, err
	}

	return resp.Courses, nil
}
