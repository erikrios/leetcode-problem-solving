class Solution {
    fun canVisitAllRooms(rooms: List<List<Int>>): Boolean {
    val size = rooms.size
    val roomsId = HashSet<Int>()
    val keys = HashSet<Int>()
 
    for (i in 0 until size) {
        roomsId.add(i)
    }
 
    while (roomsId.isNotEmpty()) {
        val ids = roomsId.toList()
 
        var openedCount = 0
 
        for (room in ids){
            if (room == 0 || keys.contains(room)){
                for(each in rooms[room]){
                    keys.add(each)
                }
 
                roomsId.remove(room)
                openedCount++
            }
        }
 
        if (openedCount == 0) {
            return false
        }
    }
 
    return true
  }
}
