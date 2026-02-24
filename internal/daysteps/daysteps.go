package daysteps

import (
	"errors"
	"log"
	"strconv"
	"time"
	"strings"

	"github.com/kisaneodna/go_pirogova_elizaveta.git/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
	return 0, 0, errors.New("")
}

	stepsStr := parts[0]
	durationStr := strings.TrimSpace(parts[1])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return 0, 0, errors.New("")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, 0, errors.New("")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
	log.Println("количество шагов должно быть положительным")
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return "Количество шагов: " + strconv.Itoa(steps) + ".\n" +
		"Дистанция составила " + strconv.FormatFloat(distance, 'f', 2, 64) + " км.\n" +
		"Вы сожгли " + strconv.FormatFloat(calories, 'f', 2, 64) + " ккал.\n"
}
