package course

import "fmt"

type Service interface {
	Describe(courseID uint64) (*Course, error)
	List(cursor uint64, limit uint64) ([]Course, error)
	Create(course Course) (uint64, error)
	Update(courseID uint64, course Course) error
	Remove(courseID uint64) (bool, error)
}

type DummyService struct {
	courses []Course
}

func NewDummyService() *DummyService {
	var courses []Course
	for i := 0; i < 33; i++ {
		courses = append(courses, Course{
			ID:          uint64(i),
			Title:       fmt.Sprintf("[%d] Course title", i),
			Description: fmt.Sprintf("[%d] Course desc", i),
		})
	}

	return &DummyService{courses: courses}
}

func (d *DummyService) Describe(courseID uint64) (*Course, error) {
	for _, course := range d.courses {
		if course.ID == courseID {
			return &course, nil
		}
	}
	return nil, fmt.Errorf("course %d not found", courseID)
}

func (d *DummyService) List(cursor uint64, limit uint64) ([]Course, error) {
	if cursor >= uint64(len(d.courses)) {
		return nil, fmt.Errorf("cursor out of bounds")
	}

	end := cursor + limit
	if end > uint64(len(d.courses)) {
		end = uint64(len(d.courses))
	}

	return d.courses[cursor:end], nil
}

func (d *DummyService) Create(course Course) (uint64, error) {
	course.ID = uint64(len(d.courses))
	d.courses = append(d.courses, course)
	return course.ID, nil
}

func (d *DummyService) Update(courseID uint64, course Course) error {
	for i, c := range d.courses {
		if c.ID == courseID {
			course.ID = c.ID
			d.courses[i] = course
			return nil
		}
	}
	return fmt.Errorf("course %d not found", courseID)
}

func (d *DummyService) Remove(courseID uint64) (bool, error) {
	for i, c := range d.courses {
		if c.ID == courseID {
			d.courses = append(d.courses[:i], d.courses[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
