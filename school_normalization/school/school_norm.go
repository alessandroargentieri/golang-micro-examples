package school

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SchoolName holds both the original and normalized versions of a school name
type SchoolName struct {
	Original    string `json:"original" db:"original_name"`
	Normalized  string `json:"normalized" db:"normalized_name"`
	Abbreviated string `json:"abbreviated,omitempty" db:"abbreviated_name"`
	SearchIndex string `json:"search_index" db:"search_index"` // For full-text search
}

var (
	// Regex patterns
	multipleSpacesRegex = regexp.MustCompile(`\s+`)
	nonAlphaNumRegex    = regexp.MustCompile(`[^\p{L}\p{N}\s]`)
	numbersRegex        = regexp.MustCompile(`\d+`)

	// Common words to remove
	commonWords = map[string]bool{
		"the": true, "and": true, "of": true, "for": true, "in": true, "at": true, "by": true,
	}

	// Educational institution type identifiers
	schoolTypes = map[string]bool{
		"school":        true,
		"schools":       true,
		"college":       true,
		"university":    true,
		"institute":     true,
		"academy":       true,
		"elementary":    true,
		"primary":       true,
		"secondary":     true,
		"high":          true,
		"middle":        true,
		"junior":        true,
		"senior":        true,
		"international": true,
		"grammar":       true,
		"prep":          true,
		"preparatory":   true,
		"higher":        true,
		"lower":         true,
	}

	// Common abbreviations and their expansions
	abbreviations = map[string]string{
		"st":     "saint",
		"st.":    "saint",
		"sch":    "school",
		"tech":   "technology",
		"tech.":  "technology",
		"univ":   "university",
		"univ.":  "university",
		"intl":   "international",
		"intl.":  "international",
		"elem":   "elementary",
		"elem.":  "elementary",
		"acad":   "academy",
		"acad.":  "academy",
		"coll":   "college",
		"coll.":  "college",
		"inst":   "institute",
		"inst.":  "institute",
		"dept":   "department",
		"dept.":  "department",
		"natl":   "national",
		"natl.":  "national",
		"sci":    "science",
		"sci.":   "science",
		"eng":    "engineering",
		"eng.":   "engineering",
		"med":    "medical",
		"med.":   "medical",
		"jhs":    "junior high school",
		"shs":    "senior high school",
		"ms":     "middle school",
		"ps":     "public school",
		"p.s.":   "public school",
		"hs":     "high school",
		"h.s.":   "high school",
		"ies":    "instituto de educacion secundaria", // Spanish
		"ceip":   "centro de educacion infantil y primaria", // Spanish
		"ib":     "international baccalaureate",
		"k-12":   "kindergarten through twelfth grade",
		"k12":    "kindergarten through twelfth grade",
	}
)

// NormalizeSchoolName creates a SchoolName object with both original and normalized forms
func NormalizeSchoolName(name string) SchoolName {
	if name == "" {
		return SchoolName{}
	}

	// Create result struct with original name
	result := SchoolName{
		Original: name,
	}

	// Normalize the name
	normalized := name

	// Convert to lowercase
	normalized = strings.ToLower(normalized)

	// Remove accents and diacritics
	normalized = removeAccents(normalized)

	// Replace common abbreviations with full forms (for search index)
	expandedName := expandAbbreviations(normalized)

	// Remove non-alphanumeric characters
	normalized = nonAlphaNumRegex.ReplaceAllString(normalized, " ")

	// Replace multiple spaces with a single space
	normalized = multipleSpacesRegex.ReplaceAllString(normalized, " ")

	// Trim spaces
	normalized = strings.TrimSpace(normalized)

	// Create abbreviation by taking first letter of each word
	words := strings.Fields(normalized)
	var abbr strings.Builder
	for _, word := range words {
		if !commonWords[word] && len(word) > 0 {
			abbr.WriteRune(rune(word[0]))
		}
	}
	result.Abbreviated = strings.ToUpper(abbr.String())

	// Store normalized name
	result.Normalized = normalized

	// Create optimized search index (expanded abbreviations, without common words)
	result.SearchIndex = createSearchIndex(expandedName)

	return result
}

// removeAccents removes diacritical marks
func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// expandAbbreviations replaces common abbreviations with their full forms
func expandAbbreviations(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if expansion, ok := abbreviations[word]; ok {
			words[i] = expansion
		}
	}
	return strings.Join(words, " ")
}

// createSearchIndex removes common words and school type identifiers
// for a more focused search index
func createSearchIndex(s string) string {
	words := strings.Fields(s)
	var result []string

	for _, word := range words {
		// Skip common words and school type identifiers unless they're part of a larger token
		if !commonWords[word] && !schoolTypes[word] {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

// ExtractInitialism extracts an acronym/initialism from a school name
// (useful for finding schools commonly referred to by their initials)
func ExtractInitialism(name string) string {
	// Handle cases where school is already written as initials (like "MIT" or "UCLA")
	if isLikelyInitialism(name) {
		return strings.ToUpper(strings.ReplaceAll(name, ".", ""))
	}

	words := strings.Fields(strings.ToLower(name))
	var initials strings.Builder

	// Extract first letter of significant words
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 0 && !commonWords[word] {
			initials.WriteRune(rune(word[0]))
		}
	}

	return strings.ToUpper(initials.String())
}

// isLikelyInitialism checks if the name is already in initialism form
func isLikelyInitialism(name string) bool {
	// Remove periods and spaces
	cleaned := strings.ReplaceAll(name, ".", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	
	// If it's 2-6 characters and all uppercase, it's likely an initialism
	return len(cleaned) >= 2 && len(cleaned) <= 6 && cleaned == strings.ToUpper(cleaned)
}

