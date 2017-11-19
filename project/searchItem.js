db.getCollection('musicInfo').find().pretty()

#show all artist, return the item id of the target item
db.getCollection('musicInfo').find({artist: {$exists: true}}, {_id: 1})

#search by artist contain 'head', return the item id of the target item
db.getCollection('musicInfo').find({artist: {$regex : ".*head.*"}}, {_id: 1})

# search by song/album/details contain 'oo', return the item id of the target item
db.getCollection('musicInfo').find(
{ $or: [
        { "albums.title": {$regex : ".*oo.*"} },
        { "albums.songs.title": {$regex: ".*oo.*"} },
        { "albums.description": {$regex: ".*oo.*"} }
        ]
},
{_id: 1})

# view albums ordered by price in descending order
db.getCollection('musicInfo').find({},{"_id": 1,"albums.price": 1}).sort( {"albums.price": -1})