package statedb

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var emptyCodeHash = crypto.Keccak256(nil)



type Account struct {
	Nonce    uint64
	Balance  *big.Int
	CodeHash []byte
}


func NewEmptyAccount() *Account {
	return &Account{
		Balance:  new(big.Int),
		CodeHash: emptyCodeHash,
	}
}


func (acct Account) IsContract() bool {
	return !bytes.Equal(acct.CodeHash, emptyCodeHash)
}


type Storage map[common.Hash]common.Hash


func (s Storage) SortedKeys() []common.Hash {
	keys := make([]common.Hash, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return bytes.Compare(keys[i].Bytes(), keys[j].Bytes()) < 0
	})
	return keys
}


type stateObject struct {
	db *StateDB

	account Account
	code    []byte

	
	originStorage Storage
	dirtyStorage  Storage

	address common.Address

	
	dirtyCode bool
	suicided  bool
}


func newObject(db *StateDB, address common.Address, account Account) *stateObject {
	if account.Balance == nil {
		account.Balance = new(big.Int)
	}
	if account.CodeHash == nil {
		account.CodeHash = emptyCodeHash
	}
	return &stateObject{
		db:            db,
		address:       address,
		account:       account,
		originStorage: make(Storage),
		dirtyStorage:  make(Storage),
	}
}


func (s *stateObject) empty() bool {
	return s.account.Nonce == 0 && s.account.Balance.Sign() == 0 && bytes.Equal(s.account.CodeHash, emptyCodeHash)
}

func (s *stateObject) markSuicided() {
	s.suicided = true
}



func (s *stateObject) AddBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	s.SetBalance(new(big.Int).Add(s.Balance(), amount))
}



func (s *stateObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	s.SetBalance(new(big.Int).Sub(s.Balance(), amount))
}


func (s *stateObject) SetBalance(amount *big.Int) {
	s.db.journal.append(balanceChange{
		account: &s.address,
		prev:    new(big.Int).Set(s.account.Balance),
	})
	s.setBalance(amount)
}

func (s *stateObject) setBalance(amount *big.Int) {
	s.account.Balance = amount
}

//

//


func (s *stateObject) Address() common.Address {
	return s.address
}


func (s *stateObject) Code() []byte {
	if s.code != nil {
		return s.code
	}
	if bytes.Equal(s.CodeHash(), emptyCodeHash) {
		return nil
	}
	code := s.db.keeper.GetCode(s.db.ctx, common.BytesToHash(s.CodeHash()))
	s.code = code
	return code
}



func (s *stateObject) CodeSize() int {
	return len(s.Code())
}


func (s *stateObject) SetCode(codeHash common.Hash, code []byte) {
	prevcode := s.Code()
	s.db.journal.append(codeChange{
		account:  &s.address,
		prevhash: s.CodeHash(),
		prevcode: prevcode,
	})
	s.setCode(codeHash, code)
}

func (s *stateObject) setCode(codeHash common.Hash, code []byte) {
	s.code = code
	s.account.CodeHash = codeHash[:]
	s.dirtyCode = true
}


func (s *stateObject) SetNonce(nonce uint64) {
	s.db.journal.append(nonceChange{
		account: &s.address,
		prev:    s.account.Nonce,
	})
	s.setNonce(nonce)
}

func (s *stateObject) setNonce(nonce uint64) {
	s.account.Nonce = nonce
}


func (s *stateObject) CodeHash() []byte {
	return s.account.CodeHash
}


func (s *stateObject) Balance() *big.Int {
	return s.account.Balance
}


func (s *stateObject) Nonce() uint64 {
	return s.account.Nonce
}


func (s *stateObject) GetCommittedState(key common.Hash) common.Hash {
	if value, cached := s.originStorage[key]; cached {
		return value
	}
	
	value := s.db.keeper.GetState(s.db.ctx, s.Address(), key)
	s.originStorage[key] = value
	return value
}


func (s *stateObject) GetState(key common.Hash) common.Hash {
	if value, dirty := s.dirtyStorage[key]; dirty {
		return value
	}
	return s.GetCommittedState(key)
}


func (s *stateObject) SetState(key common.Hash, value common.Hash) {
	
	prev := s.GetState(key)
	if prev == value {
		return
	}
	
	s.db.journal.append(storageChange{
		account:  &s.address,
		key:      key,
		prevalue: prev,
	})
	s.setState(key, value)
}

func (s *stateObject) setState(key, value common.Hash) {
	s.dirtyStorage[key] = value
}
