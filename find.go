package repository

import (
	"github.com/globalsign/mgo"
)

func Find(r Repository, dst interface{}, params ...interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		query, err := Query(r, db.C(r.GetCollectionName()), params...)
		if err != nil {
			return err
		}
		return query.One(dst)
	})
}
