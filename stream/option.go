package stream

type Options struct {
	workSize int
}

type Option func(options *Options)
