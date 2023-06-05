function post() {
    var message = document.getElementById("message").value; // Récupérer la valeur du champ de texte
    var messageContainer = document.getElementById("message-container"); // Récupérer l'élément où afficher le message

    var newMessage = document.createElement("p");
    newMessage.textContent = "Message : " + message;

    messageContainer.appendChild(newMessage);

    document.getElementById("message").value = "";
}