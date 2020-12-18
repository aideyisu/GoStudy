module library

go 1.15

replace greetings => ../internal/greetings

replace booksystem => ../internal/booksystem

require (
	booksystem v0.0.0-00010101000000-000000000000
	greetings v0.0.0-00010101000000-000000000000
)
