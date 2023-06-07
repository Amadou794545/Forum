function post() {
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

// Gestionnaire d'événement pour le bouton "Like"
var likeButton = newDiv.querySelector("#likeButton");
var likeCount = newDiv.querySelector("#likeCount");
var likeValue = 0;

likeButton.addEventListener("click", function() {
    likeValue++;
    likeCount.textContent = likeValue;
});

// Gestionnaire d'événement pour le bouton "Dislike"
var dislikeButton = newDiv.querySelector("#dislikeButton");
var dislikeCount = newDiv.querySelector("#dislikeCount");
var dislikeValue = 0;

dislikeButton.addEventListener("click", function() {
    dislikeValue++;
    dislikeCount.textContent = dislikeValue;
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
        document.getElementById("Description").focus(); // Placer le focus sur le champ de saisie "message"
    }
}

$(document).ready(function() {
    // Écouteur d'événement pour les modifications des cases à cocher
    $('#filters input[type="checkbox"]').on('change', function() {
        var selectedFilters = [];

        // Parcourir les cases à cocher et récupérer les filtres sélectionnés
        $('#filters input[type="checkbox"]:checked').each(function() {
            selectedFilters.push($(this).val());
        });

        // Afficher ou masquer les éléments en fonction des filtres sélectionnés
        if (selectedFilters.length > 0) {
            $('#container .post').hide(); // Masquer tous les éléments

            // Afficher les éléments correspondant aux filtres sélectionnés
            for (var i = 0; i < selectedFilters.length; i++) {
                var filter = selectedFilters[i];
                $('#container .post.' + filter).show();
            }
        } else {
            $('#container .post').show(); // Afficher tous les éléments si aucun filtre n'est sélectionné
        }
    });
});
