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

  $(document).ready(function() {
    var errorContainer = $('#errorContainer'); // Error message container
  
    $('#uploadForm').submit(function(event) {
      event.preventDefault(); // Prevent the default form submission
  
      // Get the form data
      var titre = $('#titre').val();
      var description = $('#description').val();
      var dada = $('input[name=dada]:checked').val();
      var image = $('#image').prop('files')[0];
    
      if (!checkSessionCookie()) {
        alert("You have to be connected to post. Please login.");
        // Create a form dynamically
        var form = $('<form>');
        form.attr('action', '/login');
        form.attr('method', 'get');
        $('body').append(form);
        form.submit();
        return;
    }

      // Perform validation
      if (!titre) {
        showError('Please enter a title.'); // Show the error message
        return;
      }
      
  
      // Clear any previous error message
      clearError();
  
      // Create a new FormData object
      var formData = new FormData();
      // Add form data to FormData object
      formData.append('titre', titre);
      formData.append('description', description);
      formData.append('dada', dada);
      if (image) {
        formData.append('image', image);
      }
  
      // Send the form data to the server using AJAX
      $.ajax({
        url: '/upload', // Update the URL to your server endpoint
        type: 'POST',
        data: formData,
        processData: false,
        contentType: false,
        success: function(response) {
          // Handle the response from the server
          console.log(response); // You can do further processing or display a success message
          post(); // Call the post() function to display the new post
        },
        error: function(xhr, status, error) {
          // Handle the error
          console.log(error); // You can display an error message or perform other error handling
        }
      });
    });
    // Function to show the error message
    function showError(message) {
      errorContainer.text(message);
    }
    // Function to clear the error message
    function clearError() {
      errorContainer.text('');
    }

    function checkSessionCookie() {
        var cookieValue = getCookie("session");
        return cookieValue !== null && cookieValue !== "";
    }

    function getCookie(name) {
        var cookies = document.cookie.split(';');
        for (var i = 0; i < cookies.length; i++) {
            var cookie = cookies[i].trim();
            if (cookie.startsWith(name + '=')) {
                return cookie.substring(name.length + 1);
            }
        }
        return null;
    }
  });