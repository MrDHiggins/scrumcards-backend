package utils

import (
	"math"
	"strconv"

	"github.com/MrDHiggins/scrumdcards-backend/internal/models"
)

func CalculateVoteAverage(votes map[string]*models.Vote) float64 {
	var sum, count int
	for _, v := range votes {
		if num, err := strconv.Atoi(v.Value); err == nil {
			sum += num
			count++
		}
	}
	if count == 0 {
		return 0
	}

	avg := float64(sum) / float64(count)
	return math.Round(avg*100) / 100
}
