// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package runtime

//如果不是windows平台占位用
//go:nosplit
func osPreemptExtEnter(mp *m) {}

//go:nosplit
func osPreemptExtExit(mp *m) {}