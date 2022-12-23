migrate((db) => {
  const collection = new Collection({
    "id": "5fzfo36oufdleyp",
    "created": "2022-12-23 10:45:23.757Z",
    "updated": "2022-12-23 10:45:23.757Z",
    "name": "artists",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "oqhrr1qf",
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
        "id": "diwabdyf",
        "name": "url",
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
  const collection = dao.findCollectionByNameOrId("5fzfo36oufdleyp");

  return dao.deleteCollection(collection);
})
