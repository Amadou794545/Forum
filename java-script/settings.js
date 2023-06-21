// script.js

function updateUsernameMessageColor() {
    var usernameMessage = document.querySelector('.usernameMessage');
    
    if (usernameMessage) {
      if (usernameMessage.textContent === 'Entrez un pseudo pour lancer la vérification') {
        usernameMessage.style.color = '#5c4f39';
      } else if (usernameMessage.textContent === 'Pseudo déjà utilisé') {
        usernameMessage.style.color = 'red';
      } else if (usernameMessage.textContent === 'Pseudo modifié avec succès') {
        usernameMessage.style.color = 'green';
      }
    }
  }

  window.addEventListener('DOMContentLoaded', updateUsernameMessageColor);
  