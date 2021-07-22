package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pterm/pterm"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//playSound execute play of song
func playSound() {
	f, err := os.Open(os.Getenv("song"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

//loopClock execute loop events
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
	//play sound
	playSound()

}

//run function in the loop
func run(pomo string) {
	//loopClock
	if pomo == "pomo" {
		loopClock(1500)
	} else if pomo == "repose" {
		loopClock(300)
	}
}

//next execute new interaction.
func next(count int) {
	pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgCyan)).Print("\U0001F449 [1]NEW POMODORO, [2]REPOSE, [3]QUIT: ")
	//reader
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}

	switch char {
	case '1':
		count++
		Start(count, "pomo")
	case '2':
		Start(count, "repose")
	case '3':
		os.Exit(0)
	default:

	}

}

//initialise function prepare for execution.
func initialise(count int, pomo string) {
	//header
	fmt.Println("\n\n")
	pterm.DefaultCenter.Println(pterm.FgYellow.Sprint("\U000023F1 GCLOCK"))
	pterm.DefaultCenter.Println(pterm.FgLightYellow.Sprint("Clock for recording time using the pomodoro technique."))

	pterm.DefaultCenter.Println(pterm.FgCyan.Sprint("made with \U00002764 by wellington gadelha"))
	fmt.Println("\n\n")

	pterm.DefaultCenter.Println("Pomodoros: ", pterm.FgLightCyan.Sprint(count))
	if count == 0 {
		area, err := pterm.DefaultArea.WithCenter().Start()
		if err != nil {
			panic(err)
		}
		str, err := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(fmt.Sprint("25:00"))).Srender()
		if err != nil {
			panic(err)
		}
		area.Update(str)

	} else {
		run(pomo)
	}
}

//clean function execute clear in terminal.
func clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

//Start function execute all functions the events.
func Start(count int, pomo string) {
	clear()
	initialise(count, pomo)
	next(count)
}

func main() {
	os.Setenv("song", "./song/alarm.mp3")
	count := 0
	Start(count, "pomo")

}
