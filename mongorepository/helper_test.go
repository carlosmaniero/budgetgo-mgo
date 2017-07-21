package mongorepository

import mgo "gopkg.in/mgo.v2"

var session, _ = mgo.Dial("localhost")
var db = session.DB("budgetgo")
