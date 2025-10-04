package main

import (
	"fmt"
	"strings"
	"time"
)

type EventSummary struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Date    string   `json:"date"`
	Country string   `json:"countr"`
	Tags    []string `json:"tags"`
}

func (app *application) getMainPage() ([]EventSummary, error) {
	query := `
		SELECT
			e.id,
			e.title,
			e.date,
			e.country,
			COALESCE(string_agg(t.name, ','), '') as tags
		FROM events e
		LEFT JOIN event_tags et ON e.id = et.event_id
		LEFT JOIN tags t ON et.tag_id = t.id
		WHERE e.active IS NOT NULL
		GROUP BY e.id, e.title, e.date, e.country
		ORDER BY e.date DESC
	`
	rows, err := app.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []EventSummary
	for rows.Next() {
		var event EventSummary
		var tagsStr string
		var dateTime time.Time

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&dateTime,
			&event.Country,
			&tagsStr,
		)
		if err != nil {
			return nil, err
		}

		event.Date = dateTime.Format("2006 Jan 02")

		if tagsStr != "" {
			event.Tags = strings.Split(tagsStr, ",")
		} else {
			event.Tags = []string{}
		}

		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", events)

	return events, nil
}
