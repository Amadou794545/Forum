function toggleDiv() {
  var div = document.getElementById("myDiv");
  if (div.style.display === "block") {
      div.style.display = "none";
  } else {
      div.style.display = "block";
      document.getElementById("Description"); // Placer le focus sur le champ de saisie "message"
  }
}



function post() {
  var description = document.getElementById("description").value;
  var titre = document.getElementById("titre").value;
  var image = document.getElementById("image").files[0];
  var Container = document.getElementById("postContainer");
  var newDiv = document.createElement("div");
  var newTitre = document.createElement("h1");
  var newDescription = document.createElement("p");
  var newImage = document.createElement("img");
  newTitre.textContent = "Titre : " + titre;
  newDescription.textContent = "Description : " + description;
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
  newDiv.innerHTML = postContent;
  newDiv.appendChild(newTitre);
  newDiv.appendChild(newDescription);

  // Redimensionner et afficher l'image
  if (image) {
    var reader = new FileReader();
    reader.onload = function (event) {
      newImage.src = event.target.result; // Définir la source de l'image
      newImage.classList.add("image-style"); // Ajouter une classe CSS pour styliser l'image
      newDiv.appendChild(newImage); // Ajouter l'élément <img> à la div nouvellement créée
    };
    reader.readAsDataURL(image); // Lire le fichier d'image en tant qu'URL de données
  }

  var firstDiv = Container.firstChild;
  Container.insertBefore(newDiv, firstDiv); // Insérer la nouvelle div avant le premier enfant existant

  Container.insertAdjacentHTML("beforeend", "<br>");

  // Gestionnaire d'événement pour le bouton "Like"
  var likeButton = newDiv.querySelector("#likeButton");
  var likeCount = newDiv.querySelector("#likeCount");
  var likeValue = 0;
  likeButton.addEventListener("click", function () {
    likeValue++;
    likeCount.textContent = likeValue;
  });

  // Gestionnaire d'événement pour le bouton "Dislike"
  var dislikeButton = newDiv.querySelector("#dislikeButton");
  var dislikeCount = newDiv.querySelector("#dislikeCount");
  var dislikeValue = 0;
  dislikeButton.addEventListener("click", function () {
    dislikeValue++;
    dislikeCount.textContent = dislikeValue;
  });

  // Vider les champs du formulaire
  document.getElementById("description").value = "";
  document.getElementById("titre").value = "";
  document.getElementById("image").value = "";
}



// Variables to track the current page and number of posts to load
let currentPage = 1;
const postsPerPage = 25;

// Function to fetch posts from the server
function fetchPosts() {
  fetch(`/api/posts?page=${currentPage}&limit=${postsPerPage}`)
    .then(response => response.json())
    .then(data => {
      // Process the retrieved post data
      console.log(data); // Print the data to the console as an example

      // Update the UI to display the posts
      const postContainer = document.getElementById('postContainer');
      data.forEach(post => {
        const postElement = document.createElement('div');
        postElement.className = 'div-item';
        postElement.innerHTML = `
          <h3 class="Post-Title">${post.Title}</h3>
          <p class="Post-Desc">${post.Description}</p>
          <img src="${post.ImagePath}" alt="Post Image" class="Post-IMG">
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
