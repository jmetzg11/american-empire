package database

import (
	"good-guys/backend/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	var count int64
	db.Model(&models.Event{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}
	activeDate := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)

	events := []models.Event{
		{
			Title:   "CIA Operation Mockingbird",
			Date:    time.Date(1950, 3, 15, 0, 0, 0, 0, time.UTC),
			Country: "USA",
			Description: `Operation Mockingbird was a large-scale program of the United States Central Intelligence Agency that began in the early years of the Cold War.

   The operation sought to influence media organizations and journalists to disseminate propaganda favorable to CIA interests. This included recruiting leading American journalists into a network to help present the CIA's views, and funded some student and cultural organizations, and magazines as fronts.

   Key aspects included:
   - Infiltration of major news organizations
   - Creation of propaganda materials
   - Coordination with foreign intelligence services`,
			Active: &activeDate,
			Sources: []models.Source{
				{Name: "Church Committee Report", URL: "https://archive.org/details/churchcommittee"},
				{Name: "Declassified CIA Documents", URL: "https://cia.gov/mockingbird-files"},
			},
			Medias: []models.Media{
				{Type: "photo", Path: "/data/photos/1/sample.jpg"},
				{Type: "youtube", URL: "VPzp51qfqB8"},
			},
		},
		{
			Title:   "Operation Condor South America",
			Date:    time.Date(1975, 11, 28, 0, 0, 0, 0, time.UTC),
			Country: "Chile",
			Description: `Operation Condor was a United States-backed campaign of political repression and state terror involving intelligence operations and assassination.

   The operation was officially implemented in November 1975 by the right-wing dictatorships of the Southern Cone of South America. The program was intended to eradicate communist or Soviet influence and ideas, and to suppress active or potential opposition movements against the participating governments.

   Coordinated efforts included:
   - Cross-border intelligence sharing
   - Joint military operations
   - Systematic human rights violations`,
			Active: &activeDate,
			Sources: []models.Source{
				{Name: "National Security Archive", URL: "https://nsarchive.gwu.edu/condor-files"},
				{Name: "FBI FOIA Release", URL: "https://fbi.gov/condor-documents"},
			},
			Medias: []models.Media{
				{Type: "photo", Path: "/data/photos/2/sample.jpeg"},
				{Type: "youtube", URL: "VPzp51qfqB8"},
			},
		},
		{
			Title:   "Iran-Contra Weapons Sales",
			Date:    time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC),
			Country: "Nicaragua",
			Description: `The Iran-Contra affair was a political scandal that occurred during the second term of the Reagan Administration.

   Senior administration officials secretly facilitated the sale of arms to Iran, which was then under an arms embargo. The officials hoped that the arms sales would secure the release of American hostages and allow U.S. intelligence agencies to fund the Nicaraguan Contras.

   This covert operation involved:
   - Illegal arms sales to Iran
   - Diversion of proceeds to Contra rebels
   - Circumvention of congressional oversight`,
			Active: &activeDate,
			Sources: []models.Source{
				{Name: "Tower Commission Report", URL: "https://reagan.library.gov/tower-commission"},
				{Name: "Walsh Report", URL: "https://justice.gov/walsh-final-report"},
			},
			Medias: []models.Media{
				{Type: "photo", Path: "/data/photos/3/sample.jpg"},
				{Type: "youtube", URL: "VPzp51qfqB8"},
			},
		},
	}

	db.Create(&events)
	log.Println("Database seeded successfully!")
}
