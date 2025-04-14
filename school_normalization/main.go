package main

import (
//    "database/sql"
    "fmt"
    
//    _ "github.com/lib/pq"
    "norm/school"
)

func main() {
    // Example school names
    examples := []string{
        "St. Mary's International School",
        "IES Ramón y Cajal",
        "P.S. 121 Nelson A. Rockefeller",
        "Massachusetts Institute of Technology",
        "UCLA",
        "École Normale Supérieure de Lyon",
    }
    
    for _, name := range examples {
        normalized := school.NormalizeSchoolName(name)
        fmt.Printf("Original: %s\n", normalized.Original)
        fmt.Printf("Normalized: %s\n", normalized.Normalized)
        fmt.Printf("Abbreviated: %s\n", normalized.Abbreviated)
        fmt.Printf("Search Index: %s\n", normalized.SearchIndex)
        fmt.Println("---")
    }
}

