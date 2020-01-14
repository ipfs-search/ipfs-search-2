package procesor

import "context"
import "errors"

type input interface{}
type output interface{}

// Processor takes a context and input and returns an output, or an error.
type Processor interface {
	Process(context.Context, input) (output, error)
}

// Error for processing.
type Error struct {
	Err   error
	Input input
}

// Unwrap yields the underlying error for a Error.
func (e *Error) Unwrap() error { return e.Err }

func (e *Error) Error() string {
	return fmt.Sprintf("%v while processing %v", e.Err, e.Input)
}

// ProcessChannel loops on a channel of inputs, yielding a channel of outputs and errors.
func ProcessChannel(ctx context.Context, p Processor, inputc <-chan input, outputc chan<- output, errc chan<- error) nil {
	for {

		select {
		case <-ctx.Done():
			// Context closed, return
			return

		case input := <-inputc:
			err := p.Process(ctx, input)

			if err != nil {
				errc <- &Error{
					err,
					input,
				}
			}
		}
	}
}
