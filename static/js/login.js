let logForm = document.getElementById("login");
let logEmail = document.getElementById("log-email");
let logPassword = document.getElementById("log-password");
let info = document.getElementById("info");




logForm.addEventListener('submit', e => {
    e.preventDefault();
    sendLogin();
});


function sendLogin() {
    if(logEmail.value === "" || logPassword.value === ""){
        info.innerHTML = `<div class="alert alert-danger" role="alert">
                            Invalid input !
                        </div>`
        return;
    }
    var logData = {
        "email": logEmail.value,
        "password": logPassword.value
    }
    console.log(logData)

    axios.post("/login", logData)
        .then(res => {
            console.log(res)
            location.href = "/"
        })
        .catch(err => {
            console.log(err)
            switch (err.response.status){
                case 401:
                info.innerHTML = `<div class="alert alert-danger" role="alert">
                This user do not exist! / ce profil n'existe pas
              </div>`;
                break;
                case 400:
                info.innerHTML = `<div class="alert alert-danger" role="alert">
                Wrong password ! / Mot de passe incorrect
              </div>`;
                break;
                default:
                info.innerHTML = `<div class="alert alert-danger" role="alert">
                Internal server error ! /  Error interne serveur !
              </div>`
            }
        })
}