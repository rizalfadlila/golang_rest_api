package helpers

// Block :nodoc:
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception :nodoc:
type Exception error

// Throw :nodoc:
func Throw(up Exception) {
	if up != nil {
		panic(up)
	}
}

// Do :nodoc:
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r.(error))
			}
		}()
	}
	tcf.Try()
}
