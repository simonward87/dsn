// Copyright (C) 2025 Simon Ward <39803787+simonward87@users.noreply.github.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dsn

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
)

// SQLite generates a valid dsn string for mattn/go-sqlite3 driver, using a
// file name and an Options object.
func SQLite(file string, opts Options) string {
	if file == "" {
		return ""
	}
	if opts == (Options{}) {
		return fmt.Sprint("file:", file)
	}
	o, err := opts.parse()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("file:%s?%s", file, o)
}

// Options represents all currently available options when creating a
// connection string for mattn/go-sqlite3.
type Options struct {
	AutoVacuum             AutoVaccuum      `url:"_auto_vacuum,omitempty"`
	BusyTimeout            int              `url:"_busy_timeout,omitempty"`
	CacheSize              int              `url:"_cache_size,omitempty"`
	CaseSensitiveLike      bool             `url:"_case_sensitive_like,omitempty"`
	DeferForeignKeys       bool             `url:"_defer_foreign_keys,omitempty"`
	ForeignKeys            bool             `url:"_foreign_keys,omitempty"`
	IgnoreCheckConstraints bool             `url:"_ignore_check_constraints,omitempty"`
	Immutable              bool             `url:"immutable,omitempty"`
	JournalMode            JournalMode      `url:"_journal_mode,omitempty"`
	LockingMode            LockingMode      `url:"_locking_mode,omitempty"`
	Mode                   Mode             `url:"mode,omitempty"`
	MutexLocking           MutexLocking     `url:"_mutex,omitempty"`
	QueryOnly              bool             `url:"_query_only,omitempty"`
	RecursiveTriggers      bool             `url:"_recursive_triggers,omitempty"`
	SecureDelete           any              `url:"_secure_delete,omitempty"`
	SharedCacheMode        SharedCacheMode  `url:"cache,omitempty"`
	Synchronous            Synchronous      `url:"_synchronous,omitempty"`
	TimeZoneLocation       TimeZoneLocation `url:"_loc,omitempty"`
	TransactionLock        TransactionLock  `url:"_txlock,omitempty"`
	UACrypt                UACrypt          `url:"_auth_crypt,omitempty"`
	UAPassword             string           `url:"_auth_pass,omitempty"`
	UASalt                 string           `url:"_auth_salt,omitempty"`
	UAUsername             string           `url:"_auth_user,omitempty"`
	WritableSchema         bool             `url:"_writable_schema,omitempty"`
}

func (o Options) parse() (string, error) {
	var dsn strings.Builder

	switch {
	case o.UAPassword != "",
		o.UASalt != "",
		o.UAUsername != "":

		dsn.WriteString("_auth&")
	}
	v, err := query.Values(o)
	if err != nil {
		return "", err
	}
	s := v.Encode()
	if s == "" {
		return s, nil
	}
	dsn.WriteString(s)
	return dsn.String(), nil
}
