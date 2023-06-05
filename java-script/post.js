function post() {
    var description = document.getElementById("description").value;
    var titre = document.getElementById("titre").value;

    var Container = document.getElementById("container");

    var newDiv = document.createElement("div");
    var newTitre = document.createElement("div");
    var newDescription = document.createElement("div");

    newTitre.textContent = "Titre : " + titre;
    newDescription.textContent = "Description : " + description;

    // Ajouter une classe CSS spécifique au titre
    newTitre.classList.add("titre-style");
    newDiv.classList.add("div-item");

    newDiv.appendChild(newTitre);
    newDiv.appendChild(newDescription);
    Container.appendChild(newDiv);
    Container.insertAdjacentHTML("beforeend", "<br>");


    document.getElementById("description").value = "";
    document.getElementById("titre").value = "";
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