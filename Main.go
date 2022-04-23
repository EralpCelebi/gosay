package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type RainbowFunction func(int) (int,int,int)

const Frequency = .1;
const Offset		= 2;

func ReadInput() string {
	Reader := bufio.NewReader(os.Stdin);
	String, _ := Reader.ReadString(0);

	return String;
}

func OutputMessage(Input string, Function RainbowFunction) {	
	Index, Color := 0, 0; 

	for Index < len(Input) {
		if Input[Index] == '\n' {
			Color = 0;
		}
		
		Red, Green, Blue := Function(Color);
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", Red, Green, Blue, Input[Index]);
		Index++;
		Color++;
	}
}

func DefaultRainbow(Index int) (int,int,int) {
	Shift := time.Now().Second();

	Red 	:= math.Sin(float64(Index + Shift) * Frequency) * 127 + 128;
	Green := math.Sin(float64(Index + Shift) * Frequency + Offset) * 127 + 128;
	Blue 	:= math.Sin(float64(Index + Shift) * Frequency + Offset * 2) * 127 + 128;

	return int(Red), int(Green), int(Blue)
}

func main() {	
	OutputMessage(ReadInput(), DefaultRainbow);
}
