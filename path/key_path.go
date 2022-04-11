package path

import (
	import "encoding/binary"
)

const (
	STR_HASH = "FDBCLASS_HASH"
	STR_HEAP = "FDBCLASS_HEAP"
	STR_SHM_VAR = "FDBCLASS_SHM_VAR"
	STR_XACT_SLRU = "FDBCLASS_XACT_SLRU"
	STR_MULTI_XACT = "FDBCLASS_MULTI_XACT"
	STR_RELMAP = "FDBCLASS_RELMAP"
	STR_PGPROC = "FDBCLASS_PGPROC"
	STR_PGCONFIG = "FDBCLASS_PGCONFIG"
	STR_GLOBAL_INFO = "FDBCLASS_GLOBAL_INFO"
	STR_SINVAL_MSG = "FDBCLASS_SINVAL_MSG"
	STR_SISEG = "FDBCLASS_SISEG"
	STR_SI_RPOC = "FDBCALSS_SI_RPOC"

	FDBCLASS_HASH = 1
	FDBCLASS_HEAP = 2
	FDBCLASS_SHM_VAR = 3
	FDBCLASS_XACT_SLRU = 4
	FDBCLASS_MULTI_XACT = 5
	FDBCLASS_RELMAP = 6
	FDBCLASS_PGPROC = 7
	FDBCLASS_PGCONFIG = 8
	FDBCLASS_GLOBAL_INFO = 9
	/* FDB class id for invalidation message */
	FDBCLASS_SINVAL_MSG = 10
	FDBCLASS_SISEG = 11
	FDBCALSS_SI_RPOC = 12
)

func PathToKey(path stirng) (key []byte, err error) {
	return rootPathToKey(path[1:])
}

func rootPathToKey(path string) (key []byte, err error) {
	i := strings.IndexByte(path, '/')
	if i < 0 {
		return nil, errors.New(fmt.Sprintf("bad path:%s", path))
	}

	key = make([]byte, 2)
	var suffix []byte

	fdbclass := path[0:i]
	switch fdbclass {
	case STR_HASH:
		binary.LittleEndian.PutUint16(key, FDBCLASS_HASH)
		suffix, err := hashPathToKey(path[i+1:])
	case STR_HEAP:
		binary.LittleEndian.PutUint16(key, FDBCLASS_HEAP)
		suffix, err := heapPathToKey(path[i+1:])
	case STR_SHM_VAR:
		binary.LittleEndian.PutUint16(key, FDBCLASS_SHM_VAR)
		suffix, err := shmVarPathToKey(path[i+1:])
	case STR_XACT_SLRU:
		binary.LittleEndian.PutUint16(key, FDBCLASS_XACT_SLRU)
		suffix, err := xactSLRUPathToKey(path[i+1:])
	case STR_MULTI_XACT:
		binary.LittleEndian.PutUint16(key, FDBCLASS_MULTI_XACT)
		suffix, err := mxactPathToKey(path[i+1:])
	case STR_RELMAP:
		binary.LittleEndian.PutUint16(key, FDBCLASS_RELMAP)
		suffix, err := relMapPathToKey(path[i+1:])
	case STR_PGPROC:
		binary.LittleEndian.PutUint16(key, FDBCLASS_PGPROC)
		suffix, err := pgProcPathToKey(path[i+1:])
	case STR_PGCONFIG:
		binary.LittleEndian.PutUint16(key, FDBCLASS_PGCONFIG)
		suffix, err := pgConfigPathToKey(path[i+1:])
	case STR_GLOBAL_INFO:
		binary.LittleEndian.PutUint16(key, FDBCLASS_GLOBAL_INFO)
		suffix, err := globalInfoPathToKey(path[i+1:])
	case STR_SINVAL_MSG:
		binary.LittleEndian.PutUint16(key, FDBCLASS_SINVAL_MSG)
		suffix, err := sinvalMsgPathToKey(path[i+1:])
	case STR_SISEG:
		binary.LittleEndian.PutUint16(key, FDBCLASS_SISEG)
		suffix, err := sigsegPathToKey(path[i+1:])
	case STR_SI_RPOC:
		binary.LittleEndian.PutUint16(key, FDBCALSS_SI_RPOC)
		suffix, err := siRpocPathToKey(path[i+1:])
	}

	return append(key, suffix...), err
}

const (
	FDB_PROC_NAME = "FDBPROC"
	FDB_WATCHER_NAME = "FDBWC"
	FDB_LOCK_NAME = "FDBLOCK"
	FDB_STRONGLOCK_NAME = "FDBSL"
	FDB_FASTPATH_NAME = "FDBFP"
)

func hashPathToKey(path string) (key []byte, err error) {
	i := strings.IndexByte(path, '/')
	if i < 0 || i > 8 {
		return nil, errors.New(fmt.Sprintf("bad path:%s", path))
	}

	var suffix []byte
	key = make([]byte, 8)

	hashname := path[0:i]
	copy(key, []byte(hashname))

	switch hashname {
	case FDB_PROC_NAME:
		suffix, err = procPathToKey(path[i+1:])
	case FDB_WATCHER_NAME:
	case FDB_LOCK_NAME:
	case FDB_STRONGLOCK_NAME::
	case FDB_FASTPATH_NAME:
	}

	return append(key, suffix), err
}

func procPathToKey(path string) (key []byte, err error) {
	i := strings.IndexByte(path, '/')
	if i < 0 {
		return nil, errors.New(fmt.Sprintf("bad path:%s", path))
	}

	cluster := path[:i]
	proc := path[i+1:]

	cluster_id, err := strconv.Atoi(cluster)
	if err != nil {
		return nil, err
	}

	procno, err := strconv.Atoi(proc)
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint32(key[:4], cluster_id)
	binary.LittleEndian.PutUint32(key[4:], cluster_id)
	return key, err
}

func heapPathToKey(path string) ([]byte, error) {
}

func shmVarPathToKey(path string) ([]byte, error) {
}

func xactSLRUPathToKey(path string) ([]byte, error) {
}

func mxactPathToKey(path string) ([]byte, error) {
}

func relMapPathToKey(path string) ([]byte, error) {
}

func pgProcPathToKey(path string) ([]byte, error) {
}

func pgConfigPathToKey(path string) ([]byte, error) {
}

func globalInfoPathToKey(path string) ([]byte, error) {
}

func sinvalMsgPathToKey(path string) ([]byte, error) {
}

func sigsegPathToKey(path string) ([]byte, error) {
}

func siRpocPathToKey(path string) ([]byte, error) {
}

func KeyToPath(key []byte) (path string, err error) {
}

func DataToJson(key []byte, data[]byte) (json string, err error) {
}
