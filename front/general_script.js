let logOutButton = document.querySelector('[id="log-out"]');

let nombreUsuario = document.querySelector('[id="profile-name"]');

nombreUsuario.innerHTML = `${sessionStorage.getItem("nombreUsuario")} <i class="fa-solid fa-circle-user"></i>`;


logOutButton.addEventListener("click",()=>{

    Swal.fire({
        title: "Estás a punto de salir...",
        icon: "warning",
        showCancelButton: true,
        cancelButtonText: "Cancelar",

    }).then((result)=>{

        if (result.isConfirmed) {

            sessionStorage.clear()

            window.location.assign('index.html');

        }

    })


})

