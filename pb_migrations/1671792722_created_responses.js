migrate((db) => {
  const collection = new Collection({
    "id": "t9kom0hbpnxys70",
    "created": "2022-12-23 10:52:02.304Z",
    "updated": "2022-12-23 10:52:02.304Z",
    "name": "responses",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "jrshmisn",
        "name": "response",
        "type": "select",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "values": [
            "interested",
            "going",
            "not-going"
          ]
        }
      },
      {
        "system": false,
        "id": "birfqstq",
        "name": "event",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "collectionId": "t8vgzp8x1rt9s5o",
          "cascadeDelete": true
        }
      },
      {
        "system": false,
        "id": "1dwzizxo",
        "name": "user",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "collectionId": "_pb_users_auth_",
          "cascadeDelete": true
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
  const collection = dao.findCollectionByNameOrId("t9kom0hbpnxys70");

  return dao.deleteCollection(collection);
})
