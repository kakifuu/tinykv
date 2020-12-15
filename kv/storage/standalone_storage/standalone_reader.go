package standalone_storage

import (
	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
)

type StandAloneReader struct {
	db  *badger.DB
	txn *badger.Txn
}

func (s *StandAloneReader) GetCF(cf string, key []byte) ([]byte, error) {
	val, err := engine_util.GetCF(s.db, cf, key)
	if err != nil && err != badger.ErrKeyNotFound {
		return nil, err
	}
	return val, nil
}

func (s *StandAloneReader) IterCF(cf string) engine_util.DBIterator {
	return engine_util.NewCFIterator(cf, s.txn)
}

func (s *StandAloneReader) Close() {
	s.txn.Discard()
}
