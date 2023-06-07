document.forms['register-form'].onsubmit = function(event) {
    // username
    if (this.username.value.trim() === "") {
        document.querySelector(".username-error").innerHTML = "Veuillez saisir un nom d’utilisateur";
        document.querySelector(".username-error").style.display = "block";
        event.preventDefault();
        return false;
    }
    if (this.username.value.trim().length < 5) {
        document.querySelector(".username-error").innerHTML = "Le nom d'utilisateur doit contenir au moins 5 caractères";
        document.querySelector(".username-error").style.display = "block";
        event.preventDefault();
        return false;
    }

    // email
    if (this.email.value.trim() === "") {
        document.querySelector(".email-error").innerHTML = "Veuillez saisir un e-mail";
        document.querySelector(".email-error").style.display = "block";
        event.preventDefault();
        return false;
    }

    if (!this.email.value.trim().includes(".") || !this.email.value.trim().includes("@") ) {
        document.querySelector(".email-error").innerHTML = "Veuillez saisir un e-mail correct";
        document.querySelector(".email-error").style.display = "block";
        event.preventDefault();
        return false;
    }

    // password
    var password = this.password.value.trim();
    if (password === "") {
        document.querySelector(".password-error").innerHTML = "Veuillez entrer un mot de passe";
        document.querySelector(".password-error").style.display = "block";
        event.preventDefault();
        return false;
    }
    // Check password requirements
    if (!/\d/.test(password)) {
        document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins un chiffre";
        document.querySelector(".password-error").style.display = "block";
        event.preventDefault();
        return false;
    }
    if (!/[A-Z]/.test(password)) {
        document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins une lettre majuscule";
        document.querySelector(".password-error").style.display = "block";
        event.preventDefault();
        return false;
    }
    if (!/[!@#$%^&*]/.test(password)) {
        document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins un caractère spécial";
        document.querySelector(".password-error").style.display = "block";
        event.preventDefault();
        return false;
    }
    // All data is valid, allow form submission
    return true;
};