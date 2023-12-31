function toggleDiv() {
  var div = document.getElementById("myDiv");
  if (div.style.display === "block") {
    div.style.display = "none";
  } else {
    div.style.display = "block";
    document.getElementById("Description") // Placer le focus sur le champ de saisie "message"
  }
}

function Post() {
  var description = document.getElementById("description").value;
  var titre = document.getElementById("titre").value;
  var image = document.getElementById("image").files[0];
  var Container = document.getElementById("postContainer");
  // Création de la div principale
  var newDiv = document.createElement("div");
  var newTitre = document.createElement("h1");
  newTitre.className = 'Post-Title';
  var newDescription = document.createElement("p");
  newDescription.className = 'Post-Desc';
  var newImage = document.createElement("img");
  newTitre.textContent = "Titre : " + titre;
  newDescription.textContent = "Description : " + description;
  // Création de la div newDivPost
  var newDivPost = document.createElement("div");
  newDivPost.classList.add("div-Post");
  // Ajouter des classes CSS
  newTitre.classList.add("titre-style");
  newDiv.classList.add("div-item");
  var postContent = `
      <div id="post">
          <button id="likeButton">Like</button>
          <span id="likeCount">0</span>
          <button id="dislikeButton">Dislike</button>
          <span id="dislikeCount">0</span>
      </div>
  `;
  newDivPost.innerHTML = postContent;
  newDivPost.appendChild(newTitre);
  newDivPost.appendChild(newDescription);
  // Redimensionner et afficher l'image
  if (image) {
    var reader = new FileReader();
    reader.onload = function (event) {
      newImage.src = event.target.result; // Définir la source de l'image
      newImage.classList.add("image-style"); // Ajouter une classe CSS pour styliser l'image
      newDivPost.appendChild(newImage); // Ajouter l'élément <img> à la div nouvellement créée
    };
    reader.readAsDataURL(image); // Lire le fichier d'image en tant qu'URL de données
  }
  var firstDiv = Container.firstChild;
  Container.insertBefore(newDiv, firstDiv); // Insérer la nouvelle div avant le premier enfant existant
  Container.insertAdjacentHTML("beforeend", "<br>");
  // Gestionnaire d'événement pour le bouton "Like"
  var likeButton = newDivPost.querySelector("#likeButton");
  var likeCount = newDivPost.querySelector("#likeCount");
  var likeValue = 0;
  likeButton.addEventListener("click", function () {
    likeValue++;
    likeCount.textContent = likeValue;
  });
  let currentPag = 0;
  const postsPerPag = 1
  fetch(`/api/posts?page=${currentPag}&limit=${postsPerPag}`)
    .then(response => response.json())
    .then(data => {
      console.log(data);
      // Update the UI to display the posts
      data.forEach(post => {
        const postElement = document.createElement('div');
        postElement.innerHTML = `
  <form id="ID">  
  <input name="ID"  id="input" type="hidden" value="${post.ID}">
  </form>
        `;
        newDivPost.appendChild(postElement);
      })
    })
  newDiv.appendChild(newDivPost);
  document.getElementById("description").value = "";
  document.getElementById("titre").value = "";
  document.getElementById("image").value = "";
}

function toggleDivCom() {
  var div = document.getElementById("myDivCom");
  var commentDivs = document.getElementsByClassName("comment-container");
  if (div.style.display === "block") {
    div.style.display = "none";
    // Cacher toutes les div des commentaires
    for (var i = 0; i < commentDivs.length; i++) {
      commentDivs[i].style.display = "none";
    }
  } else {
    div.style.display = "block";
    document.getElementById("Comments"); // Placer le focus sur le champ de saisie "message"
    // Afficher toutes les div des commentaires précédemment créées
    for (var i = 0; i < commentDivs.length; i++) {
      commentDivs[i].style.display = "block";
    }
  }
}

function scrollListener() {
  if (isScrolledToBottom()) {
    fetchPosts();



  }
}

let currentPage = 1;
const postsPerPage = 25;

$(document).ready(function () {
  var errorContainer = $('#errorContainer'); // Error message container
  $('#filterform').submit(function (event) {
    event.preventDefault(); // Prevent the default form submission
    clearallpost();
    currentPage = 1;
    fetchPosts()
    
  })
  window.addEventListener('scroll', scrollListener);
});

function clearallpost() {
  const postContainer = document.getElementById('postContainer');
  postContainer.innerHTML = "";
}

