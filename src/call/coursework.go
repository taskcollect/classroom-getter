package call

import "google.golang.org/api/classroom/v1"

func ListCourseWork(srv *classroom.Service, course *classroom.Course) ([]*classroom.CourseWork, error) {
	resp, err := srv.Courses.CourseWork.List(course.Id).Do()

	if err != nil {
		return nil, err
	}

	return resp.CourseWork, nil
}

func GetCourseWorkByID(srv *classroom.Service, course *classroom.Course, id string) (*classroom.CourseWork, error) {
	work, err := srv.Courses.CourseWork.Get(course.Id, id).Do()

	if err != nil {
		return nil, err
	}

	return work, nil
}
