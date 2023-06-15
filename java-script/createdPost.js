fetch('/api/user/posts?id_user=USER_ID')

    .then(response => response.json())
    .then(posts => {
        const postsContainer = document.getElementById('posts');

        // Loop through the posts and create HTML elements to display them
        posts.forEach(post => {
            const postElement = document.createElement('div');
            postElement.innerHTML = `
                <h2>${post.Title}</h2>
                <p>${post.Description}</p>
                <img src="${post.ImagePath}" alt="Post Image">
            `;

            postsContainer.appendChild(postElement);
        });
    })
    .catch(error => {
        console.error('Error fetching posts:', error);
    });