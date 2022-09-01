package main
import(
		"fmt" //formatted I/O
		"github.com/gocolly/colly" //scraping framework
    "github.com/gocolly/colly/extensions"
//		"strconv"
//		"strings"
        "os"
        "encoding/csv"
        "log"
        "time"    
)

//Results:div.s-result-list.s-search-results.sg-row
//Items:div.a-section.a-spacing-base
//Name:span.a-size-base-plus.a-color-base.a-text-normal
//Price:span.a-price-whole
//Stars:span.a-icon-alt

func main(){
        fileName := "keyboards.csv"
        file, err := os.Create(fileName)
         if err != nil {
          log.Fatalf("Could not create %s", fileName)
        }
        defer file.Close()
       
        writer := csv.NewWriter(file)
        defer writer.Flush()
       
        writer.Write([]string{"Product Name", "Stars", "Price"})

		c := colly.NewCollector(colly.Async(false),colly.AllowedDomains("www.amazon.in"))
     //   extensions.RandomUserAgent(c)
		c.OnRequest(func(r *colly.Request){
				fmt.Println("Link of the page:", r.URL)
		})
        c.Limit(&colly.LimitRule{
        	RandomDelay: 5 * time.Second,
        	Parallelism: 4,
            })
		c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h*colly.HTMLElement){
				h.ForEach("div.a-section.a-spacing-base", func(_ int, h*colly.HTMLElement){
						var name string
						name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
						var stars string
						stars = h.ChildText("span.a-icon-alt")
					    var price string
                        price = h.ChildText("span.a-price-whole")
                    //	v := h.ChildText("span.a-price-whole")
					//	v = strings.Replace(v, ",", "", -1)
					//	price, err := strconv.ParseFloat(v, 32)
						if err != nil{
								fmt.Println("Error! Can't parse Price to Int")
								return
						}
					//	price = strconv.ParseInt(priceAsString)
					    writer.Write([]string{
                              name,
                              stars,
                              price,
                              })


          //  fmt.Println("ProductName: ", name)
					//	fmt.Println("Ratings: ", stars)
					//	fmt.Println("Price: ", price)

				})
		})

       for i :=1; i<=1; i++{    //change this to scrape multiple pages
       URL := fmt.Sprintf("https://www.amazon.in/s?k=keyboard&page=%d", i)
   c.Visit(URL)
	}
}
