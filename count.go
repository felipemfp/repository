package repository

import "github.com/globalsign/mgo"

func Count(r Repository, params ...interface{}) (n int, err error) {
	err = r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		query, err := Query(r, db.C(r.GetCollectionName()), params...)
		if err != nil {
			return err
		}
		n, err = query.Count()
		return err
	})
	return
}
