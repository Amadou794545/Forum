// Variables to track the current page and number of posts to load
let currentPage = 1;
const postsPerPage = 25;

// Function to fetch posts from the server
function fetchPosts() {
  // Send a GET request to the server to retrieve the posts
  fetch(`/api/posts?page=${currentPage}&limit=${postsPerPage}`)
    .then(response => response.json())
    .then(data => {
      // Process the retrieved post data
      console.log(data); // Print the data to the console as an example

      // Update the UI to display the posts
      const postContainer = document.getElementById('postContainer');
      data.forEach(post => {
        const postElement = document.createElement('div');
        postElement.className = 'post';
        postElement.innerHTML = `
          <h3 class="Post-Title">${post.Title}</h3>
          <p class="Post-Desc">${post.Description}</p>
          <img src="${post.ImagePath}" alt="Post Image" class="Post-IMG">
        `;
        postContainer.appendChild(postElement);
      });

      // Increment the current page number
      currentPage++;

      // Check if there are more posts to load
      if (data.length < postsPerPage) {
        // Display the "The End" message
        const endMessage = document.createElement('h2');
        endMessage.textContent = 'The End';
        postContainer.appendChild(endMessage);

        // Remove the scroll listener since there are no more posts
        window.removeEventListener('scroll', scrollListener);
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });
}

// Function to check if the user has scrolled to the bottom of the page
function isScrolledToBottom() {
  return window.innerHeight + window.scrollY >= document.body.offsetHeight;
}

// Event listener for scroll events
function scrollListener() {
  if (isScrolledToBottom()) {
    fetchPosts();
  }
}

// Attach the scroll listener to the window
window.addEventListener('scroll', scrollListener);

// Initial call to fetch posts when the page loads
fetchPosts();
