package redovalnica

import (
	"fmt"
)

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

var stOcen, MinOcena, MaxOcena int

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Napaka: študent z vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	if ocena < 0 || ocena > 10 {
		fmt.Println("Napaka: ocena mora biti med 0 in 10.")
		return
	}

	st.ocene = append(st.ocene, ocena)
	studenti[vpisnaStevilka] = st
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(st.ocene) < 6 {
		return 0
	}

	var vsota int
	for _, ocena := range st.ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(st.ocene))
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.ime, st.priimek, st.ocene)
	}
}
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		p := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.ime, student.priimek, p)

		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}
