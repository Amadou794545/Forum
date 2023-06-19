function toggleLikePost(userId, postId) {
    const existingLike = db.run(`SELECT * FROM likes WHERE userId = ${userId} AND postId = ${postId}`);
  
    if (existingLike) {
      // User has already liked the post, remove the like
      db.run(`UPDATE posts SET likes = likes - 1 WHERE id = ${postId}`);
      db.run(`DELETE FROM likes WHERE userId = ${userId} AND postId = ${postId}`);
      console.log("Like removed from the post.");
    } else {
      // User has not liked the post, add the like
      db.run(`UPDATE posts SET likes = likes + 1 WHERE id = ${postId}`);
      db.run(`INSERT INTO likes (userId, postId) VALUES (${userId}, ${postId})`);
      console.log("Post liked successfully!");
    }
  }
  
  function toggleDislikePost(userId, postId) {
    const existingDislike = db.run(`SELECT * FROM likes WHERE userId = ${userId} AND postId = ${postId}`);
  
    if (existingDislike) {
      // User has already disliked the post, remove the dislike
      db.run(`UPDATE posts SET dislikes = dislikes - 1 WHERE id = ${postId}`);
      db.run(`DELETE FROM likes WHERE userId = ${userId} AND postId = ${postId}`);
      console.log("Dislike removed from the post.");
    } else {
      // User has not disliked the post, add the dislike
      db.run(`UPDATE posts SET dislikes = dislikes + 1 WHERE id = ${postId}`);
      db.run(`INSERT INTO likes (userId, postId) VALUES (${userId}, ${postId})`);
      console.log("Post disliked successfully!");
    }
  }
  
  
  
  function toggleLikeComment(userId, commentId) {
    const existingLike = db.run(`SELECT * FROM likes WHERE userId = ${userId} AND commentId = ${commentId}`);
  
    if (existingLike) {
      db.run(`UPDATE posts SET likes = likes - 1 WHERE id = ${commentId}`);
      db.run(`DELETE FROM likes WHERE userId = ${userId} AND postId = ${commentId}`);
      console.log("Like removed from the post.");
    } else {
      db.run(`UPDATE posts SET likes = likes + 1 WHERE id = ${commentId}`);
      db.run(`INSERT INTO likes (userId, postId) VALUES (${userId}, ${commentId})`);
      console.log("Post liked successfully!");
    }
  }
  
  function toggleDislikeComment(userId, commentId) {
    const existingDislike = db.run(`SELECT * FROM likes WHERE userId = ${userId} AND postId = ${commentId}`);
  
    if (existingDislike) {
      db.run(`UPDATE posts SET dislikes = dislikes - 1 WHERE id = ${commentId}`);
      db.run(`DELETE FROM likes WHERE userId = ${userId} AND postId = ${commentId}`);
      console.log("Dislike removed from the post.");
    } else {
      db.run(`UPDATE posts SET dislikes = dislikes + 1 WHERE id = ${commentId}`);
      db.run(`INSERT INTO likes (userId, postId) VALUES (${userId}, ${commentId})`);
      console.log("Post disliked successfully!");
    }
  }
  