// Attach the scroll listener to the window
// Function to fetch posts from the server
function fetchPosts() {
  // Send a GET request to the server to retrieve the posts
  var filters = [];
  $('input[type="checkbox"]:checked').each(function () {
    filters.push($(this).val());
  });
  console.log(filters)

  fetch(`/api/posts?page=${currentPage}&limit=${postsPerPage}&filters=${filters}`)
    .then(response => response.json())
    .then(data => {
      // Process the retrieved post data
      console.log(data); // Print the data to the console as an example
      // Update the UI to display the posts
      const postContainer = document.getElementById('postContainer');
      data.forEach(post => {
        const postElement = document.createElement('div');
        postElement.className = 'post';
        postElement.innerHTML = `
            <h3 class="Post-Title">${post.Title}</h3>
            <p class="Post-Desc">${post.Description}</p>
            <img src="${post.ImagePath}" alt="Post Image" class="Post-IMG">
            <input type="button" value="likes" name="Likes-${post.ID}" onclick="handleRadioClick(this)">
            <input type="button" value="Dislikes" name="Dislikes-${post.ID}" onclick="handleRadioClickDislike(this)">
            <input type="radio" id="Likes-${post.ID}" class="hidden" name="radio-${post.ID}" value="1">
            <input type="radio" id="Dislikes-${post.ID}" class="hidden" name="radio-${post.ID}" value="-1">
            <input type="radio" id="neutral-${post.ID}" class="hidden" name="radio-${post.ID}" value="0" checked>
            <p> Likes : ${post.Likes}</p>
            <p> Dislikes : ${post.Dislikes}</p>
            <div id="comments-${post.ID}" class="comments"></div> 
            <form id="commentForm-${post.ID}" class="comment-form">
              <input type="text" id="comment-${post.ID}" class="comment-input" placeholder="Add a comment">
              <input type="submit" class="comment-submit">
            </form>
  `;
        postContainer.appendChild(postElement);
        // Appeler la fonction pour afficher les commentaires pour chaque publication
        displayPostComments(post.ID);
      });
      const commentForms = document.querySelectorAll('.comment-form');
      commentForms.forEach((form) => {
        form.addEventListener('submit', (event) => {
          event.preventDefault();
          const postID = form.id.split('-')[1];
          const commentInput = document.getElementById(`comment-${postID}`);
          const commentContent = commentInput.value;
          // Call the function to send the comment to the server
          submitComment(postID, commentContent);
          // Reset the comment input field
          commentInput.value = '';
          // Display the comment
          displayComment(postID, commentContent);
        });
      });
      // fetch("/api/comments")
      //     .then(response => response.json())
      //     .then(data => {
      //         const postContainer = document.getElementById('postContainer');
      //         data.forEach(comment => {
      //             const commentElement = document.createElement('div');
      //             commentElement.innerHTML = `<div>${comment.Description}</div>`;
      //             postContainer.appendChild(commentElement);
      //         });
      //     });
      // Increment the current page number
      currentPage++;
      // Check if there are more posts to load
      if (data.length < postsPerPage) {
        // Display the "The End" message
        const endMessage = document.createElement('h2');
        endMessage.textContent = 'The End';
        postContainer.appendChild(endMessage);
        // Remove the scroll listener since there are no more posts
        window.removeEventListener('scroll', scrollListener);
      }
    });
}

function handleRadioClick(button) {
  const postId = button.name.split('-')[1];
  const likesRadio = document.getElementById(`Likes-${postId}`);
  const neutralRadio = document.getElementById(`neutral-${postId}`);

  if (!likesRadio.checked) {
    likesRadio.checked = true;
    sendLike(postId, 1);
  } else {
    neutralRadio.checked = true;
    sendLike(postId, 0);
  }
}

function handleRadioClickDislike(button) {
  const postId = button.name.split('-')[1];
  const dislikesRadio = document.getElementById(`Dislikes-${postId}`);
  const neutralRadio = document.getElementById(`neutral-${postId}`);

  if (!dislikesRadio.checked) {
    dislikesRadio.checked = true;
    sendLike(postId, -1);
  } else {
    neutralRadio.checked = true;
    sendLike(postId, 0);
  }
}

function sendLike(postId, likeValue) {
  fetch(`/api/like`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      postId: postId,
      likeValue: likeValue
    })
  })
  .then(response => {
    if (response.ok) {
      console.log(`Like sent for post ${postId} with value ${likeValue}`);
      // Handle success if needed
    } else {
      console.log(`Failed to send like for post ${postId}`);
      // Handle failure if needed
    }
  })
  .catch(error => {
    console.log(`Error occurred while sending like for post ${postId}:`, error);
    // Handle error if needed
  });
}

