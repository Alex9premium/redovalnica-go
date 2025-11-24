// Package redovalnica ponuja redovalnico za študente.
//
// Paket omogoča dodajanje ocen študentom, izpis vseh ocen, izračunavanje
// končnega uspeha na podlagi ocen in določenih
// nastavitev.
//
//
// Primer uporabe:
//
//    Možna stikala:
//    StOcen - ki definira najmanjše število ocen potrebnih za pozitivno oceno
//    MinOcena = najmanjša možna ocena
//    MaxOcena = največja možna ocena
//
//    Ustvari, novega studenta:
//    s := redovalnica.NoviStudent("Janez", "Novak")
//    studenti["63230111"] = redovalnica.NoviStudent("Janez", "Novak")
//
//    Funkcije:
//    DodajOceno(studenti, "12345", 8)
//    DodajOceno(studenti, "12345", 9)
//    IzpisVsehOcen(studenti)
//    IzpisiKoncniUspeh(studenti)

package redovalnica

import (
	"fmt"
)

// Možnosti za študenta
type Student struct {
	ime     string
	priimek string
	ocene   []int
}

// Nastavitve za redovalnico
var StOcen, MinOcena, MaxOcena int


func NoviStudent(ime, priimek string) Student {
    return Student{
        ime:     ime,
        priimek: priimek,
        ocene:   []int{},
    }
}

// DodajOceno doda oceno študentu z dano vpisno številko.
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Napaka: študent z vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	if ocena < MinOcena || ocena > MaxOcena {
		fmt.Printf("Napaka: ocena mora biti med %d in %d.\n", MinOcena, MaxOcena)
		return
	}

	st.ocene = append(st.ocene, ocena)
	studenti[vpisnaStevilka] = st
}

// Povprecje izračuna povprečno oceno študenta z dano vpisno številko.
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(st.ocene) < StOcen {
		return 0
	}

	var vsota int
	for _, ocena := range st.ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(st.ocene))
}

// IzpisVsehOcen izpiše vse ocene vseh študentov.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.ime, st.priimek, st.ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh vseh študentov.
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
