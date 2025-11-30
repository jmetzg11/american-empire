package main

import (
	"encoding/json"
	"os"
	"time"
)

type Event struct {
	ID      int
	Title   string
	Date    string
	Country string
	Tags    []string
}

func (app *application) getMainPage() ([]Event, error) {
	query := `
		SELECT
			e.id,
			e.title,
			e.date,
			e.country,
			COALESCE(json_agg(DISTINCT t.name) FILTER (WHERE t.id IS NOT NULL), '[]') as tags
		FROM events e
		LEFT JOIN event_tags et ON e.id = et.event_id
		LEFT JOIN tags t ON et.tag_id = t.id
		WHERE e.active IS NOT NULL AND e.flagged = False
		GROUP BY e.id, e.title, e.date, e.country
		ORDER BY e.date DESC
	`
	rows, err := app.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var tagsJSON string
		var dateTime time.Time

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&dateTime,
			&event.Country,
			&tagsJSON,
		)
		if err != nil {
			return nil, err
		}

		event.Date = dateTime.Format("2006 Jan 02")

		json.Unmarshal([]byte(tagsJSON), &event.Tags)

		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

type BookEvent struct {
	Id      int
	Title   string
	Date    string
	Country string
}

type BookMain struct {
	Title  string
	Author string
	Link   string
	Events []BookEvent
}

func (app *application) getBooks() ([]BookMain, error) {
	query := `
		SELECT
			b.title,
			b.author,
			b.link,
			COALESCE(json_agg(DISTINCT jsonb_build_object(
				'id', e.id,
				'title', e.title,
				'date', e.date,
				'country', e.country
			)) FILTER (WHERE e.id IS NOT NULL), '[]') as events
		FROM books b
		LEFT JOIN book_events be ON b.id = be.book_id
		LEFT JOIN events e ON be.event_id = e.id
		GROUP BY b.id, b.title, b.author, b.link
		ORDER BY b.title
	`
	rows, err := app.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []BookMain
	for rows.Next() {
		var book BookMain
		var eventsJSON string

		err := rows.Scan(
			&book.Title,
			&book.Author,
			&book.Link,
			&eventsJSON,
		)
		if err != nil {
			return nil, err
		}

		var events []BookEvent
		json.Unmarshal([]byte(eventsJSON), &events)

		for i := range events {
			if events[i].Date != "" {
				dateTime, err := time.Parse("2006-01-02T15:04:05Z", events[i].Date)
				if err == nil {
					events[i].Date = dateTime.Format("2006 Jan 02")
				}
			}
		}
		book.Events = events
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

type Source struct {
	Name string
	URL  string
}

type Media struct {
	Type    string
	URL     string
	Path    string
	Caption string
}

type Book struct {
	Title  string
	Author string
	Link   string
}

type DetailedEvent struct {
	ID          int
	Title       string
	Date        string
	Country     string
	Description string
	Tags        []string
	Sources     []Source
	Medias      []Media
	Books       []Book
}

func (app *application) getEvent(id string) (DetailedEvent, error) {
	query := `
		SELECT
			e.id,
			e.title,
			e.date,
			e.country,
			e.description,
			COALESCE(json_agg(DISTINCT t.name) FILTER (WHERE t.id IS NOT NULL), '[]') as tags,
			COALESCE(json_agg(DISTINCT jsonb_build_object(
				'name', s.name,
				'url', s.url
			)) FILTER (WHERE s.id IS NOT NULL), '[]') as sources,
			COALESCE(json_agg(DISTINCT jsonb_build_object(
				'type', m.type,
				'url', m.url,
				'path', m.path,
				'caption', m.caption
			)) FILTER (WHERE m.id IS NOT NULL), '[]') as medias,
			COALESCE(json_agg(DISTINCT jsonb_build_object(
				'title', b.title,
				'author', b.author,
				'link', b.link
			)) FILTER (WHERE b.id IS NOT NULL), '[]') as books
		FROM events e
		LEFT JOIN event_tags et ON e.id = et.event_id
		LEFT JOIN tags t ON et.tag_id = t.id
		LEFT JOIN sources s ON e.id = s.event_id
		LEFT JOIN media m ON e.id = m.event_id
		LEFT JOIN book_events be ON e.id = be.event_id
		LEFT JOIN books b on be.book_id = b.id
		WHERE e.active IS NOT NULL
		AND e.id = $1
		GROUP BY e.id, e.title, e.date, e.country, e.description
	`

	var event DetailedEvent
	var dateTime time.Time
	var tagsJSON, sourcesJSON, mediasJSON, booksJSON string

	err := app.db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Title,
		&dateTime,
		&event.Country,
		&event.Description,
		&tagsJSON,
		&sourcesJSON,
		&mediasJSON,
		&booksJSON,
	)
	if err != nil {
		return DetailedEvent{}, err
	}

	event.Date = dateTime.Format("2006, Jan 02")
	json.Unmarshal([]byte(tagsJSON), &event.Tags)
	json.Unmarshal([]byte(sourcesJSON), &event.Sources)
	json.Unmarshal([]byte(booksJSON), &event.Books)

	json.Unmarshal([]byte(mediasJSON), &event.Medias)
	for i := range event.Medias {
		if event.Medias[i].Type == "photo" && event.Medias[i].Path != "" {
			event.Medias[i].URL = os.Getenv("PHOTOS_URL") + event.Medias[i].Path
		}
	}

	return event, nil
}
