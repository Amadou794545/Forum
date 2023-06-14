function post() {
    var description = document.getElementById("description").value;
    var titre = document.getElementById("titre").value;
    var image = document.getElementById("image").files[0];
    var postContainer = document.getElementById("postContainer");

    // Création de la div principale
    var newDiv = document.createElement("div");
    newDiv.classList.add("main-div");

    // Création de la div newDivPost
    var newDivPost = document.createElement("div");
    newDivPost.classList.add("div-Post");

    // Création de la div newDivComm
    var newDivComm = document.createElement("div");
    newDivComm.classList.add("div-Comm");
    newDivComm.setAttribute("id","div-Comm")

    var newTitre = document.createElement("h1");
    var newDescription = document.createElement("p");
    var newImage = document.createElement("img");

    newTitre.textContent = titre;

    // Ajout des classes CSS
    newTitre.classList.add("titre-style");

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

    // Redimensionner et afficher l'image
    if (image) {
        var reader = new FileReader();
        reader.onload = function (event) {
            newImage.src = event.target.result;
            newImage.classList.add("image-style");
            newDivPost.appendChild(newImage);
        };
        reader.readAsDataURL(image);
    }

    var maxDescriptionLength = 50;
    if (description.split(" ").length > maxDescriptionLength) {
        var shortDescription = description.split(" ", maxDescriptionLength).join(" ");
        var remainingDescription = description.split(" ").slice(maxDescriptionLength).join(" ");

        var descriptionText = document.createElement("span");
        descriptionText.textContent = shortDescription + " ";
        newDescription.appendChild(descriptionText);

        var showMoreButton = document.createElement("a");
        showMoreButton.textContent = "Afficher plus";
        showMoreButton.href = "#";
        showMoreButton.addEventListener("click", function () {
            descriptionText.textContent = description;
            newDescription.appendChild(remainingDescriptionText);
            newDescription.removeChild(showMoreButton);
        });
        newDescription.appendChild(showMoreButton);

        var remainingDescriptionText = document.createElement("span");
        remainingDescriptionText.textContent = " " + remainingDescription;
        remainingDescriptionText.style.display = "none";
    } else {
        newDescription.textContent = description;
    }

    newDivPost.appendChild(newDescription);


    var comContent = `
    <div id="myDivCom" style="display: none;">
        <!-- Contenu de la div des commentaires -->
        <input type="text" id="Description" placeholder="Saisissez votre commentaire...">
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

    var firstDiv = postContainer.firstChild;
    postContainer.insertBefore(newDiv, firstDiv);

    postContainer.insertAdjacentHTML("beforeend", "<br>");

    // Gestionnaire d'événement pour le bouton "Like"
    var likeButton = newDivPost.querySelector("#likeButton");
    var likeCount = newDivPost.querySelector("#likeCount");
    var likeValue = 0;
    likeButton.addEventListener("click", function () {
        likeValue++;
        likeCount.textContent = likeValue;
    });

    // Gestionnaire d'événement pour le bouton "Dislike"
    var dislikeButton = newDivPost.querySelector("#dislikeButton");
    var dislikeCount = newDivPost.querySelector("#dislikeCount");
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
        document.getElementById("Description").focus(); // Placer le focus sur le champ de saisie "message"

        // Afficher toutes les div des commentaires précédemment créées
        for (var i = 0; i < commentDivs.length; i++) {
            commentDivs[i].style.display = "block";
        }
    }
}

function displayComment() {
    var commentInput = document.getElementById("Description");
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


function toggleDiv() {
    var div = document.getElementById("myDiv");
    if (div.style.display === "block") {
        div.style.display = "none";
    } else {
        div.style.display = "block";
        document.getElementById("Description"); // Placer le focus sur le champ de saisie "message"
    }
}

