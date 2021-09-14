package errorconstant

import "errors"

var ErrorUnableToSave error = errors.New("unable to save data")
var ErrorUnableToDelete error = errors.New("unable to delete data")
var ErrorUnableToFetch error = errors.New("unable to fetch data")
var IDNotFound error = errors.New("id not found ")
