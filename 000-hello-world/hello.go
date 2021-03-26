package hello

const english = "English"
const german = "German"
const french = "French"

const englishHelloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case german:
		prefix = germanHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
