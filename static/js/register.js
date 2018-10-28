let regName = document.getElementById("reg-name");
let regEmail = document.getElementById("reg-email");
let regPass1 = document.getElementById("reg-password1");
let regPass2 = document.getElementById("reg-password2");
let info = document.getElementById("info");








// sendRegistration check if form is in valid form, then make an object from user's
// inputed data and then send this data to server with ajax request
function sendRegistration(){
    if ((regPass1.value !== regPass2.value) && regPass1.value !== ""){
        alert("password didn't match")
        return
    }
    var regData = {
        "name": regName.value,
        "email": regEmail.value,
        "password": regPass1.value
    }

    axios.post("/register", regData)
        .then(res => {
            if(res.status === 200){
                location.href = "/"
            }
        })
        .catch(err => {
            switch (err.response.status){
                case 400:
                info.innerHTML = "<p class='danger'>Donnees fourni invalid";
                break;
                case 409:
                info.innerHTML = "<p class='danger'>Compte existe deja merci de vous connecter</p> <a href='/login'>Login</a>";
                break;
                default:
                info.innerText = "<p class='danger'>error interne serveur</p>"
            }
        })
}


