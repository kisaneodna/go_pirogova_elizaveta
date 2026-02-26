package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("")
	}

	stepsStr := parts[0]
	activityType := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат шагов: %s", stepsStr)
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат продолжительности: %s", durationStr)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть положительной: %v", duration)
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLen := height * stepLengthCoefficient

	return (float64(steps) * stepLen) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64
	switch activityType {
	case "Ходьба":
		calories, _ = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, _ = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	return "Тип тренировки: " + activityType + "\n" +
		"Длительность: " + strconv.FormatFloat(duration.Hours(), 'f', 2, 64) + " ч.\n" +
		"Дистанция: " + strconv.FormatFloat(dist, 'f', 2, 64) + " км.\n" +
		"Скорость: " + strconv.FormatFloat(speed, 'f', 2, 64) + " км/ч\n" +
		"Сожгли калорий: " + strconv.FormatFloat(calories, 'f', 2, 64) + "\n", nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	speed := meanSpeed(steps, height, duration)
	calories := ((weight * speed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient

	return calories, nil

}
