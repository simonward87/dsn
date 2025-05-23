// package dsn contains convenience functions for generating connection strings
// for github.com/mattn/go-sqlite3
//
// Copyright (C) 2025 Simon Ward <39803787+simonward87@users.noreply.github.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dsn

// AutoVaccuum represents all possible values for PRAGMA auto_vaccuum
type AutoVaccuum int

const (
	AutoVaccuumNone AutoVaccuum = iota
	AutoVaccuumFull
	AutoVaccuumIncremental
)

func (av AutoVaccuum) String() string {
	switch av {
	case AutoVaccuumNone:
		return "none"
	case AutoVaccuumFull:
		return "full"
	case AutoVaccuumIncremental:
		return "incremental"
	default:
		return ""
	}
}

// JournalMode represents all possible values for PRAGMA journal_mode
type JournalMode int

const (
	JournalModeDelete JournalMode = iota
	JournalModeTruncate
	JournalModePersist
	JournalModeMemory
	JournalModeWal
	JournalModeOff
)

func (jm JournalMode) String() string {
	switch jm {
	case JournalModeDelete:
		return "DELETE"
	case JournalModeTruncate:
		return "TRUNCATE"
	case JournalModePersist:
		return "PERSIST"
	case JournalModeMemory:
		return "MEMORY"
	case JournalModeWal:
		return "WAL"
	case JournalModeOff:
		return "OFF"
	default:
		return ""
	}
}

// LockingMode represents all possible values for PRAGMA locking_mode
type LockingMode int

const (
	LockingModeNormal LockingMode = iota
	LockingModeExclusive
)

func (lm LockingMode) String() string {
	switch lm {
	case LockingModeNormal:
		return "NORMAL"
	case LockingModeExclusive:
		return "EXCLUSIVE"
	default:
		return ""
	}
}

// Mode represents all possible values for SQLite open access mode
type Mode int

const (
	ModeRO Mode = iota
	ModeRW
	ModeRWC
	ModeMemory
)

func (m Mode) String() string {
	switch m {
	case ModeRO:
		return "ro"
	case ModeRW:
		return "rw"
	case ModeRWC:
		return "rwc"
	case ModeMemory:
		return "memory"
	default:
		return ""
	}
}

// MutexLocking represents all possible values for mutex mode
type MutexLocking int

const (
	MutexLockingNo MutexLocking = iota
	MutexLockingFull
)

func (ml MutexLocking) String() string {
	switch ml {
	case MutexLockingNo:
		return "no"
	case MutexLockingFull:
		return "full"
	default:
		return ""
	}
}

// SharedCacheMode represents all possible values for SQLite shared cache mode
type SharedCacheMode int

const (
	SharedCacheModeShared SharedCacheMode = iota
	SharedCacheModePrivate
)

func (scm SharedCacheMode) String() string {
	switch scm {
	case SharedCacheModeShared:
		return "shared"
	case SharedCacheModePrivate:
		return "private"
	default:
		return ""
	}
}

// Synchronous represents all possible values for PRAGMA synchronous
type Synchronous int

const (
	SynchronousOff Synchronous = iota
	SynchronousNormal
	SynchronousFull
	SynchronousExtra
)

func (s Synchronous) String() string {
	switch s {
	case SynchronousOff:
		return "OFF"
	case SynchronousNormal:
		return "NORMAL"
	case SynchronousFull:
		return "FULL"
	case SynchronousExtra:
		return "EXTRA"
	default:
		return ""
	}
}

// TimeZoneLocation represents all possible values for time zone location
type TimeZoneLocation int

const (
	TimeZoneLocationAuto TimeZoneLocation = iota
)

func (tzl TimeZoneLocation) String() string {
	if tzl != TimeZoneLocationAuto {
		return ""
	}
	return "auto"
}

// TransactionLock represents all possible values for transaction locking behaviour
type TransactionLock int

const (
	TransactionLockImmediate TransactionLock = iota
	TransactionLockDeferred
	TransactionLockExclusive
)

func (tl TransactionLock) String() string {
	switch tl {
	case TransactionLockImmediate:
		return "immediate"
	case TransactionLockDeferred:
		return "deferred"
	case TransactionLockExclusive:
		return "exclusive"
	default:
		return ""
	}
}

// UACrypt represents all possible values for user authentication module option crypt
type UACrypt int

const (
	UACryptSHA1 UACrypt = iota
	UACryptSSHA1
	UACryptSHA256
	UACryptSSHA256
	UACryptSHA384
	UACryptSSHA384
	UACryptSHA512
	UACryptSSHA512
)

func (c UACrypt) String() string {
	switch c {
	case UACryptSHA1:
		return "SHA1"
	case UACryptSSHA1:
		return "SSHA1"
	case UACryptSHA256:
		return "SHA256"
	case UACryptSSHA256:
		return "SSHA256"
	case UACryptSHA384:
		return "SHA384"
	case UACryptSSHA384:
		return "SSHA384"
	case UACryptSHA512:
		return "SHA512"
	case UACryptSSHA512:
		return "SSHA512"
	default:
		return ""
	}
}
