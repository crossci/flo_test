package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	r := gin.Default()
	r.GET("/shuffle/:cards", handle_shuffle)
	r.Run(":8080")
	//
	//s := []int{1, 2, 3, 4, 5}
	//fmt.Printf("%v \n", s)
	//rand.Seed(time.Now().Unix())
	//for i := 0; i < 100; i += 1 {
	//	randn(&s)
	//	fmt.Printf("%v \n", s)
	//	//fmt.Printf("%v \n", rand.Float64()*(4-0+1))
	//}
}
func handle_shuffle(c *gin.Context) {
	str := c.Param("cards")
	fmt.Printf("%v \n", str)
	cards := strings.Split(str, ",")
	randn(&cards)
	c.JSON(200, cards)
}
func get_random(floor, cell float64) float64 {
	return math.Floor(rand.Float64()*(cell-floor+1)) + floor
}
func swap(s []string, i, j int) {
	s[i], s[j] = s[j], s[i]
}
func randn(s *[]string) {
	n := len(*s)
	for i := n - 1; i > 0; i -= 1 {
		//j := rand.Int() % (i + 1)
		j := int(get_random(0, float64(i)))
		swap(*s, i, j)
	}
}
