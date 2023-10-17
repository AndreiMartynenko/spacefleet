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

const addArmamentButton = document.querySelector('#add-armament-button');
const newSpaceshipInputArmamentTitle = document.querySelector('#new-spaceship-input-armament-title');
const newSpaceshipInputArmamentQty = document.querySelector('#new-spaceship-input-armament-qty');
const armamentSelected = document.querySelector('#armament-selected');

const armaments = [];

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

            const formData = new FormData();
            formData.append('name', newSpaceshipInputName.value);
            formData.append('class', newSpaceshipInputClass.value)
            formData.append('crew', newSpaceshipInputCrew.value)
            formData.append('image', newSpaceshipInputImage.value)
            formData.append('value', newSpaceshipInputValue.value)
            formData.append('status', newSpaceshipInputStatus.value)
            formData.append('armaments', JSON.stringify(armaments))


            fetch('/spaceship/', {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                // body: new URLSearchParams({
                //     'name': newSpaceshipInputName.value,
                //     'class': newSpaceshipInputClass.value,
                //     'crew': newSpaceshipInputCrew.value,
                //     'image': newSpaceshipInputImage.value,
                //     'value': newSpaceshipInputValue.value,
                //     'status': newSpaceshipInputStatus.value,
                //     'armament': newSpaceshipInputArmament.value
            })

            body: formData
        })
            .then(resp => resp.json())
            .then(data => console.log(data))
            .catch(err => console.log(err))
    });

addArmamentButton.addEventListener('click', e => {
    const title = newSpaceshipInputArmamentTitle.value;
    const qty = newSpaceshipInputArmamentQty.value;

    if (title == '' || Number.isNaN(qty) || qty < 1) {
        alert('error');
        return;
    }

    const armament = document.createElement('div');
    armament.textContent = `${title}: ${qty}`;
    armament.classList.add('armament-item')
    armamentSelected.appendChild(armament)

    armaments.push({ title, qty })




})