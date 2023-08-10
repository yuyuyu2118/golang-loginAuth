document
  .getElementById("loginForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:60180/login", {
      // adjust the URL as per your server setup
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.success) {
          document.getElementById("result").innerText = "ログイン完了";
        } else {
          document.getElementById("result").innerText = "ログイン失敗";
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  });
