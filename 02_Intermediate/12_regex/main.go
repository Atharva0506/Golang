package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 1. Compile a pattern once and reuse it.
	// Use MustCompile when the pattern is a compile-time constant — it panics on invalid patterns
	// instead of silently returning a nil pointer like Compile does.
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)

	// 2. MatchString — does the string contain a match?
	fmt.Println("--- MatchString ---")
	fmt.Println(emailPattern.MatchString("hello@example.com")) // true
	fmt.Println(emailPattern.MatchString("not-an-email"))      // false

	// 3. FindString — return the FIRST match (or empty string if none)
	fmt.Println("\n--- FindString ---")
	text := "Contact us at support@acme.com or sales@acme.com"
	first := emailPattern.FindString(text)
	fmt.Println("First email:", first) // support@acme.com

	// 4. FindAllString — return ALL matches (-1 means "no limit")
	fmt.Println("\n--- FindAllString ---")
	all := emailPattern.FindAllString(text, -1)
	fmt.Println("All emails:", all) // [support@acme.com sales@acme.com]

	// 5. FindStringSubmatch — capture groups
	// Wrap parts of a pattern in () to create a "capture group"
	fmt.Println("\n--- FindStringSubmatch (capture groups) ---")
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	match := datePattern.FindStringSubmatch("Today is 2024-07-15 and tomorrow is 2024-07-16")
	if len(match) > 0 {
		fmt.Println("Full match:", match[0]) // 2024-07-15
		fmt.Println("Year:      ", match[1]) // 2024
		fmt.Println("Month:     ", match[2]) // 07
		fmt.Println("Day:       ", match[3]) // 15
	}

	// 6. ReplaceAllString — find-and-replace
	fmt.Println("\n--- ReplaceAllString ---")
	redacted := emailPattern.ReplaceAllString(text, "[REDACTED]")
	fmt.Println(redacted) // Contact us at [REDACTED] or [REDACTED]

	// 7. ReplaceAllStringFunc — replace using a function
	fmt.Println("\n--- ReplaceAllStringFunc ---")
	masked := emailPattern.ReplaceAllStringFunc(text, func(email string) string {
		// Keep only the domain part for logging
		domainPattern := regexp.MustCompile(`@.+`)
		return "***" + domainPattern.FindString(email)
	})
	fmt.Println(masked) // Contact us at ***@acme.com or ***@acme.com

	// 8. Split — split a string by a regex pattern
	fmt.Println("\n--- Split ---")
	csvLine := "apple , banana,  cherry ,date"
	splitPattern := regexp.MustCompile(`\s*,\s*`)
	parts := splitPattern.Split(csvLine, -1)
	fmt.Println(parts) // [apple banana cherry date]
}
