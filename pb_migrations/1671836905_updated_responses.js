migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t9kom0hbpnxys70")

  collection.updateRule = "@request.auth.id = user.id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t9kom0hbpnxys70")

  collection.updateRule = null

  return dao.saveCollection(collection)
})
