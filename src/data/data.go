package data

import (
	"database/sql"
	"fmt"

	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func Connect(user string, pass string, name string, host string, port int) (contract.DataManager, error) {

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, name)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	p := new(pool)
	p.pool = db
	p.repo = newRepo(db)

	return p, nil
}

var _ contract.DataManager = &pool{}

type pool struct {
	repo
	pool *sql.DB
}

func (p *pool) Begin() (contract.TransactionManager, error) {

	tx, err := p.pool.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	t := new(transaction)
	t.transaction = tx
	t.repo = newRepo(tx)

	return t, nil
}

func (p *pool) Close() error {

	err := p.pool.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

var _ contract.TransactionManager = &transaction{}

type transaction struct {
	repo
	transaction *sql.Tx
	committed   bool
	rolledback  bool
}

func (t *transaction) Rollback() error {

	if !t.committed && !t.rolledback {

		err := t.transaction.Rollback()
		if err != nil {
			return errors.WithStack(err)
		}

		t.rolledback = true
	}

	return nil
}

func (t *transaction) Commit() error {

	err := t.transaction.Commit()
	if err != nil {
		return errors.WithStack(err)
	}

	t.committed = true
	return nil
}

var _ contract.RepoManager = repo{}

type repo struct {
	champ  champRepo
	team   teamRepo
	trophy trophyRepo
}

func newRepo(ex executor) repo {

	var r repo
	r.champ = champRepo{ex}
	r.team = teamRepo{ex}
	r.trophy = trophyRepo{ex}
	return r
}

func (r repo) Champ() contract.ChampRepo {
	return r.champ
}

func (r repo) Team() contract.TeamRepo {
	return r.team
}

func (r repo) Trophy() contract.TrophyRepo {
	return r.trophy
}

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
