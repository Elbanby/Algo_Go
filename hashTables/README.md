# Hash Table

## Note
I implemented the hash table on top of the map. However, I disabled the functionality of growing your map past the size you specify. Since, maps must have unique keys, I also didn't allow duplicate keys. However, you can add any logic you desire to handle duplicates easily.

The use case I visualized while building this, is a user name system where I wouldn't like to have two users share the same name. Also I added the update method to update the value for a specified key. This way I prevent accidental value change for a key.
