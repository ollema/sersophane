migrate((db) => {
  const collection = new Collection({
    "id": "lo9oksd528cajzz",
    "created": "2022-12-23 10:42:32.039Z",
    "updated": "2022-12-23 10:42:32.039Z",
    "name": "cities",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "cfmoqnti",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": true,
        "options": {
          "min": 2,
          "max": 40,
          "pattern": ""
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
  const collection = dao.findCollectionByNameOrId("lo9oksd528cajzz");

  return dao.deleteCollection(collection);
})
