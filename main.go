package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	e := echo.New()

	e.GET("/min", min)
	e.GET("/max", max)
	e.GET("/avg", avg)
	e.GET("/median", median)
	e.GET("/percentile", percentile)

	e.Logger.Fatal(e.Start(":9191"))
}

func min(c echo.Context) error {
	numbers := strings.Split(c.QueryParam("numbers"), ",")
	intNumbers, _ := sliceAtoi(numbers)
	min := Min(intNumbers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"min": min,
	})
}

func max(c echo.Context) error {
	numbers := strings.Split(c.QueryParam("numbers"), ",")
	intNumbers, _ := sliceAtoi(numbers)
	max := Max(intNumbers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"max": max,
	})
}

func avg(c echo.Context) error {
	numbers := strings.Split(c.QueryParam("numbers"), ",")
	intNumbers, _ := sliceAtoi(numbers)
	avg := Avg(intNumbers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"avg": avg,
	})
}

func median(c echo.Context) error {
	numbers := strings.Split(c.QueryParam("numbers"), ",")
	intNumbers, _ := sliceAtoi(numbers)
	median := Median(intNumbers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"median": median,
	})
}

func percentile(c echo.Context) error {
	numbers := strings.Split(c.QueryParam("numbers"), ",")
	intNumbers, _ := sliceAtoi(numbers)

	percent, _ := strconv.ParseFloat(c.QueryParam("percent"), 64)

	percentile := PercentileNearestRank(intNumbers, percent)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"percentile": percentile,
	})
}
