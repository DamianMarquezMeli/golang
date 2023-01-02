package repositories

/*
   Interacción con la fuente de datos (CRUD) (ej: BBDD).
*/

import (
	"context"
	"time"

	db "github.com/devpablocristo/transactions/databases"
	trs "github.com/devpablocristo/transactions/models/trss"
	errors "github.com/devpablocristo/transactions/utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection = db.GetCollection("transactions")
var ctx = context.Background()

//var ctx = context.TODO()

func CreateTrs(u trs.Trs) (*trs.Trs, *errors.RestErr) {
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Inserting new document.")
		return nil, restErr
	}
	return &u, nil
}

func GetTrss() (*trs.Trss, *errors.RestErr) {
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Reading collection.")
		return nil, restErr
	}

	var urs trs.Trss
	for cur.Next(ctx) {
		var u trs.Trs
		err := cur.Decode(&u)
		if err != nil {
			restErr := errors.BadRequestError("ERROR! Reading documents.")
			return nil, restErr
		}
		urs = append(urs, &u)
	}
	return &urs, nil
}

func GetTrs(uId string) (*trs.Trs, *errors.RestErr) {
	oid, _ := primitive.ObjectIDFromHex(uId)
	filter := bson.M{"_id": oid}

	var u trs.Trs
	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Reading document.")
		return nil, restErr
	}

	return &u, nil
}

func UpdateTrs(t trs.Trs, uId string) (*trs.Trs, *errors.RestErr) {
	var err error
	oid, _ := primitive.ObjectIDFromHex(uId)
	filter := bson.M{"_id": oid}

	l, _ := time.LoadLocation("America/Buenos_Aires")
	tm := time.Now()

	now := tm.In(l)

	/*update := bson.M{
		"$set": bson.M{
			"trsname":   t.Trsname,
			"password":   t.Password,
			"email":      t.Email,
			"updated_at": now,
		},
	}*/

	update := bson.M{
		"$set": bson.M{
			"description": t.Description,
			"trsname":     t.Username,
			"password":    t.Password,
			"fullname":    t.Fullname,
			"phone":       t.Phone,
			"date":        now,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Trying to update.")
		return nil, restErr
	}

	updU, rErr := GetTrs(uId)
	if rErr != nil {
		restErr := errors.BadRequestError("ERROR! Finding new insertion.")
		return nil, restErr
	}

	return updU, nil
}

func DeleteTrs(uId string) (*int64, *errors.RestErr) {
	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(uId)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Using the uId.")
		return nil, restErr
	}

	filter := bson.M{"_id": oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Try to delete document.")
		return nil, restErr
	}

	return &result.DeletedCount, nil
}

func GetIdLastInseted() (bson.M, *errors.RestErr) {

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	var lastDocument bson.M
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&lastDocument)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Trying to get last inseted document.")
		return nil, restErr
	}
	//sId := lastrecord["_id"].(primitive.ObjectID).Hex()
	//fmt.Println(err, sId)

	return lastDocument, nil
}
