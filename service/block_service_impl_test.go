package blockchain

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"os"
	"it-chain/domain"
)

func TestLedger_CreateBlock(t *testing.T) {
	path := "./test_db"
	defer func(){
		os.RemoveAll(path)
	}()
	ledger := NewLedger(path)
	txList := make([]*domain.Transaction, 0)
	for i := 0; i < 999; i++{
		tx := domain.CreateNewTransaction("1", "1", domain.General, time.Now(), domain.SetTxData("", domain.Invoke, domain.SetTxMethodParameters(0, "", []string{""}), ""))
		tx.GenerateHash()
		txList = append(txList, tx)
	}

	_, err := ledger.CreateBlock(txList, "123")

	assert.NoError(t, err)
}

func TestLedger_AddBlock(t *testing.T) {
	path := "./test_db"
	defer func(){
		os.RemoveAll(path)
	}()
	ledger := NewLedger(path)
	txList := make([]*domain.Transaction, 0)
	for i := 0; i < 999; i++{
		tx := domain.CreateNewTransaction("1", "1", domain.General, time.Now(), domain.SetTxData("", domain.Invoke, domain.SetTxMethodParameters(0, "", []string{""}), ""))
		tx.GenerateHash()
		txList = append(txList, tx)
	}
	blk, err := ledger.CreateBlock(txList, "123")
	assert.NoError(t, err)

	_, err = ledger.AddBlock(blk)
	assert.NoError(t, err)
}

func TestLedger_GetLastBlock(t *testing.T) {
	path := "./test_db"
	defer func(){
		os.RemoveAll(path)
	}()
	ledger := NewLedger(path)
	txList := make([]*domain.Transaction, 0)
	for i := 0; i < 999; i++{
		tx := domain.CreateNewTransaction("1", "1", domain.General, time.Now(), domain.SetTxData("", domain.Invoke, domain.SetTxMethodParameters(0, "", []string{""}), ""))
		tx.GenerateHash()
		txList = append(txList, tx)
	}
	blk, err := ledger.CreateBlock(txList, "123")
	assert.NoError(t, err)

	_, err = ledger.AddBlock(blk)
	assert.NoError(t, err)

	lastblk, err := ledger.GetLastBlock()

	assert.Equal(t, blk, lastblk)
}

func TestLedger_LookUpBlock(t *testing.T) {
	path := "./test_db"
	defer func(){
		os.RemoveAll(path)
	}()
	ledger := NewLedger(path)
	txList := make([]*domain.Transaction, 0)
	for i := 0; i < 999; i++{
		tx := domain.CreateNewTransaction("1", "1", domain.General, time.Now(), domain.SetTxData("", domain.Invoke, domain.SetTxMethodParameters(0, "", []string{""}), ""))
		tx.GenerateHash()
		txList = append(txList, tx)
	}
	blk, err := ledger.CreateBlock(txList, "123")
	assert.NoError(t, err)

	_, err = ledger.AddBlock(blk)
	assert.NoError(t, err)

	retBlk, err := ledger.LookUpBlock(0)
	assert.NoError(t, err)
	assert.Equal(t, blk, retBlk)

	retBlk, err = ledger.LookUpBlock(blk.Header.BlockHash)
	assert.NoError(t, err)
	assert.Equal(t, blk, retBlk)
}

func TestLedger_VerifyBlock(t *testing.T) {
	path := "./test_db"
	defer func(){
		os.RemoveAll(path)
	}()
	ledger := NewLedger(path)
	txList := make([]*domain.Transaction, 0)
	for i := 0; i < 999; i++{
		tx := domain.CreateNewTransaction("1", "1", domain.General, time.Now(), domain.SetTxData("", domain.Invoke, domain.SetTxMethodParameters(0, "", []string{""}), ""))
		tx.GenerateHash()
		txList = append(txList, tx)
	}
	blk, err := ledger.CreateBlock(txList, "123")
	assert.NoError(t, err)

	_, err = ledger.AddBlock(blk)
	assert.NoError(t, err)

	ledger.VerifyBlock(blk)
}