package dao

import (
	"log"

	. "github.com/serj113/contacts/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContactsDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "contact"
)

func (m *ContactsDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ContactsDao) FindAll() ([]Contact, error) {
	var contacts []Contact
	err := db.C(COLLECTION).Find(bson.M{}).All(&contacts)
	return contacts, err
}

func (m *ContactsDao) FindById(id string) (Contact, error) {
	var contact Contact
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&contact)
	return contact, err
}

func (m *ContactsDao) Insert(contact Contact) error {
	err := db.C(COLLECTION).Insert(&contact)
	return err
}

func (m *ContactsDao) Delete(contact Contact) error {
	err := db.C(COLLECTION).Remove(&contact)
	return err
}

func (m *ContactsDao) Update(contact Contact) error {
	err := db.C(COLLECTION).UpdateId(contact.ID, &contact)
	return err
}
