package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SchoolTypeInfo holds information about a school type
type SchoolTypeInfo struct {
	Type    string // The full lowercase name of the school type
	Acronym string // The acronym for the school type
}

// NormalizeSchoolName normalizes a school name and identifies its type
// Returns the normalized name and the school type information
func NormalizeSchoolName(schoolName string) (string, SchoolTypeInfo) {
	// Define school type mappings - all keys in lowercase
	schoolTypeMap := map[string]SchoolTypeInfo{
		// Istituto Comprensivo variations
		"i. c.":                {Type: "istituto comprensivo", Acronym: "ic"},
		"i.c.":                 {Type: "istituto comprensivo", Acronym: "ic"},
		"ic":                   {Type: "istituto comprensivo", Acronym: "ic"},
		"i.c":                  {Type: "istituto comprensivo", Acronym: "ic"},
		"ic.":                  {Type: "istituto comprensivo", Acronym: "ic"},
		"i. comprensivo":       {Type: "istituto comprensivo", Acronym: "ic"},
		"i c":                  {Type: "istituto comprensivo", Acronym: "ic"},
		"ist. comprensivo":     {Type: "istituto comprensivo", Acronym: "ic"},
		"ist. compr.":          {Type: "istituto comprensivo", Acronym: "ic"},
		"ist.comprensivo":      {Type: "istituto comprensivo", Acronym: "ic"},
		"ist.compr.":           {Type: "istituto comprensivo", Acronym: "ic"},
		"ist comprensivo":      {Type: "istituto comprensivo", Acronym: "ic"},
		"istituto comprensivo": {Type: "istituto comprensivo", Acronym: "ic"},
		"istituto compr.":      {Type: "istituto comprensivo", Acronym: "ic"},

		// Circolo Didattico variations
		"c. d.": {Type: "circolo didattico", Acronym: "cd"},
		"c.d.":  {Type: "circolo didattico", Acronym: "cd"},
		"cd":    {Type: "circolo didattico", Acronym: "cd"},
		"d. d.": {Type: "direzione didattica", Acronym: "dd"},
		"d.d.":  {Type: "direzione didattica", Acronym: "dd"},
		"dd":    {Type: "direzione didattica", Acronym: "dd"},

		// Istituto Istruzione Superiore variations
		"i.i.s.":   {Type: "istituto istruzione superiore", Acronym: "iis"},
		"i.i.s":    {Type: "istituto istruzione superiore", Acronym: "iis"},
		"iis":      {Type: "istituto istruzione superiore", Acronym: "iis"},
		"i.i.s.s.": {Type: "istituto istruzione superiore statale", Acronym: "iiss"},
		"i.i.s.s":  {Type: "istituto istruzione superiore statale", Acronym: "iiss"},
		"iiss":     {Type: "istituto istruzione superiore statale", Acronym: "iiss"},
		"i.is":     {Type: "istituto istruzione superiore", Acronym: "iis"},
		"iss":      {Type: "istituto superiore statale", Acronym: "iss"},
		"isis":     {Type: "istituto statale istruzione superiore", Acronym: "isis"},
		"isiss":    {Type: "istituto statale istruzione superiore statale", Acronym: "isiss"},
		"iisss":    {Type: "istituto istruzione secondaria superiore statale", Acronym: "iisss"},

		// Istituto Omnicomprensivo variations
		"i. omnicompr":             {Type: "istituto omnicomprensivo", Acronym: "io"},
		"i. omnicomprensivo":       {Type: "istituto omnicomprensivo", Acronym: "io"},
		"i.omnicomprensivo":        {Type: "istituto omnicomprensivo", Acronym: "io"},
		"i.omnicompr":              {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist. omnicomprensivo":     {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist. omnicompr":           {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist. onnicomprensivo":     {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist.omnicomprensivo":      {Type: "istituto omnicomprensivo", Acronym: "io"},
		"istituto omnicomprensivo": {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist omni":                 {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist omnicom":              {Type: "istituto omnicomprensivo", Acronym: "io"},
		"ist omnicomprensivo":      {Type: "istituto omnicomprensivo", Acronym: "io"},

		// Liceo variations
		"liceo":                   {Type: "liceo", Acronym: "l"},
		"liceo classico":          {Type: "liceo classico", Acronym: "lc"},
		"l.classico":              {Type: "liceo classico", Acronym: "lc"},
		"l. classico":             {Type: "liceo classico", Acronym: "lc"},
		"lc":                      {Type: "liceo classico", Acronym: "lc"},
		"liceo scientifico":       {Type: "liceo scientifico", Acronym: "ls"},
		"l.scientifico":           {Type: "liceo scientifico", Acronym: "ls"},
		"l. scientifico":          {Type: "liceo scientifico", Acronym: "ls"},
		"ls":                      {Type: "liceo scientifico", Acronym: "ls"},
		"liceo artistico":         {Type: "liceo artistico", Acronym: "la"},
		"l. artistico":            {Type: "liceo artistico", Acronym: "la"},
		"liceo statale":           {Type: "liceo statale", Acronym: "ls"},
		"l.statale":               {Type: "liceo statale", Acronym: "ls"},
		"liceo linguistico":       {Type: "liceo linguistico", Acronym: "ll"},
		"liceo scienze":           {Type: "liceo scienze umane", Acronym: "lsu"},
		"liceo musicale":          {Type: "liceo musicale", Acronym: "lm"},
		"liceo ginnasio":          {Type: "liceo classico ginnasio", Acronym: "lcg"},
		"liceo pluricomprensivo":  {Type: "liceo pluricomprensivo", Acronym: "lp"},
		"liceo artisticomusicale": {Type: "liceo artistico musicale", Acronym: "lam"},
		"liceo e":                 {Type: "liceo e istituto tecnico", Acronym: "lit"},
		"licei":                   {Type: "licei", Acronym: "l"},

		// Istituto Tecnico variations
		"istituto tecnico": {Type: "istituto tecnico", Acronym: "it"},
		"ist. tecnico":     {Type: "istituto tecnico", Acronym: "it"},
		"iti":              {Type: "istituto tecnico industriale", Acronym: "iti"},
		"itis":             {Type: "istituto tecnico industriale statale", Acronym: "itis"},
		"itc":              {Type: "istituto tecnico commerciale", Acronym: "itc"},
		"itcg":             {Type: "istituto tecnico commerciale e geometri", Acronym: "itcg"},
		"itg":              {Type: "istituto tecnico per geometri", Acronym: "itg"},
		"itet":             {Type: "istituto tecnico economico tecnologico", Acronym: "itet"},
		"ite":              {Type: "istituto tecnico economico", Acronym: "ite"},
		"itt":              {Type: "istituto tecnico tecnologico", Acronym: "itt"},
		"itts":             {Type: "istituto tecnico tecnologico statale", Acronym: "itts"},
		"itst":             {Type: "istituto tecnico settore tecnologico", Acronym: "itst"},
		"its":              {Type: "istituto tecnico superiore", Acronym: "its"},
		"itaer":            {Type: "istituto tecnico aeronautico", Acronym: "itaer"},
		"ita":              {Type: "istituto tecnico agrario", Acronym: "ita"},
		"itsct":            {Type: "istituto tecnico settore commerciale tecnologico", Acronym: "itsct"},
		"ittur":            {Type: "istituto tecnico per il turismo", Acronym: "ittur"},

		// Istituto Professionale variations
		"istituto professionale": {Type: "istituto professionale", Acronym: "ip"},
		"ist. professionale":     {Type: "istituto professionale", Acronym: "ip"},
		"ist prof":               {Type: "istituto professionale", Acronym: "ip"},
		"ist.prof":               {Type: "istituto professionale", Acronym: "ip"},
		"ipsia":                  {Type: "istituto professionale statale industria artigianato", Acronym: "ipsia"},
		"ipsar":                  {Type: "istituto professionale statale alberghiero ristorazione", Acronym: "ipsar"},
		"ipssar":                 {Type: "istituto professionale statale servizi alberghieri ristorazione", Acronym: "ipssar"},
		"ipsseoa":                {Type: "istituto professionale statale servizi enogastronomia ospitalità alberghiera", Acronym: "ipsseoa"},
		"ipseoa":                 {Type: "istituto professionale servizi enogastronomia ospitalità alberghiera", Acronym: "ipseoa"},
		"ipsct":                  {Type: "istituto professionale statale commercio turismo", Acronym: "ipsct"},
		"ipss":                   {Type: "istituto professionale statale servizi", Acronym: "ipss"},
		"ipssa":                  {Type: "istituto professionale statale servizi alberghieri", Acronym: "ipssa"},
		"ipsceoa":                {Type: "istituto professionale statale commercio enogastronomia ospitalità alberghiera", Acronym: "ipsceoa"},
		"ips":                    {Type: "istituto professionale statale", Acronym: "ips"},
		"ipia":                   {Type: "istituto professionale industria artigianato", Acronym: "ipia"},
		"ipalb":                  {Type: "istituto professionale alberghiero", Acronym: "ipalb"},
		"ipsas":                  {Type: "istituto professionale statale arte servizi", Acronym: "ipsas"},
		"ipsasr":                 {Type: "istituto professionale statale agricoltura sviluppo rurale", Acronym: "ipsasr"},
		"ipa":                    {Type: "istituto professionale agrario", Acronym: "ipa"},
		"i p":                    {Type: "istituto professionale", Acronym: "ip"},

		// Other institution types
		"convitto nazionale":  {Type: "convitto nazionale", Acronym: "cn"},
		"conv. naz.":          {Type: "convitto nazionale", Acronym: "cn"},
		"educandato":          {Type: "educandato", Acronym: "ed"},
		"scuola media":        {Type: "scuola media", Acronym: "sm"},
		"scuola secondaria":   {Type: "scuola secondaria", Acronym: "ss"},
		"scuola primaria":     {Type: "scuola primaria", Acronym: "sp"},
		"sms":                 {Type: "scuola media statale", Acronym: "sms"},
		"sm":                  {Type: "scuola media", Acronym: "sm"},
		"sspg":                {Type: "scuola secondaria primo grado", Acronym: "sspg"},
		"isc":                 {Type: "istituto scolastico comprensivo", Acronym: "isc"},
		"ics":                 {Type: "istituto comprensivo statale", Acronym: "ics"},
		"is":                  {Type: "istituto superiore", Acronym: "is"},
		"isi":                 {Type: "istituto statale istruzione", Acronym: "isi"},
		"c.p.i.a.":            {Type: "centro provinciale istruzione adulti", Acronym: "cpia"},
		"istituto magistrale": {Type: "istituto magistrale", Acronym: "im"},
		"ist. magistrale":     {Type: "istituto magistrale", Acronym: "im"},
		"scuola europea":      {Type: "scuola europea", Acronym: "se"},
		"istituto agrario":    {Type: "istituto agrario", Acronym: "ia"},
		"i. s":                {Type: "istituto superiore", Acronym: "is"},
		"i s":                 {Type: "istituto superiore", Acronym: "is"},
		"i.s":                 {Type: "istituto superiore", Acronym: "is"},
		"istituto":            {Type: "istituto", Acronym: "i"},
		"istituto superiore":  {Type: "istituto superiore", Acronym: "is"},
		"ist. superiore":      {Type: "istituto superiore", Acronym: "is"},
		"istituto istruzione": {Type: "istituto istruzione", Acronym: "ii"},
		"ist. istruzione":     {Type: "istituto istruzione", Acronym: "ii"},
		"ist.istruzione":      {Type: "istituto istruzione", Acronym: "ii"},
		"istituto di":         {Type: "istituto di istruzione", Acronym: "idi"},
		"ist. di":             {Type: "istituto di istruzione", Acronym: "idi"},
		"istituto istr":       {Type: "istituto istruzione", Acronym: "ii"},
		"i.t":                 {Type: "istituto tecnico", Acronym: "it"},
		"istruzione":          {Type: "istituto istruzione", Acronym: "ii"},
		"statale":             {Type: "istituto statale", Acronym: "is"},
		"scuole":              {Type: "scuole", Acronym: "s"},
		"scuol":               {Type: "scuola", Acronym: "s"},
		"i circolo":           {Type: "primo circolo", Acronym: "pc"},
		"ii":                  {Type: "secondo istituto comprensivo", Acronym: "sic"},
		"iii":                 {Type: "terzo istituto comprensivo", Acronym: "tic"},
		"iv":                  {Type: "quarto istituto comprensivo", Acronym: "qic"},
		"ix":                  {Type: "nono istituto comprensivo", Acronym: "nic"},
		"i ic":                {Type: "primo istituto comprensivo", Acronym: "pic"},
		"secondo":             {Type: "secondo istituto comprensivo", Acronym: "sic"},
		"settore":             {Type: "settore tecnologico", Acronym: "st"},
	}

	if schoolName == "" {
		return "", SchoolTypeInfo{"unknown", "u"}
	}

	// Convert the input to lowercase first
	lowercaseSchoolName := strings.ToLower(schoolName)

	// Identify school type
	var schoolType SchoolTypeInfo
	var nameWithoutType string

	// Default if no match is found
	schoolType = SchoolTypeInfo{"unknown", "u"}
	nameWithoutType = schoolName

	// Find the longest matching pattern at the beginning of the name
	longestMatch := 0

	for pattern, typeInfo := range schoolTypeMap {
		// Create regex to match at beginning of string
		re := regexp.MustCompile("^" + strings.ReplaceAll(regexp.QuoteMeta(pattern), " ", "\\s+") + "\\s+")
		if re.MatchString(lowercaseSchoolName) {
			loc := re.FindStringIndex(lowercaseSchoolName)
			if loc != nil && loc[1] > longestMatch {
				longestMatch = loc[1]
				schoolType = typeInfo
				// Use original case for the name without type
				nameWithoutType = schoolName[loc[1]:]
			}
		}
	}

	// Normalize the name
	normalized := strings.ToLower(nameWithoutType)

	// Remove quotes
	normalized = strings.ReplaceAll(normalized, "'", "")
	normalized = strings.ReplaceAll(normalized, "-", "_")
	normalized = strings.ReplaceAll(normalized, "\"", "")

	// Replace special characters
	re := regexp.MustCompile("[\\.,;:\\/\\\\()\\[\\]{}#%&]")
	normalized = re.ReplaceAllString(normalized, "")

	// Replace spaces with underscores
	normalized = strings.ReplaceAll(normalized, " ", "_")

	// Remove leading/trailing underscores
	normalized = strings.Trim(normalized, "_")

	// Remove diacritics
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, _ = transform.String(t, normalized)

	// Replace multiple underscores with single underscore
	re = regexp.MustCompile("_+")
	normalized = re.ReplaceAllString(normalized, "_")

	// Combine acronym and normalized name
	if schoolType.Acronym != "u" {
		normalized = schoolType.Acronym + "_" + normalized
	}

	return normalized, schoolType
}

// Example usage in main
func main() {
	// Test with some examples
	examples := []string{
		"I.C. LANCIANO \"DON L.MILANI\"",
		"LICEO SCIENTIFICO \"G.GALILEI\" PESCARA",
		"I.OMNICOMPR.\"CIAMPOLI-SPAVENTA\"",
		"ISTITUTO COMPRENSIVO D.ALIGHIE",
		"I.I.S. \"DE TITTA - FERMI\" - LANCIANO",
		"CONVITTO NAZIONALE \"M.DELFICO\"",
		"C.P.I.A. PESCARA - CHIETI",
		"IPSIA \"DI MARZIO-MICHETTI\" PESCARA",
		"ISTITUTO TECNICO ECONOMICO FEDERICO II",
		"LICEO SCIENTIFICO M.Dell'Atti",
		"Marzio-Michetti",
	}

	for _, example := range examples {
		normalized, typeInfo := NormalizeSchoolName(example)

		println("Original:", example)
		println("Normalized:", normalized)
		println("Type:", typeInfo.Type)
		println("Acronym:", typeInfo.Acronym)
		println("---")
	}
}
