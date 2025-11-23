// Package redovalnica provides a simple grade register for students.
//
// The package exports three functions/methods required by the assignment:
//   - DodajOceno
//   - IzpisVsehOcen
//   - IzpisiKoncniUspeh
//
// The helper function povprecje is intentionally unexported.
package redovalnica

import (
	"errors"
	"fmt"
	"sort"
)

// Student represents a student in the register.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Redovalnica holds students and configuration for valid grades.
type Redovalnica struct {
	studenti map[string]Student
	StOcen   int
	MinOcena int
	MaxOcena int
}

// New creates a new Redovalnica with the given configuration.
func New(stOcen, minOcena, maxOcena int) *Redovalnica {
	return &Redovalnica{
		studenti: make(map[string]Student),
		StOcen:   stOcen,
		MinOcena: minOcena,
		MaxOcena: maxOcena,
	}
}

// AddStudent adds a new student record (helper utility used by main).
func (r *Redovalnica) AddStudent(vpisna, ime, priimek string, ocene []int) error {
	for _, ocena := range ocene {
		if ocena < r.MinOcena || ocena > r.MaxOcena {
			return fmt.Errorf("grade %d out of bounds (%d..%d)", ocena, r.MinOcena, r.MaxOcena)
		}
	}
	r.studenti[vpisna] = Student{Ime: ime, Priimek: priimek, Ocene: ocene}
	return nil
}

// DodajOceno adds a grade to a student identified by vpisna.
func (r *Redovalnica) DodajOceno(vpisna string, ocena int) error {
	if ocena < r.MinOcena || ocena > r.MaxOcena {
		return fmt.Errorf("grade %d out of bounds (%d..%d)", ocena, r.MinOcena, r.MaxOcena)
	}
	s, ok := r.studenti[vpisna]
	if !ok {
		return errors.New("student does not exist")
	}
	s.Ocene = append(s.Ocene, ocena)
	r.studenti[vpisna] = s
	return nil
}

// IzpisVsehOcen returns a textual representation of the whole register.
func (r *Redovalnica) IzpisVsehOcen() string {
	keys := make([]string, 0, len(r.studenti))
	for k := range r.studenti {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := "REDOVALNICA:\n"
	for _, k := range keys {
		s := r.studenti[k]
		out += fmt.Sprintf("%s - %s %s: %v\n", k, s.Ime, s.Priimek, s.Ocene)
	}
	return out
}

// IzpisiKoncniUspeh returns a textual summary of final results for all students.
func (r *Redovalnica) IzpisiKoncniUspeh() string {
	keys := make([]string, 0, len(r.studenti))
	for k := range r.studenti {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := ""
	for _, k := range keys {
		s := r.studenti[k]
		avg := r.povprecje(k)
		out += fmt.Sprintf("%s %s: povprečna ocena %.1f -> ", s.Ime, s.Priimek, avg)
		if avg == -1 {
			out += "No grades\n"
			continue
		}
		if len(s.Ocene) < r.StOcen {
			out += "Neuspešen študent (premalo ocen)\n"
			continue
		}
		if avg >= 9 {
			out += "Odličen student!\n"
		} else if avg >= 6 {
			out += "Povprečen študent\n"
		} else {
			out += "Neuspešen študent\n"
		}
	}
	out += "\n"
	return out
}

// povprecje computes the average grade for student vpisna.
func (r *Redovalnica) povprecje(vpisna string) float64 {
	s, ok := r.studenti[vpisna]
	if !ok {
		return -1
	}
	if len(s.Ocene) == 0 {
		return -1
	}
	var sum float64
	for _, g := range s.Ocene {
		sum += float64(g)
	}
	avg := sum / float64(len(s.Ocene))
	if avg < 6 {
		return 0.0
	}
	return avg
}
