// Copyright 2016 The go-irchain Authors
// This file is part of the go-irchain library.
//
// The go-irchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-irchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-irchain library. If not, see <http://www.gnu.org/licenses/>.

// Contains initialization code for the mbile library.

package girc

import (
	"os"
	"runtime"

	"github.com/irchain/go-irchain/log"
)

func init() {
	// Initialize the logger
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))

	// Initialize the goroutine count
	runtime.GOMAXPROCS(runtime.NumCPU())
}
