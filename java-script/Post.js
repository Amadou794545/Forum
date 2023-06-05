function post() {
    var message = document.getElementById("message").value;
    var titre = document.getElementById("titre").value;

    var titreContainer = document.getElementById("titre-container");
    var messageContainer = document.getElementById("message-container");

    var newTitre = document.createElement("div");
    var newMessage = document.createElement("div");

    newTitre.textContent = "titre : " + titre;
    newMessage.textContent = "Message : " + message;

    titreContainer.appendChild(newTitre);
    messageContainer.appendChild(newMessage);

    document.getElementById("message").value = "";
    document.getElementById("titre").value = "";
}
