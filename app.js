const spacecraftsSelect = document.querySelector('#spacecrafts');
const spacecraftsInfo = document.querySelector('#spacecrafts-info');

const addSpacecraft = document.querySelector('#add-spacecraft');
const newSpaceshipInputName = document.querySelector('#new-spaceship-input-name');
const newSpaceshipInputClass = document.querySelector('#new-spaceship-input-class');
const newSpaceshipInputCrew = document.querySelector('#new-spaceship-input-crew');
const newSpaceshipInputImage = document.querySelector('#new-spaceship-input-image');
const newSpaceshipInputValue = document.querySelector('#new-spaceship-input-value');
const newSpaceshipInputStatus = document.querySelector('#new-spaceship-input-status');
const newSpaceshipInputArmament = document.querySelector('#new-spaceship-input-armament');

fetch('/spaceship', {
    method: 'GET',
    headers: {
        "Content-Type": "application/json"
    }
})
    .then(resp => resp.json())
    .then(data => {

        console.log(data.payload)
        data.payload.forEach(element => {
            const option = document.createElement('option');
            option.setAttribute('value', element.name);
            option.textContent = element.name;
            spacecraftsSelect.appendChild(option);
            option.addEventListener('click', e => {
                alert(element.id)

                option.addEventListener('click', e => {
                    loadSpaceShip(element.id);
                })

            });
        })
            .catch(err => console.log(err));

        const loadSpaceShip = (id) => {
            fetch(`/spaceship/${id}`,
                {
                    method: 'GET',
                    headers: {
                        "Accept": "application/json"
                    }
                })
                .then(resp => resp.json())
                .then(data => {
                    if (JSON.stringify(data).includes('error:')) {
                        throw new Error(data);
                    }
                    spacecraftsInfo.textContent = JSON.stringify(data)
                })
                .catch(err => console.log(err));
        }

        addSpacecraft.addEventListener('click', e => {
            fetch('/spaceship/', {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: new URLSearchParams({
                    'name': newSpaceshipInputName.value,
                    'class': newSpaceshipInputClass.value,
                    'crew': newSpaceshipInputCrew.value,
                    'image': newSpaceshipInputImage.value,
                    'value': newSpaceshipInputValue.value,
                    'status': newSpaceshipInputStatus.value,
                    'armament': newSpaceshipInputArmament.value
                })
            })
                .then(resp => resp.json())
                .then(data => console.log(data))
                .catch(err => console.log(err))
        });
    }