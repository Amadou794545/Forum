function ajouterPost() {
    var description = document.getElementById("description").value;
    var titre = document.getElementById("titre").value;
    var image = document.getElementById("image").files[0]; // Récupérer le fichier d'image sélectionné
    var Container = document.getElementById("container");
    var newDiv = document.createElement("div");
    var newTitre = document.createElement("div");
    var newDescription = document.createElement("div");
    var newImage = document.createElement("img"); // Créer un élément <img>
    newTitre.textContent = "Titre : " + titre;
    newDescription.textContent = "Description : " + description;
    // Ajouter des classes CSS
    newTitre.classList.add("titre-style");
    newDiv.classList.add("div-item");

    var titre = document.getElementById("titre").value;
    var contenu = document.getElementById("description").value;

    var filter = document.getElementById('filter');
    var categorie = filter.value;

    if (categorie === 'all') {
        alert("Veuillez sélectionner un filtre avant de publier le post.");
        return; // Arrête l'exécution de la fonction si aucun filtre n'est sélectionné
    }

    var nouveauPost = {
        titre: titre,
        contenu: contenu,
        categorie: categorie
    };

    posts.push(nouveauPost);

    generatePosts(filter.value);

    // Vider les champs du formulaire
    document.getElementById("description").value = "";
    document.getElementById("titre").value = "";
    document.getElementById("image").value = "";

    var postContent = `
        <div id="post">
            <button id="likeButton">Like</button>
            <span id="likeCount">0</span>
            <button id="dislikeButton">Dislike</button>
            <span id="dislikeCount">0</span>
        </div>
    `;

    newDiv.innerHTML = postContent;
    Container.appendChild(newDiv);
    Container.insertAdjacentHTML("beforeend", "<br>");

    var likeButton = document.querySelector("#likeButton");
    var likeCount = document.querySelector("#likeCount");
    var likeValue = 0;
    
    likeButton.addEventListener("click", function() {
        likeValue++;
        likeCount.textContent = likeValue;
        likeButton.classList.toggle("liked");
        dislikeButton.classList.remove("disliked");
    });
    
    var dislikeButton = document.querySelector("#dislikeButton");
    var dislikeCount = document.querySelector("#dislikeCount");
    var dislikeValue = 0;
    
    dislikeButton.addEventListener("click", function() {
        dislikeValue++;
        dislikeCount.textContent = dislikeValue;
        dislikeButton.classList.toggle("disliked");
        likeButton.classList.remove("liked");
    });
    

    newDiv.appendChild(newTitre);
    newDiv.appendChild(newDescription);
    Container.appendChild(newDiv);
    Container.insertAdjacentHTML("beforeend", "<br>");

    // Redimensionner et afficher l'image
    if (image) {
        var reader = new FileReader();
        reader.onload = function(event) {
            newImage.src = event.target.result; // Définir la source de l'image
            newImage.classList.add("image-style"); // Ajouter une classe CSS pour styliser l'image
            Container.appendChild(newImage);
        };
        reader.readAsDataURL(image); // Lire le fichier d'image en tant qu'URL de données
    }

    // Vider les champs du formulaire
    document.getElementById("description").value = "";
    document.getElementById("titre").value = "";
    document.getElementById("image").value = "";
}

function toggleDiv() {
    var div = document.getElementById("myDiv");
    if (div.style.display === "block") {
        div.style.display = "none";
    } else {
        div.style.display = "block";
        document.getElementById("description").focus(); // Placer le focus sur le champ de saisie "description"
    }
}
