package gklang

import "log"

// Er exit with error log
func Er(e error) {
	log.Panicln(e)
}
