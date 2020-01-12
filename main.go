package main

// NOTE: It seems that pkger has problems to create a bundle
// if it does not see a .go file on modules root.
// This main.go file is not really needed and will be deleted
// after finding a solution to this issue.

import (
	_ "golang.org/x/text/language"
	_ "golang.org/x/text/message"
	_ "golang.org/x/text/message/catalog"
)

//go:generate gotext -srclang=en update -out=pkg/auth/web/i18n.go -lang=en,es,de,pl

func main() {
	return
}
