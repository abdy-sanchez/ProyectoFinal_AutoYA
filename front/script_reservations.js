let homeButton = document.querySelector('[id="home-icon"]');

let tipoParameter = document.querySelector('[class="input-tipo"]');

let marcaParameter = document.querySelector('[class="input-marca"]');

let combustibleParameter = document.querySelector('[class="input-combustible"]');

let transmisionParameter = document.querySelector('[class="input-transmision"]');

let modeloParameter = document.querySelector('[class="input-name-modelo"]');

let SearchButton = document.querySelector('[id="buscar-button"]');



homeButton.addEventListener("click",()=>{

    window.location.assign('home.html');

})


SearchButton.addEventListener("click",()=>{

    let filterURL = `marca=${marcaParameter.value}&tipo=${tipoParameter.value}&combustible=${combustibleParameter.value}&transmision=${transmisionParameter.value}&ModeloNombre=${modeloParameter.value}`;

    ShowFilterResults(filterURL);

})


async function FilterQuery(filterURL){

    let fullURL = 'http://localhost:8080/cars/filter?' + filterURL;

    try {

        const response = await axios.get(fullURL);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

}

async function ShowFilterResults(filterURL){

    [...document.querySelectorAll('[class="car-card"]')].forEach((x)=>x.remove());

    let numeroResults = document.querySelector('[id="numero-results"] span');

    const queryCars = await FilterQuery(filterURL);

    if(queryCars[0] == null){

        Swal.fire({
            title: "Qué mal...",
            text: "No se encontró ningun carro con esos criterios.",
            icon: "error",
            timer: 3200,
            timerProgressBar: true,
            showConfirmButton: false
    
        })
      

    }else{


        let CarCardMainContainer = document.querySelector('[id="results"]');

        for(carElement of queryCars){

            let carCard = document.createElement("div");

            carCard.setAttribute("class","car-card");

            let absBool = "No";

            if(absBool) absBool = "Si";

            carCard.innerHTML = `<div class="car-ref-iamge" style='background-image: url("${carElement.img}");width:190px;height:100px;background-position:center;background-size:cover;'></div>
            <div class="car-marca">Marca: <span class="marca-value">${carElement.marca}</span></div>
            <div class="car-modelo-name">Modelo Nombre: <span class="modelo-name-value">${carElement.modeloNombre}</span></div>
            <div class="car-modelo-year">Modelo Año: <span class="modelo-year-value">${carElement.modeloYear}</span></div>
            <div class="car-transmision">Transmisión: <span class="transmision-value">${carElement.transmision}</span></div>
            <div class="car-combustible">Combustible: <span class="combustible-value">${carElement.combustible}</span></div>
            <div class="car-tipo">Tipo: <span class="tipo-value">${carElement.tipo}</span></div>
            <div class="car-puertas">No. Puertas: <span class="puertas-value">${carElement.puertas}</span></div>
            <div class="car-precio">Precio: <span class="precio-value">${carElement.precio} $</span></div>
            <div class="car-abs"><span class="abs-value">ABS: ${absBool}</span></div>`;

            const reservaButton = document.createElement("button");

            if(carElement.estadoReserva == false){  //se puede reservar [disponible]

                reservaButton.setAttribute("class","reservar-button");

                reservaButton.setAttribute("id",carElement.carroID)

                reservaButton.innerHTML = "Reservar";

            }else{

                reservaButton.setAttribute("class","no-disponible-button");

                reservaButton.innerHTML = "No disponible";

            }


            carCard.appendChild(reservaButton);

            CarCardMainContainer.appendChild(carCard);

        }


        //Funcionalidades Botones Reservar y NO disponible

        let carClickedNoDisponible = document.querySelectorAll('[class="no-disponible-button"]');

        for( noDispo of carClickedNoDisponible){

            noDispo.addEventListener("click",()=>{

                Swal.fire({
                    title: "No Disponible",
                    text: "Este carro ya fue previamente reservado.",
                    icon: "error",
                    timer: 3200,
                    timerProgressBar: true,
                    showConfirmButton: false
            
                })
            })

        }

        let carClickedReservar = document.querySelectorAll('[class="reservar-button"]');

        for( Dispo of carClickedReservar){

            Dispo.addEventListener("click",(clickedButton)=>{

                Swal.fire({
                    title: "Reservar",
                    text: "Deseas realizar esta reserva?",
                    icon: "question",
                    showCancelButton: true,
                    cancelButtonText: "Cancelar",
                }).then((result)=>{

                    if(result.isConfirmed){

                        let carData = clickedButton.target.parentNode;

                        let carName = carData.querySelector('[class="marca-value"]').textContent.trim();

                        let carModel = carData.querySelector('[class="modelo-name-value"]').textContent.trim();

                        let carYear = carData.querySelector('[class="modelo-year-value"]').textContent.trim();

                        let randomReserva = Math.floor(Math.random() * 10000) + 1;

                        reservaOBJCT ={
                            usuario_ID: parseInt(sessionStorage.getItem("usuarioID")),
                            carro_ID: parseInt(clickedButton.target.id),
                            reservaID:randomReserva,
                            resumen:`${carName} ${carModel} (${carYear})`
                        };

                        MakeReserva(reservaOBJCT);
                    }
                })

            })

        }

     
    }

    numeroResults.innerHTML = queryCars.length;

}


async function ReservaQuery(reservaOBJCT){

    try {

        const response = await axios.post('http://localhost:8080/reservations', reservaOBJCT);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

}


async function MakeReserva(reservaOBJCT){

    const reservaData = await ReservaQuery(reservaOBJCT);

    if(reservaData != null){

        Swal.fire("¡Listo!", "Reservación correctamente realizada.", "success").then((result2)=>{

            window.location.assign('home.html');
        })
    }
}
