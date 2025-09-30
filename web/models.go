package main 

import (
	"time"
)

type EventSummary struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Date time.Time `json:"date"`
	Country string `json:"countr"`
	Tags []string `json:"tags"`
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

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Date,
			&event.Country,
			&tagsStr,
		)
		if err != nil {
			return nil, err
		}

		// Split comma-separated tags into slice
		if tagsStr != "" {
			event.Tags = []string{}
			for _, tag := range []string{tagsStr} {
				if tag != "" {
					event.Tags = append(event.Tags, tag)
				}
			}
		} else {
			event.Tags = []string{}
		}

		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
