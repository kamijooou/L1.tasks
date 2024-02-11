// паттерн адаптер применяется тогда, когда нужно к одному интерфейсу
// применять разное поведение, чтобы в итоге получить единую реализацию

// ATTENTION! Описанные классы - просто шутка, если кого-то задел, приношу
// свои искренние извинения.
package main

import (
	"fmt"
	"time"
)

type Mentor interface {
	Exam() bool
}

type Lovetskiy struct {
	Spec    string
	YouKnow bool
}

func (m *Lovetskiy) Exam() bool {
	fmt.Println("Ну что, обсудим базу?")
	if m.YouKnow {
		return true
	}
	return false
}

type Chertkov struct {
	Spec     string
	Task     string
	Answered bool
	Time     time.Duration
}

func (m *Chertkov) Exam() bool {
	fmt.Println("Достаем листочки, тянем билет")
	if m.Answered {
		return true
	}
	return false
}

type Roman struct {
	Spec     string
	Leetcode bool
	Base     bool
}

func (m *Roman) Exam() bool {
	fmt.Println("На старт, внимание... GO!")
	if m.Leetcode && m.Base {
		return true
	}
	return false
}

func startExam(mentor Mentor) {
	if mentor.Exam() {
		fmt.Println("Nice job, dude!")
		return
	}
	fmt.Println("YOU SHALL NOT PASS!")
}

func main() {
	ivan := &Lovetskiy{
		Spec:    "Golang under the hood",
		YouKnow: true,
	}
	maxim := &Chertkov{
		Spec:     "Questions owner",
		Task:     "Как обеспечить отказоустойчивость?",
		Answered: true,
		Time:     time.Duration(10 * time.Minute),
	}
	roman := &Roman{
		Spec:     "Base + leetcode",
		Leetcode: false,
		Base:     true,
	}

	startExam(ivan)
	startExam(maxim)
	startExam(roman)
}
