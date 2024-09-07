class Twitter() {
 
    val users : HashMap<Int, HashSet<Int>> = HashMap()
    val posts: HashMap<Int, MutableList<Post>> = HashMap()
    var sequenceNumber = 0
 
    fun postTweet(userId: Int, tweetId: Int) {
        addPost(userId,tweetId)
    }
 
    fun getNewsFeed(userId: Int): List<Int> {
 
        val listOfFollowee = users[userId]
        val totalPost = mutableListOf<Post>()
        listOfFollowee?.let{
            for (followee in listOfFollowee){
                val followeePost = posts[followee]
 
                if(followeePost != null){
                    for (post in followeePost){
                        totalPost.add(post)
                    }
                }
            }
        }
 
        val userPosts = posts[userId]
 
        if(userPosts != null){
            for (post in userPosts){
                totalPost.add(post)
            }
        }
 
        return totalPost.sortedByDescending {it.sequenceNumber}.map{it.postId}.take(10)
    }
 
 
 
    fun follow(followerId: Int, followeeId: Int) {
        if(!isUserExist(followerId)){
            insertUser(followerId)
        }
 
        users[followerId]?.add(followeeId)
    }
 
    fun unfollow(followerId: Int, followeeId: Int) {    
        users[followerId]?.remove(followeeId)
    }
 
    private fun addPost(userId: Int, tweetId: Int){
        if(!isUserExist(userId)){
            insertUser(userId)
        } 
 
        val postList = posts.get(userId)
        sequenceNumber++
        postList?.apply{ 
            add(Post(tweetId, sequenceNumber))
            posts[userId] = postList ?: mutableListOf()
        } ?: run{
            posts[userId] = mutableListOf(Post(tweetId, sequenceNumber))
        }
    }
 
    private fun isUserExist(userId: Int): Boolean =
       users.containsKey(userId)
 
    private fun insertUser(userId : Int){
        users.put(userId, hashSetOf<Int>())
    }
}
 
data class Post(
    val postId: Int,
    val sequenceNumber: Int
)
