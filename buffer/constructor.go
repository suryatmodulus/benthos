/*
Copyright (c) 2014 Ashley Jeffs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package buffer

import (
	"github.com/jeffail/benthos/buffer/ring"
	"github.com/jeffail/benthos/types"
	"github.com/jeffail/util/log"
	"github.com/jeffail/util/metrics"
)

//--------------------------------------------------------------------------------------------------

var constructors = map[string]func(conf Config, log log.Modular, stats metrics.Aggregator) (Type, error){}

//--------------------------------------------------------------------------------------------------

// Config - The all encompassing configuration struct for all input types.
type Config struct {
	Type   string            `json:"type" yaml:"type"`
	Mmap   ring.MmapConfig   `json:"mmap_file" yaml:"mmap_file"`
	Memory ring.MemoryConfig `json:"memory" yaml:"memory"`
}

// NewConfig - Returns a configuration struct fully populated with default values.
func NewConfig() Config {
	return Config{
		Type:   "none",
		Mmap:   ring.NewMmapConfig(),
		Memory: ring.NewMemoryConfig(),
	}
}

//--------------------------------------------------------------------------------------------------

// Construct - Create an input type based on an input configuration.
func Construct(conf Config, log log.Modular, stats metrics.Aggregator) (Type, error) {
	if c, ok := constructors[conf.Type]; ok {
		return c(conf, log, stats)
	}
	return nil, types.ErrInvalidBufferType
}

//--------------------------------------------------------------------------------------------------