function post() {
    var description = document.getElementById("description").value;
    var titre = document.getElementById("titre").value;

    var descriptionContainer = document.getElementById("container");

    var newDiv = document.createElement("div");
    var newTitre = document.createElement("div");
    var newDescription = document.createElement("div");

    newDiv.classList.add("post-item");
    newTitre.textContent = "Titre : " + titre;
    newDescription.textContent = "Description : " + description;

    newDiv.appendChild(newTitre);
    newDiv.appendChild(newDescription);

    descriptionContainer.appendChild(newDiv);

    document.getElementById("description").value = "";
    document.getElementById("titre").value = "";
}

document.getElementById("imageInput").addEventListener("change", function(event) {
    var file = event.target.files[0];
    if (file) {
        var reader = new FileReader();
        reader.onload = function() {
            var imageUrl = reader.result;
            var imageElement = document.createElement("img");
            imageElement.src = imageUrl;
            document.getElementById("container").appendChild(imageElement);
        };
        reader.readAsDataURL(file);
    }
});

function toggleDiv() {
    var div = document.getElementById("myDiv");
    if (div.style.display === "block") {
        div.style.display = "none";
    } else {
        div.style.display = "block";
        document.getElementById("Description").focus(); // Placer le focus sur le champ de saisie "message"
    }
}