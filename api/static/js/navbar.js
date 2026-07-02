const token = localStorage.getItem("token");

const nav = document.getElementById("navButtons");

const card = document.getElementById("shortenerCard");
const result = document.getElementById("result");
const table = document.getElementById("myUrlsSection");
const guest = document.getElementById("guestMessage");

if (token) {

 nav.innerHTML = `

    <span class="welcomeText">

        👋 Welcome

    </span>

    <button class="nav-btn" onclick="logout()">
        Logout
    </button>

`;

    guest.style.display = "none";

    card.style.display = "block";
    result.style.display = "block";
    table.style.display = "block";

} else {

    guest.style.display = "flex";

    card.style.display = "none";
    result.style.display = "none";
    table.style.display = "none";

}