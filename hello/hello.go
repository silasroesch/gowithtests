package main

const spanish = "spanish"
const french = "french"
const german = "german"
const prefixEnglish = "Hello, "
const prefixGerman = "Hallo, "
const prefixFrench = "Bonjour, "
const prefixSpanish = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case german:
		prefix = prefixGerman
	case french:
		prefix = prefixFrench
	case spanish:
		prefix = prefixSpanish
	default:
		prefix = prefixEnglish
	}
	return
}

func main() {
}
