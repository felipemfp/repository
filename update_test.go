package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"."
)

var _ = Describe("Update", func() {
	BeforeEach(func() {
		Expect(clearSession()).To(BeNil())
	})

	It("should update an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.Update(r, obj.ID, testRepObject{
			ID:   obj.ID,
			Name: "Scarlett",
			Age:  22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Scarlett"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})

	It("should update partially an object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		obj := &testRepObject{
			ID:   bson.NewObjectId(),
			Name: "Snake Eyes",
			Age:  33,
		}
		Expect(repository.Create(r, obj)).To(BeNil())
		Expect(repository.Update(r, obj.ID, testRepObject{
			Age: 22,
		})).To(BeNil())
		Expect(defaultQueryRunner.RunWithDB(func(db *mgo.Database) error {
			c := db.C(r.GetCollectionName())
			objs := make([]testRepObject, 0)
			Expect(c.Find(nil).All(&objs)).To(BeNil())
			Expect(objs).To(HaveLen(1))
			Expect(objs[0].Name).To(Equal("Snake Eyes"))
			Expect(objs[0].Age).To(Equal(22))
			return nil
		})).To(BeNil())
	})

	It("should fail updating an non existent object", func() {
		r := &testRepNoDefaultCriteriaNoDefaultSorting{}
		Expect(repository.Update(r, bson.NewObjectId(), testRepObject{
			Name: "Scarlett",
			Age:  22,
		})).To(Equal(mgo.ErrNotFound))
	})
})