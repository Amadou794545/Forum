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


