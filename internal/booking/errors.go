package booking

import "errors"

var HouseDoesNotExist = errors.New("House with this ID does not exist")
var InvalidHouseData = errors.New("Provided house data is invalid")
var UnexpectedFailure = errors.New("Unexpected database error occured")
