migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("lo9oksd528cajzz")

  collection.listRule = ""
  collection.viewRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("lo9oksd528cajzz")

  collection.listRule = null
  collection.viewRule = null

  return dao.saveCollection(collection)
})
