package communication

import (
	"fmt"
	"strings"
)

type Good struct {
	ShopName string
	Name     string
	Price    int
	OldPrice int
	ImageUrl string
	Url      string
	Discount int
}

func (g *Good) ToStringForTelegram() string {
	q := "#%v\n\n%v\n\n*%v*\n\n%v"

	return fmt.Sprintf(q, strings.ToUpper(g.ShopName), getDiscount(g.Discount, g.ImageUrl), g.Name, g.Url)
}

func getDiscount(discount int, image string) string {
	res := "[🔥](%v) *Скидка %v%s* 🔥"
	if discount >= 80 {
		res = "🔥[🔥](%v)🔥 *Скидка %v%s* 🔥🔥🔥"
	}
	if discount >= 50 && discount < 80 {
		res = "🔥[🔥](%v) *Скидка %v%s* 🔥🔥"
	}
	return fmt.Sprintf(res, image, discount, "%")
}
