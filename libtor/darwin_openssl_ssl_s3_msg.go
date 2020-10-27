// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build darwin,amd64 darwin,arm64 ios,amd64 ios,arm64
// +build staticOpenssl

package libtor

/*
#ifdef PREFIX_GO_LIBTOR
# pragma extern_prefix GO_LIBTOR_
#endif

#define DSO_NONE
#define OPENSSLDIR "/usr/local/ssl"
#define ENGINESDIR "/usr/local/lib/engines"

#include <../ssl/s3_msg.c>
*/
import "C"
