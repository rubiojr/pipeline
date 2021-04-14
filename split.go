package pipeline

// Split takes an interface from Collect and splits it back out into individual elements
// Usefull for batch processing pipelines (`intput chan -> Collect -> Process -> Split -> Cancel -> output chan`).
func Split(in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for is := range in {
			for _, i := range is.([]interface{}) {
				out <- i
			}
		}
	}()
	return out
}
