let today = new Date();
const dbName = 'audit-flow'
const db = db.getSiblingDB(dbName);
const collections = db.getCollectionNames();

function CreateCollection(collections, db, collectionName){   
    if (!collections.includes(collectionName)) {
      db.createCollection(collectionName);
    }
  }
  
  function upsertDocument(db, collection, filter, document, dbName) {
    let update = { $set: document};
    let options = { upsert: true};
    db.getSiblingDB(dbName)[collection].updateOne(filter, update, options);
  }

db.createUser(
    {
      user: 'auditAdmim',
      pwd: 'f0cd47b4b7364a7e9b87e1a377b7dddf',
      roles: [{ role: 'readWrite', db: dbName }],
    },
  );

  CreateCollection(collections, db, 'auditory')

//auditory index
db.auditory.createIndex({entity_id: 1});
db.auditory.createIndex({entity_id: 1, dated_in:1});
db.auditory.createIndex({entity_id: 1, dated_in:1, Type:1});
db.auditory.createIndex({entity: 1});
db.auditory.createIndex({entity: 1, dated_in:1});
db.auditory.createIndex({entity: 1, dated_in:1, Type:1});





