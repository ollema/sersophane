migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5fzfo36oufdleyp")

  collection.listRule = ""
  collection.viewRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5fzfo36oufdleyp")

  collection.listRule = null
  collection.viewRule = null

  return dao.saveCollection(collection)
})
