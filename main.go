package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book represents the structure of a book item to be stored in MongoDB.
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	PublishDate string `json:"publishDate" bson:"publishDate"`
	Rating      string `json:"rating"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// Function for extracting the rating and year from a raw string
func ExtractYear(raw string) string {
	raw = strings.TrimSpace(raw)
	// Regex for year (last 4 digits)
	yearRegex := regexp.MustCompile(`\d{4}`)
	year := yearRegex.FindString(raw)

	return year
}

func ParseBook(bookURL string) (Book, error) {
	res, err := http.Get(bookURL)
	if err != nil {
		return Book{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Book{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Book{}, err
	}

	book := Book{}
	book.Title = strings.TrimSpace(doc.Find(".Text.Text__title1").Text())
	book.Rating = strings.TrimSpace(doc.Find(".RatingStatistics__rating").Text())
	rawText := strings.TrimSpace(doc.Find(".publicationInfo").Text())
	book.PublishDate = ExtractYear(rawText)
	book.Description = strings.TrimSpace(doc.Find("div[data-testid='description'] span.Formatted").Text())
	book.Author = strings.TrimSpace(doc.Find(".ContributorLink__name").Text())
	if img, ok := doc.Find("img.ResponsiveImage").Attr("src"); ok {
		book.Cover = strings.TrimSpace(img)
	}
	return book, nil
}

func main() {
	//  Loading the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//  Connecting to MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DATABASE")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
		fmt.Println("❌ Issue connecting to MongoDB!")
	}
	fmt.Println("✅ Connected to MongoDB!")
	collection := client.Database(mongoDB).Collection("books")
	// Rabbit MQ connection
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

}
