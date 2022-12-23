migrate((db) => {
  const collection = new Collection({
    "id": "fzcwkkfaxajwz7s",
    "created": "2022-12-23 10:44:23.137Z",
    "updated": "2022-12-23 10:44:23.137Z",
    "name": "venues",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "cesdrmwq",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": true,
        "options": {
          "min": 2,
          "max": 40,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "cnyvfwb8",
        "name": "city",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "collectionId": "lo9oksd528cajzz",
          "cascadeDelete": true
        }
      },
      {
        "system": false,
        "id": "db0upfqh",
        "name": "link",
        "type": "url",
        "required": false,
        "unique": false,
        "options": {
          "exceptDomains": null,
          "onlyDomains": null
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("fzcwkkfaxajwz7s");

  return dao.deleteCollection(collection);
})
