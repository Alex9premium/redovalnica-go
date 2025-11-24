# redovalnica-go

Paket, ki ponuja redovalnico. Trenutno naredi predefiniranega študenta in mu dodeli nekaj ocen

Ponuja funkcije kot so dodajOceno, IzpisVsehOcen, IzpisiKoncniUspeh.

Možna stikala:

StOcen - ki definira najmanjše število ocen potrebnih za pozitivno oceno

MinOcena = najmanjša možna ocena

MaxOcena = največja možna ocena

Uporaba:

```go
s := redovalnica.NoviStudent("Janez", "Novak")
studenti["63230111"] = redovalnica.NoviStudent("Janez", "Novak")

DodajOceno(studenti, "12345", 8)
DodajOceno(studenti, "12345", 9)
IzpisVsehOcen(studenti)
IzpisiKoncniUspeh(studenti)
```
