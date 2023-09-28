package patterns

/*

   [foo] ->
                   foobar ->
   [bar] ->

*/

func Merge(foobar ...<-chan string) <-chan string {
	out := make(chan string)
	// from the input channels
	for _, channel := range foobar {
		go func(in <-chan string) {
			for {
				out <- <-in
			}
		}(channel)
	}
	return out
}
