package main

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type quoteText struct {
	QuoteText string `json:"quoteText"`
}

func insertQuote(ctx context.Context, conn *pgx.Conn, quote Quote) error {
	sql := `INSERT INTO public.quotes 
			(id, created_at, quote_text, author, tags, likes, quote_url) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := conn.Exec(ctx, sql, quote.Id, quote.CreatedAt, quote.QuoteText, quote.Author, quote.Tags, quote.Likes, quote.QuoteUrl)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func getQuotes(ctx context.Context, conn * pgx.Conn) ([]quoteText, error) {
	sql := `SELECT quote_text
			FROM public.quotes
			LIMIT 10`
	rows, err := conn.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var quotes []quoteText
	for rows.Next() {
		var quote quoteText
		err = rows.Scan(&quote.QuoteText)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil


}
