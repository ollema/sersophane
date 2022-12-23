migrate((db) => {
  const collection = new Collection({
    "id": "t8vgzp8x1rt9s5o",
    "created": "2022-12-23 10:50:00.256Z",
    "updated": "2022-12-23 10:50:00.256Z",
    "name": "events",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "9psptsqv",
        "name": "name",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": 2,
          "max": 100,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "8zblkgm0",
        "name": "venue",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "collectionId": "fzcwkkfaxajwz7s",
          "cascadeDelete": true
        }
      },
      {
        "system": false,
        "id": "r6aup2no",
        "name": "artists",
        "type": "relation",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 10,
          "collectionId": "5fzfo36oufdleyp",
          "cascadeDelete": false
        }
      },
      {
        "system": false,
        "id": "wjelzixl",
        "name": "type",
        "type": "select",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "values": [
            "concert",
            "music-festival",
            "movie"
          ]
        }
      },
      {
        "system": false,
        "id": "isqby0kp",
        "name": "cancelled",
        "type": "bool",
        "required": false,
        "unique": false,
        "options": {}
      },
      {
        "system": false,
        "id": "4orkx1hq",
        "name": "starts",
        "type": "date",
        "required": true,
        "unique": false,
        "options": {
          "min": "",
          "max": ""
        }
      },
      {
        "system": false,
        "id": "kqdquhza",
        "name": "ends",
        "type": "date",
        "required": true,
        "unique": false,
        "options": {
          "min": "",
          "max": ""
        }
      },
      {
        "system": false,
        "id": "4h2er5jl",
        "name": "url",
        "type": "url",
        "required": false,
        "unique": false,
        "options": {
          "exceptDomains": [],
          "onlyDomains": []
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
  const collection = dao.findCollectionByNameOrId("t8vgzp8x1rt9s5o");

  return dao.deleteCollection(collection);
})