function displayComment(postID, commentContent) {
  if (commentContent) {
    const commentContainer = document.createElement("div");
    commentContainer.classList.add("comment-container"); // Ajouter la classe CSS
    const commentText = document.createElement("p");
    commentText.textContent = commentContent;
    commentContainer.appendChild(commentText);
    const postCommentsContainer = document.getElementById(`comments-${postID}`);
    postCommentsContainer.appendChild(commentContainer);
    // Ajouter le commentaire au stockage local
    const commentsKey = `comments_${postID}`;
    const existingComments = localStorage.getItem(commentsKey);
    const updatedComments = existingComments ? JSON.parse(existingComments) : [];
    updatedComments.push(commentContent);
    localStorage.setItem(commentsKey, JSON.stringify(updatedComments));
  }
}
// Function to check if the user has scrolled to the bottom of the page
function submitComment(postID, commentContent) {
  // Create a POST request to save the comment
  fetch('/comment', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      postID: postID,
      commentContent: commentContent
    })
  })
    .then(response => {
      if (!response.ok) {
        throw new Error('Comment submission failed.');
      }
      // If the comment was successfully submitted, display it
      //  displayComment(postID, commentContent);
    })
    .catch(error => {
      console.error('Error:', error);
    });
}

// Event listener for scroll events
function isScrolledToBottom() {
  return window.innerHeight + window.scrollY >= document.body.offsetHeight;
}
//const postId = '60'; // Remplacez par l'ID du message souhaité
fetch(`/api/comments?postId=post.ID`)
  .then(response => response.json())
  .then(comments => {
    // Les commentaires sont disponibles ici
    console.log(comments);
    // Vous pouvez appeler une fonction pour traiter les commentaires
    displayPostComments(comments);
  })
  .catch(error => {
    // Gérer les erreurs de requête
    console.error('Erreur lors de la récupération des commentaires:', error);
  });
function displayPostComments(postId) {
  const commentsContainer = document.getElementById(`comments-${postId}`);

    fetch(`/api/comments?postId=${postId}`)
        .then(response => response.json())
        .then(comments => {
            comments.forEach(comment => {
                const commentElement = document.createElement('div');
                commentElement.innerHTML = `
                    <p>${comment.Description}</p>
                    <p>Likes: ${comment.Likes} Dislikes: ${comment.Dislikes}</p>
                    <input type="button" value="Likes" name="Likes-${comment.CommentID}" onclick="handleRadioClickComment(this)">
                    <input type="button" value="Dislikes" name="Dislikes-${comment.CommentID}" onclick="handleRadioClickDislikeComment(this)">
                    <input type="radio" id="Likes-${comment.CommentID}" class="hidden" name="radio-${comment.CommentID}" value="1">
                    <input type="radio" id="Dislikes-${comment.CommentID}" class="hidden" name="radio-${comment.CommentID}" value="-1">
                    <input type="radio" id="neutral-${comment.CommentID}" class="hidden" name="radio-${comment.CommentID}" value="0" checked>
                `;
                commentsContainer.appendChild(commentElement);
            });
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des commentaires :', error);
        });
}

function handleRadioClickComment(button) {
  const commentID = button.name.split('-')[1];
  const likesRadio = document.getElementById(`Likes-${commentID}`);
  const neutralRadio = document.getElementById(`neutral-${commentID}`);

  if (!likesRadio.checked) {
    likesRadio.checked = true;
    sendLikeComment(commentID, 1);
  } else {
    neutralRadio.checked = true;
    sendLikeComment(commentID, 0);
  }
}

function handleRadioClickDislikeComment(button) {
  const commentID = button.name.split('-')[1];
  const dislikesRadio = document.getElementById(`Dislikes-${commentID}`);
  const neutralRadio = document.getElementById(`neutral-${commentID}`);

  if (!dislikesRadio.checked) {
    dislikesRadio.checked = true;
    sendLikeComment(commentID, -1);
  } else {
    neutralRadio.checked = true;
    sendLikeComment(commentID, 0);
  }
}

function sendLikeComment(postId, likeValue) {
  fetch(`/api/commentlike`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      postId: postId,
      likeValue: likeValue
    })
  })
  .then(response => {
    if (response.ok) {
      console.log(`Like sent for comments ${postId} with value ${likeValue}`);
      // Handle success if needed
    } else {
      console.log(`Failed to send like for post ${postId}`);
      // Handle failure if needed
    }
  })
  .catch(error => {
    console.log(`Error occurred while sending like for post ${postId}:`, error);
    // Handle error if needed
  });
}

// Initial call to fetch posts when the page loads
fetchPosts();
