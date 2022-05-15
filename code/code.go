package code

type three_address_code struct {
	counter            int
	original_code      []string
	three_address_code []string
	individual_counter int
	iterator_counter   int
	goto_counter       int
}

func NewThreeAddressCode(original_code []string) *three_address_code {
	return &three_address_code{
		counter:            0,
		original_code:      original_code,
		three_address_code: make([]string, 0),
		individual_counter: 0,
		iterator_counter:   0,
		goto_counter:       0,
	}
}
