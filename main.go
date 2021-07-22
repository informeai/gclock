package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"time"
)

func loopClock(seconds int) {
	//write clock in center terminal
	area, err := pterm.DefaultArea.WithCenter().Start()
	if err != nil {
		panic(err)
	}
	for {

		str, err := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(fmt.Sprintf("%v:%v", seconds/60, seconds%60))).Srender()
		if err != nil {
			panic(err)
		}
		area.Update(str)
		time.Sleep(time.Second)
		if seconds == 0 {
			break
		}
		seconds--
	}
	area.Stop()

}
func run(pomo string) {
	//loopClock
	if pomo == "pomo" {
		loopClock(1500)
	} else {
		loopClock(300)
	}
}

func main() {
	//header
	pterm.DefaultCenter.Println(pterm.FgYellow.Sprint("\U000023F1 GCLOCK"))
	pterm.DefaultCenter.Println(pterm.FgLightYellow.Sprint("Clock for recording time using the pomodoro technique."))

	pterm.DefaultCenter.Println(pterm.FgCyan.Sprint("made with \U00002764 by wellington gadelha"))
	fmt.Println("\n\n")
	var pomo int = 0
	pterm.DefaultCenter.Println("Pomodoros: ", pterm.FgLightCyan.Sprint(pomo))
	//state pomodoro
	var pomodoroOrRepose string = "pomo"
	run(pomodoroOrRepose)

}
