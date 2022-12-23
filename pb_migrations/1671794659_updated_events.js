migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t8vgzp8x1rt9s5o")

  collection.listRule = ""
  collection.viewRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t8vgzp8x1rt9s5o")

  collection.listRule = null
  collection.viewRule = null

  return dao.saveCollection(collection)
})
