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
          postElement.innerHTML = `<form id="idpost"> 
<input name="ID"  type="hidden" value="${post.ID}">
</form>
        `;
          newDivPost.appendChild(postElement);
        })
      })

  $('#idpost').submit(function(event) {
    event.preventDefault(); // Prevent the default form submission
    //appel d'une ou plusieurs func en js
  });
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



function fetchPosts() {


  fetch(`/api/posts?page=${currentPage}&limit=${postsPerPage}`)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        // Update the UI to display the posts
        const postContainer = document.getElementById('postContainer');
        data.forEach(post => {
          const postElement = document.createElement('div');
          postElement.className = 'div-item';
          postElement.innerHTML = `
          <h3 class="Post-Title">${post.Title}</h3>
          <p class="Post-Desc">${post.Description}</p>
          <img src="${post.ImagePath}" alt="Post Image" class="Post-IMG">
                    <input type="hidden" value="${post.ID}">

        `;
          postContainer.appendChild(postElement);
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
// Event listener for scroll events
function scrollListener() {
  if (isScrolledToBottom()) {
    fetchPosts();
  }
}
// Attach the scroll listener to the window
window.addEventListener('scroll', scrollListener);
// Initial call to fetch posts when the page loads
fetchPosts();