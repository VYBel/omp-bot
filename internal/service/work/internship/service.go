package internship

import (
	"errors"
	"log"
	"strconv"
)

type WorkService interface {
	List(cursor uint64, limit uint64) []Internship
	Describe(internshipID uint64) (*Internship, error)
	Remove(internshipID uint64) (bool, error)
	Update(internshipID uint64, internship Internship) error
	Create(Internship) (uint64, error)
	ShortString(p Internship) string
	FullString(p Internship) string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) []Internship {
	return allEntities
}

func (s *Service) Describe(internshipID uint64) (*Internship, error) {
	var resId uint64
	var result bool = false
	for i := 0; i < len(allEntities); i++ {
		if internshipID == allEntities[i].Id {
			resId = uint64(i)
			result = true
			break
		}
	}
	if !result {
		err := errors.New("id not found")
		log.Printf("intership.Service.Get: id not found - %v", err)
		return nil, err
	}
	return &allEntities[resId], nil
}

func (s *Service) Remove(internshipID uint64) (bool, error) {
	var result bool = false
	for i := 0; i < len(allEntities); i++ {
		if internshipID == allEntities[i].Id {
			result = true
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			break
		}
	}
	return result, nil
}

func (s *Service) Update(internshipID uint64, internship Internship) error {
	err := errors.New("not implemented yet")
	return err
}

func (s *Service) Create(internship Internship) (uint64, error) {
	err := errors.New("not implemented yet")
	return 0, err
}

func (s *Service) ShortString(p Internship) string {
	var result string = ""
	result += "ID: " + strconv.FormatUint(p.Id, 10) + ":"
	if &p.Description != nil {
		result += " Description: " + p.Description
	}
	return result
}

func (s *Service) FullString(p Internship) string {
	var result string = ""
	result += "ID: " + strconv.FormatUint(p.Id, 10) + ":"
	result += " Team - " + strconv.FormatUint(p.Team_id, 10) + ";"
	if &p.Description != nil {
		result += " Description: " + p.Description + ";"
	}
	if &p.Period != nil {
		result += " Period: " + p.Period + ";"
	}
	if &p.Compensation != nil {
		if p.Compensation {
			result += " compensation: yes."
		} else {
			result += " compensation: no."
		}
	}
	return result
}
