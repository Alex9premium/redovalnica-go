package main

import (
	"context"
	"log"
	"os"

	"github.com/Alex9premium/redovalnica-go/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "test_redovalnice",
		Usage: "Redovalnica shows student grades and averages, also supports adding grades.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 6,
				Usage: "minimum number of grades for positive score",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 1,
				Usage: "minimum possible score",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "maximum possible score",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			redovalnica.StOcen = cmd.Int("stOcen")
			redovalnica.MinOcena = cmd.Int("minOcena")
			redovalnica.MaxOcena = cmd.Int("maxOcena")

			studenti := map[string]redovalnica.Student{}
			studenti["63230111"] = redovalnica.NoviStudent("Janez", "Novak")
			redovalnica.DodajOceno(studenti, "63230111", 8)
			redovalnica.DodajOceno(studenti, "63230111", 9)
			redovalnica.DodajOceno(studenti, "63230111", 10)
			redovalnica.DodajOceno(studenti, "63230111", 7)
			redovalnica.DodajOceno(studenti, "63230111", 6)
			redovalnica.DodajOceno(studenti, "63230111", 9)

			redovalnica.IzpisVsehOcen(studenti)
			redovalnica.IzpisiKoncniUspeh(studenti)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
