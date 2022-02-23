
After printing the users information and before make a host search we fetch available filters and facets, printing both lists. This is the relevant code:

``` Go
	// get available filters
	filters, err := s.Filters()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("filters:")
	for _, filter := range filters {
		fmt.Printf("\t%s\n", filter)
	}

	// get available facets
	facets, err := s.Filters()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("facets:")
	for _, facet := range facets {
		fmt.Printf("\t%s\n", facet)
	}
```