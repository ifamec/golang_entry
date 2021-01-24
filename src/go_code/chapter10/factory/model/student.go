package model

type student struct {
	Name string
	score float64
}
// student start with lowercase, only available in this file

func NewStudent(n string, s float64) *student {
	return &student{n, s}
}
func (s *student) GetScore() float64 {
	return (*s).score
}
func (s *student) SetScore(n float64) {
	(*s).score = n
}