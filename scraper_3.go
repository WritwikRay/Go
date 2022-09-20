package main
import(
		"fmt" //formatted I/O
		"github.com/gocolly/colly" //scraping framework
		"strconv"
		"strings"
		"os"
		"bufio"
)

//Results:div.s-result-list.s-search-results.sg-row
//Items:div.a-section.a-spacing-base
//Name:span.a-size-base-plus.a-color-base.a-text-normal
//Price:span.a-price-whole
//Stars:span.a-icon-alt

//

func main(){
	//	var search_word string
		fmt.Print("Please insert the amazon product link here: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		link := scanner.Text()
	//	search_word = strings.ReplaceAll(search_word, " ", "%20")
		fmt.Println(link)
		c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

		c.OnRequest(func(r *colly.Request){
				fmt.Println("Link of the page:", r.URL)
		})

		c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h*colly.HTMLElement){
				h.ForEach("div.a-section.a-spacing-base", func(_ int, h*colly.HTMLElement){
						var name string
						name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
						var stars string
						stars = h.ChildText("span.a-icon-alt")
						v := h.ChildText("span.a-price-whole")
						v = strings.Replace(v, ",", "", -1)
						price, err := strconv.ParseFloat(v, 32)
						if err != nil{
								fmt.Println("Error! Can't parse Price to Int")
								return
						}
					//	price = strconv.ParseInt(priceAsString)
						fmt.Println("ProductName: ", name)
						fmt.Println("Ratings: ",stars)
						fmt.Println("Price: ",price)

				})
		})
	//	URL := fmt.Sprintf("https://www.amazon.in/s?k=%s&ref=nb_sb_noss", search_word)
        c.Visit(link)
	}
