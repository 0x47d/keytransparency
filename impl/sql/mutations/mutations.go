// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package mutations defines operations to write and read mutations to and from
// the database.
package mutations

import (
	"database/sql"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/google/keytransparency/core/mutator"
	"github.com/google/keytransparency/core/transaction"

	tpb "github.com/google/keytransparency/core/proto/keytransparency_v1_types"
)

const (
	insertMapRowExpr = `INSERT INTO Maps (MapID) VALUES (?);`
	countMapRowExpr  = `SELECT COUNT(*) AS count FROM Maps WHERE MapID = ?;`
	insertExpr       = `
	INSERT INTO Mutations (MapID, MIndex, Mutation)
	VALUES (?, ?, ?);`
	readExpr = `
	SELECT Mutation FROM Mutations
	WHERE MapID = ? AND Sequence >= ?
	ORDER BY Sequence ASC LIMIT ?;`
)

type mutations struct {
	mapID int64
	db    *sql.DB
}

// New creates a new mutations instance.
func New(db *sql.DB, mapID int64) (mutator.Mutation, error) {
	m := &mutations{
		mapID: mapID,
		db:    db,
	}

	// Create tables and map entry.
	if err := m.create(); err != nil {
		return nil, err
	}
	if err := m.insertMapRow(); err != nil {
		return nil, err
	}
	return m, nil
}

// ReadRange reads all mutations for a specific given mapID and sequence range.
// The range is identified by a starting sequence number and a count.
func (m *mutations) ReadRange(txn transaction.Txn, startSequence uint64, count int) ([]*tpb.SignedKV, error) {
	readStmt, err := txn.Prepare(readExpr)
	if err != nil {
		return nil, err
	}
	defer readStmt.Close()
	rows, err := readStmt.Query(m.mapID, startSequence, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*tpb.SignedKV, 0, count)
	for rows.Next() {
		var mData []byte
		if err := rows.Scan(&mData); err != nil {
			return nil, err
		}
		mutation := new(tpb.SignedKV)
		if err := proto.Unmarshal(mData, mutation); err != nil {
			return nil, err
		}
		results = append(results, mutation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// Write saves the mutation in the database. Write returns the auto-inserted
// sequence number.
func (m *mutations) Write(txn transaction.Txn, mutation *tpb.SignedKV) (uint64, error) {
	index := mutation.GetKeyValue().Key
	mData, err := proto.Marshal(mutation)
	if err != nil {
		return 0, err
	}

	writeStmt, err := txn.Prepare(insertExpr)
	if err != nil {
		return 0, err
	}
	defer writeStmt.Close()
	result, err := writeStmt.Exec(m.mapID, index, mData)
	if err != nil {
		return 0, err
	}
	sequence, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(sequence), nil
}

// Create creates new database tables.
func (m *mutations) create() error {
	for _, stmt := range createStmt {
		_, err := m.db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("Failed to create mutation tables: %v", err)
		}
	}
	return nil
}

func (m *mutations) insertMapRow() error {
	// Check if a map row does not exist for the same MapID.
	countStmt, err := m.db.Prepare(countMapRowExpr)
	if err != nil {
		return fmt.Errorf("insertMapRow(): %v", err)
	}
	defer countStmt.Close()
	var count int
	if err := countStmt.QueryRow(m.mapID).Scan(&count); err != nil {
		return fmt.Errorf("insertMapRow(): %v", err)
	}
	if count >= 1 {
		return nil
	}

	// Insert a map row if it does not exist already.
	insertStmt, err := m.db.Prepare(insertMapRowExpr)
	if err != nil {
		return fmt.Errorf("insertMapRow(): %v", err)
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(m.mapID)
	if err != nil {
		return fmt.Errorf("insertMapRow(): %v", err)
	}
	return nil
}
