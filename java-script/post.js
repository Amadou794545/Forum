function toggleDiv() {
  var div = document.getElementById("myDiv");
  if (div.style.display === "block") {
    div.style.display = "none";
  } else {
    div.style.display = "block";
    document.getElementById("Description") // Placer le focus sur le champ de saisie "message"
  }
}
function post() {
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
  // Création de la div newDivComm
  var newDivComm = document.createElement("div");
  newDivComm.classList.add("div-Comm");
  newDivComm.setAttribute("id","div-Comm")

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



  let currentPag= 0;
  const postsPerPag =1
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


  $('#ID').submit(function(event) {
    event.preventDefault(); // Prevent the default form submission
    //appel d'une ou plusieurs func en js

    var formData = new FormData();
    // Add form data to FormData object
    var value =$('input').value

    formData.append('ID', value);

    $.ajax(({
      url: '/comment', // Update the URL to your server endpoint
      type: 'POST',
      data: formData,
      processData: false,
      contentType: false,

    }))

  })
//commentaire
  var comContent = `
    <div id="myDivCom" style="display: none;">
        <input type="text" id="Comments" placeholder="Saisissez votre commentaire...">
        <button onclick="displayComment()">Envoyer</button>
    </div>
    
  `;
  newDivComm.innerHTML = comContent;

  var commentButton = document.createElement("button");
  commentButton.setAttribute("id", "commentButton");
  commentButton.textContent = "Commentaire";
  commentButton.addEventListener("click", toggleDivCom);
  newDivComm.appendChild(commentButton);

  $('#com').submit(function(event) {
    event.preventDefault(); // Empêcher la soumission du formulaire par défaut
    // Appeler votre fonction de traitement des commentaires ici
  });



  newDiv.appendChild(newDivPost);
  newDiv.appendChild(newDivComm);

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

function displayComment() {
  var commentInput = document.getElementById("Comments").value;
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/submit", true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
      console.log(xhr.responseText);
    }
  };
  var data = JSON.stringify({ comment: commentInput });
  xhr.send(data);

  var commentInput = document.getElementById("Comments");
  var comment =  commentInput.value;
  if (comment) {
    var commentContainer = document.createElement("div");
    commentContainer.classList.add("comment-container"); // Ajouter la classe CSS
    var commentText = document.createElement("p");
    commentText.textContent = comment;
    commentContainer.appendChild(commentText);
    var postContainer = document.getElementById("div-Comm");
    postContainer.appendChild(commentContainer);
    commentInput.value = "";
  }
}


let currentPage = 1;
const postsPerPage = 25;

$(document).ready(function() {

  var errorContainer = $('#errorContainer'); // Error message container

  $('#filterform').submit(function(event) {
    event.preventDefault(); // Prevent the default form submission
    clearallpost();
    currentPage = 1;
    fetchPosts();
    window.addEventListener('scroll', scrollListener);
})});

function clearallpost() {
  const postContainer = document.getElementById('postContainer');
  postContainer.innerHTML = "";
}



// Attach the scroll listener to the window


// Function to fetch posts from the server
function fetchPosts() {
  // Send a GET request to the server to retrieve the posts
  var filters = [];
  $('input[type="checkbox"]:checked').each(function() {
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
          <div id="comments-${post.ID}" class="comments"></div> // Add comment section
          <form id="commentForm-${post.ID}" class="comment-form">
            <input type="text" id="comment-${post.ID}" class="comment-input" placeholder="Add a comment">
            <input type="submit" class="comment-submit">
          </form>
        `;
        postContainer.appendChild(postElement);
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
  });
});

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
    })
    .catch(error => {
      console.error('Error:', error);
    });
}
// Function to check if the user has scrolled to the bottom of the page
function isScrolledToBottom() {
  return window.innerHeight + window.scrollY >= document.body.offsetHeight;
}


function submitComment(postID, commentContent) {
  const url = '/comment';
  const data = {
    postID: postID,
    commentContent: commentContent,
  };

  fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
    .then(response => {
      if (response.ok) {
        console.log('Comment submitted successfully');
      } else {
        console.error('Failed to submit comment');
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });
}

// Event listener for scroll events


// Initial call to fetch posts when the page loads
fetchPosts();
window.addEventListener('scroll', scrollListener);