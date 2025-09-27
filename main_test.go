package main

import (
	"testing"
)

// Simplified example of the Good reads html structure for the book "To Kill a Mockingbird"
const HTML = `
<div id="BookPage__leftColumn">
	<img class="ResponsiveImage" role="presentation" src="https://images-na.ssl-images-amazon.com/images/S/compressed.photo.goodreads.com/books/1612238791i/56916837.jpg" loading="eager" onerror="this.onerror=null;this.src='https://dryofg8nmyqjw.cloudfront.net/images/no-cover.png';">
</div> 

<div class="BookPage__rightColumn>
	<div class="BookPageTitleSection">
		<h1 class="Text Text__title1" data-testid="bookTitle" aria-label="Book title:To Kill a Mockingbird">To Kill a Mockingbird</h1>
	</div>
	<div class="ContributorLinksList">
		<span class="ContributorLink__name" data-testid="name">Harper Lee</span>
	</div>
	<p data-testid="publicationInfo">First published July 11, 1960</p>
	<div class="RatingStatistics__rating">4.26</div>
	<div data-testid='description'>
		<span class="Formatted">One of the best-loved stories of all time, To Kill a Mockingbird has been translated into more than forty languages, sold more than forty million copies worldwide, served as the basis for an enormously popular motion picture, and was voted one of the best novels of the twentieth century by librarians across the country. A gripping, heart-wrenching, and wholly remarkable coming-of-age tale in a South poisoned by virulent prejudice, it views a world of great beauty and savage iniquities through the eyes of a young girl, as her father — a crusading local lawyer — risks everything to defend a black man unjustly accused of a terrible crime.</span>
	</div>
</div>
`

func TestParseBook(t *testing.T) {

	expectedBook := Book{
		Title:       "To Kill a Mockingbird",
		Cover:       "https://images-na.ssl-images-amazon.com/images/S/compressed.photo.goodreads.com/books/1612238791i/56916837.jpg",
		PublishDate: "1960",
		Rating:      "4.26",
		Author:      "Harper Lee",
		Description: `One of the best-loved stories of all time, To Kill a Mockingbird has been translated into more than forty languages, sold more than forty million copies worldwide, served as the basis for an enormously popular motion picture, and was voted one of the best novels of the twentieth century by librarians across the country. A gripping, heart-wrenching, and wholly remarkable coming-of-age tale in a South poisoned by virulent prejudice, it views a world of great beauty and savage iniquities through the eyes of a young girl, as her father — a crusading local lawyer — risks everything to defend a black man unjustly accused of a terrible crime.`,
	}

	url := "https://www.goodreads.com/book/show/56916837-to-kill-a-mockingbird?ref=nav_sb_ss_1_21"
	currentBook, err := ParseBook(url)

	if err != nil {
		t.Error("Should return nil value")
		t.Error(err)
	}

	if expectedBook.Title != currentBook.Title {
		t.Errorf("returned wrong title: got %v want %v", currentBook.Title, expectedBook.Title)
		t.Error(err)
	} else {
		t.Logf("returned correct title: ✅ ")
	}

	if expectedBook.PublishDate != currentBook.PublishDate {
		t.Errorf("returned wrong publish date: got %v ❌ want %v", currentBook.PublishDate, expectedBook.PublishDate)
		t.Error(err)
	} else {
		t.Logf("returned correct publish date: ✅ ")
	}

	if expectedBook.Description != currentBook.Description {
		t.Errorf("returned wrong description: got %v ❌ want %v", currentBook.Description, expectedBook.Description)
		t.Error(err)
	} else {
		t.Logf("returned correct description: ✅ ")
	}
}
