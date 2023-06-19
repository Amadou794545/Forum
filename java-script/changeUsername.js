var usernameMessage = document.querySelector('.username-message');
var isUsernameAvailable = usernameMessage.getAttribute('data-username-available');

if (isUsernameAvailable === "true") {
    usernameMessage.textContent = "Nom d'utilisateur changé avec succès";
    usernameMessage.style.display = "block";
} else {
    usernameMessage.textContent = "Nom d'utilisateur déjà utilisé";
    usernameMessage.style.display = "block";
}
