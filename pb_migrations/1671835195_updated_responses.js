migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t9kom0hbpnxys70")

  collection.createRule = "@request.auth.id = user.id"

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "oiwkkzdz",
    "name": "field",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t9kom0hbpnxys70")

  collection.createRule = null

  // remove
  collection.schema.removeField("oiwkkzdz")

  return dao.saveCollection(collection)
})
