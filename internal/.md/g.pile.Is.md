type anyTypeIs func(anyType) bool
    anyTypeIs represents a boolean predicate of anyType implemented as a boolean
    function for anything of type anyType.

func (a anyTypeIs) And(predicates ...func(anyType) bool) anyTypeIs
func (a anyTypeIs) Eval(arg anyType) bool
func (a anyTypeIs) Is() anyTypeIs
func (a anyTypeIs) Not() anyTypeIs
func (a anyTypeIs) Or(predicates ...func(anyType) bool) anyTypeIs
