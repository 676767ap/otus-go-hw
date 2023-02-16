package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	go func() {
		defer close(out)

		stream := in
		for _, stage := range stages {
			stream = stage(stream)
		}

		for {
			select {
			case <-done:
				return
			case i, ok := <-stream:
				if !ok {
					return
				}
				out <- i
			}
		}
	}()

	return out
}
