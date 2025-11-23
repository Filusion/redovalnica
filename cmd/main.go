package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Filusion/redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "redovalnica-cli",
		Usage: "Simple grade register CLI using package redovalnica",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 5,
				Usage: "Minimum number of grades required for positive status",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 0,
				Usage: "Minimum allowed grade",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "Maximum allowed grade",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			// read CLI flags
			stOcen := c.Int("stOcen")
			minOc := c.Int("minOcena")
			maxOc := c.Int("maxOcena")

			// create a new Redovalnica
			r := redovalnica.New(stOcen, minOc, maxOc)

			// add some sample students
			if err := r.AddStudent("35755313", "Johnny", "Depp", []int{10, 10, 9, 9, 8}); err != nil {
				fmt.Printf("Student wasn't added. AddStudent error: %v\n\n", err)
			}

			if err := r.AddStudent("29013526", "Jack", "Black", []int{8, 9, 6, 7, 8}); err != nil {
				fmt.Printf("Student wasn't added. AddStudent error: %v\n\n", err)
			}

			if err := r.AddStudent("63233333", "Peter", "Griffin", []int{5, 5, 4, 7, 6}); err != nil {
				fmt.Printf("Student wasn't added. AddStudent error: %v\n\n", err)
			}

			// print all grades
			fmt.Println(r.IzpisVsehOcen())
			fmt.Println("----- KONČNI USPEH -----")
			fmt.Print(r.IzpisiKoncniUspeh())

			// example: add a new grade
			if err := r.DodajOceno("35755313", 9); err != nil {
				fmt.Printf("DodajOceno error: %v\n", err)
			} else {
				fmt.Printf("Added grade to 35755313\n\n")
			}

			fmt.Println("After adding a grade:\n-----------------------------")
			fmt.Println(r.IzpisVsehOcen())

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
