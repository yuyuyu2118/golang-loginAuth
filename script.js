document.getElementById('loginForm').addEventListener('submit', function(event) {
  event.preventDefault();

  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  fetch('http://localhost:60180/login', {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          username: username,
          password: password
      })
  })
  .then(response => {
      if (response.ok) {
          alert('Logged in successfully');
      } else {
          alert('Failed to log in');
      }
  })
  .catch(error => console.error('Error:', error));
});
