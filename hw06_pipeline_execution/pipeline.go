package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(inCh In) (out Out)

func ExecutePipeline(inCh In, doneCh In, stages ...Stage) Out {
	outCh := inCh

	for _, stage := range stages {
		stageCh := make(Bi)

		go func(inCh Bi, outCh Out) {
			defer close(inCh)

			for {
				select {
				case <-doneCh:
					return
				case result, ok := <-outCh:
					if !ok {
						return
					}
					inCh <- result
				}
			}
		}(stageCh, outCh)

		outCh = stage(stageCh)
	}

	return outCh
}
