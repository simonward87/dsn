// Copyright (C) 2025 Simon Ward <39803787+simonward87@users.noreply.github.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dsn_test

import (
	"testing"

	"dsn"
)

func TestSQLite(t *testing.T) {
	testCases := []struct {
		expect string
		file   string
		name   string
		opts   dsn.Options
	}{
		{
			expect: "file::memory:?_busy_timeout=1000&_foreign_keys=true",
			file:   ":memory:",
			name:   "Basic",
			opts: dsn.Options{
				BusyTimeout:     1_000,
				ForeignKeys:     true,
				SharedCacheMode: dsn.SharedCacheModeShared,
			},
		},
		{
			expect: "file:api.sqlite3?_auth&_auth_crypt=SSHA256&_auth_pass=ryxdd69aP9gy&_auth_salt=TpSI7DCaOjyfMC2GL0VLmYSW&_auth_user=sqlite_user_123&_busy_timeout=10000&_foreign_keys=true&_journal_mode=WAL",
			file:   "api.sqlite3",
			name:   "WithAuth",
			opts: dsn.Options{
				BusyTimeout: 10_000,
				ForeignKeys: true,
				JournalMode: dsn.JournalModeWal,
				UACrypt:     dsn.UACryptSSHA256,
				UAPassword:  "ryxdd69aP9gy",
				UASalt:      "TpSI7DCaOjyfMC2GL0VLmYSW",
				UAUsername:  "sqlite_user_123",
			},
		},
		{
			expect: "",
			file:   "",
			name:   "NoFile",
			opts: dsn.Options{
				BusyTimeout:     1_000,
				ForeignKeys:     true,
				SharedCacheMode: dsn.SharedCacheModeShared,
			},
		},
		{
			expect: "",
			file:   "",
			name:   "NoFileNoOpts",
			opts:   dsn.Options{},
		},
		{
			expect: "file:api.sqlite3",
			file:   "api.sqlite3",
			name:   "NoOpts",
			opts:   dsn.Options{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			connStr := dsn.SQLite(tc.file, tc.opts)
			if tc.expect != connStr {
				t.Errorf("want %q, got %q", tc.expect, connStr)
			}
		})
	}
}
