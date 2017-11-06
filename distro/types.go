//
// Copyright (c) 2017
// Cavium
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"github.com/drasko/edgex-export"
)

// Sender - Send interface
type Sender interface {
	Send(data []byte)
}

// Formater - Format interface
type Formater interface {
	Format(event *export.Event) []byte
}

// Transformer - Transform interface
type Transformer interface {
	Transform(data []byte) []byte
}

// RegistrationInfo - registration info
type RegistrationInfo struct {
	registration export.Registration
	format       Formater
	compression  Transformer
	encrypt      Transformer
	sender       Sender

	chRegistration chan *export.Registration
	chEvent        chan *export.Event

	deleteMe bool
}
