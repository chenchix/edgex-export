//
// (C) Copyright 2017
// Mainflux
//
// SPDX-License-Identifier:	Apache-2.0
//

package distro

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger(l *zap.Logger) {
	logger = l
	return
}
