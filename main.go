package main

import (
	"fmt"
	"log"
	"os"
)

// Helper function I don't want to do thing
func parseLine(s string) [15]bool {
	if len(s) != 15 {
		log.Fatal("Line needs to be 16 characters")
	}
	var line [15]bool
	for i, c := range s {
		if c == '#' {
			line[i] = true
		} else {
			line[i] = false
		}
	}
	return line
}

func main() {
	var rooms []Room

	var room0 Room
	room0.Solid[0] = parseLine("###############")
	room0.Solid[1] = parseLine("#_____________#")
	room0.Solid[2] = parseLine("#_##_#####_##_#")
	room0.Solid[3] = parseLine("#_#_________#_#")
	room0.Solid[4] = parseLine("#_#_###_###_#_#")
	room0.Solid[5] = parseLine("#_#_________#_#")
	room0.Solid[6] = parseLine("#_##_##_##_##_#")
	room0.Solid[7] = parseLine("#_#__#___#__#_#")
	room0.Solid[8] = parseLine("#___#######___#")
	room0.Solid[9] = parseLine("#_#_________#_#")
	room0.Solid[10] = parseLine("#_#_##_#_##_#_#")
	room0.Solid[11] = parseLine("#_#_________#_#")
	room0.Solid[12] = parseLine("#_####_#_####_#")
	room0.Solid[13] = parseLine("#_____________#")
	room0.Solid[14] = parseLine("###############")
	rooms = append(rooms, room0)

	var room1 Room
	room1.Solid[0] = parseLine("#######_#######")
	room1.Solid[1] = parseLine("#_____________#")
	room1.Solid[2] = parseLine("#_#####_#####_#")
	room1.Solid[3] = parseLine("#_#_________#_#")
	room1.Solid[4] = parseLine("#_#_###_###_#_#")
	room1.Solid[5] = parseLine("#_#_________#_#")
	room1.Solid[6] = parseLine("#_#__##_##__#_#")
	room1.Solid[7] = parseLine("#_#_##___##_#_#")
	room1.Solid[8] = parseLine("_______#_______")
	room1.Solid[9] = parseLine("#_#_#######_#_#")
	room1.Solid[10] = parseLine("#_#_#_#_#_#_#_#")
	room1.Solid[11] = parseLine("#_#_________#_#")
	room1.Solid[12] = parseLine("#_#####_#####_#")
	room1.Solid[13] = parseLine("#_____________#")
	room1.Solid[14] = parseLine("#######_#######")
	rooms = append(rooms, room1)

	var room2 Room
	room2.Solid[0] = parseLine("#######_#######")
	room2.Solid[1] = parseLine("#___#_____#___#")
	room2.Solid[2] = parseLine("#_#_##_#_##_#_#")
	room2.Solid[3] = parseLine("#_#__#_#_#__#_#")
	room2.Solid[4] = parseLine("#_##_______##_#")
	room2.Solid[5] = parseLine("#____##_##____#")
	room2.Solid[6] = parseLine("#_##_#___#_##_#")
	room2.Solid[7] = parseLine("#_#__#####__#_#")
	room2.Solid[8] = parseLine("#___###_###___#")
	room2.Solid[9] = parseLine("#_#_#_____#_#_#")
	room2.Solid[10] = parseLine("#_#___#_#___#_#")
	room2.Solid[11] = parseLine("#_##_##_##_##_#")
	room2.Solid[12] = parseLine("#_#___#_#___#_#")
	room2.Solid[13] = parseLine("#___#_____#___#")
	room2.Solid[14] = parseLine("#######_#######")
	rooms = append(rooms, room2)

	var room3 Room
	room3.Solid[0] = parseLine("###########_###")
	room3.Solid[1] = parseLine("#___#_____#___#")
	room3.Solid[2] = parseLine("#_#___###___#_#")
	room3.Solid[3] = parseLine("#_#_#__#__#_#_#")
	room3.Solid[4] = parseLine("#___##_#_##___#")
	room3.Solid[5] = parseLine("###_________###")
	room3.Solid[6] = parseLine("____###_###____")
	room3.Solid[7] = parseLine("#_#__#___#__#_#")
	room3.Solid[8] = parseLine("#_##_#####_##_#")
	room3.Solid[9] = parseLine("#__#_______#__#")
	room3.Solid[10] = parseLine("##___#####___##")
	room3.Solid[11] = parseLine("#__#___#___#__#")
	room3.Solid[12] = parseLine("#_####___####_#")
	room3.Solid[13] = parseLine("#______#______#")
	room3.Solid[14] = parseLine("###########_###")
	rooms = append(rooms, room3)

	// Generate csv
	err := os.Remove("out.csv")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("out.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, room := range rooms {
		sT := room.GenerateMap()
		for goalY := 0; goalY < 15; goalY++ {
			for goalX := 0; goalX < 15; goalX++ {
				v, ok := sT[Pos{goalX, goalY}]
				amIOkay := true
				if !ok {
					amIOkay = false // :(
				}
				for y := 0; y < 15; y++ {
					for x := 0; x < 15; x++ {
						var value = uint(0)
						if amIOkay {
							v, ok := (*v)[Pos{x, y}]
							if ok {
								value = v.Direction
							}
						}
						_, err := fmt.Fprintf(file, "%1d, ", value)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
}
