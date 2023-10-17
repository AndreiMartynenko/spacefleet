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

const deleteSpacecraft = document.querySelector('#delete-spacecraft');

const checkboxName = document.querySelector('#checkbox-name');
const checkboxClass = document.querySelector('#checkbox-class');
const checkboxStatus = document.querySelector('#checkbox-status');
const showAllSpacecraft = document.querySelector('#show-all-spacecraft');

const editSpacecraft = document.querySelector('#edit-spacecraft');

let currentSpaceShipId = 0;
let selectedSpaceShip;

const armaments = [];

fetch('/spaceship', {
    method: 'GET',
    headers: {
        "Content-Type": "application/json"
    }
})
    .then(resp => resp.json())
    .then(data => {

        console.log(data)

        if (data) {

            data.forEach(element => {
                const option = document.createElement('option');
                option.setAttribute('value', element.id);
                option.textContent = element.name;
                spacecraftsSelect.appendChild(option);
                option.addEventListener('click', e => {
                    // alert(element.id)

                    currentSpaceShipId = element.id;
                    editSpacecraft.style.display = 'block';
                    editSpacecraft.value = `Edit ${element.name}`;

                    // option.addEventListener('click', e => {
                    loadSpaceShip(element.id);
                })

            });

        }
    })
    .catch(err => console.log(err));

const loadSpaceShip = (id) => {

    const _name = checkboxName.checked
    const _class = checkboxClass.checked;
    const _status = checkboxStatus.checked;


    let url = `/spaceship/${id}`

    if (_name || _class || _status) {
        const urlParams = new URLSearchParams({})
        if (_name) {
            urlParams.append("name", _name)
        }
        if (_class) {
            urlParams.append("class", _class)
        }
        if (_status) {
            urlParams.append("status", _status)
        }
        url = url + '?' + urlParams;
    }
    fetch(url, 
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

            selectedSpaceShip = data
            spacecraftsInfo.textContent = JSON.stringify(data)
        })
        .catch(err => console.log(err));
}

addSpacecraft.addEventListener('click', e => {

    if (addSpacecraft.value === 'Submit') {

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
        },
        body: FormData,
    })
    .then(resp => resp.json())
    .then(data => console.log(data))
    .catch(err => console.log(err))

} else if (addSpacecraft.value === 'Update') {
    addSpacecraft.value = 'Submit';
}

let url = `/spaceship/${selectedSpaceShip.id}`
const urlParams = new URLSearchParams({})

if (selectedSpaceShip.name !== newSpaceshipInputName.value) {
    urlParams.append('name', newSpaceshipInputName.value);
}
if (selectedSpaceShip.class !== newSpaceshipInputClass.value) {
    urlParams.append('class', newSpaceshipInputClass.value);
}
if (selectedSpaceShip.crew.toString() !== newSpaceshipInputCrew.value) {
    urlParams.append('crew', newSpaceshipInputCrew.value);
}
if (selectedSpaceShip.image !== newSpaceshipInputImage.value) {
    urlParams.append('image', newSpaceshipInputImage.value);
}
if (selectedSpaceShip.value.toString() !== newSpaceshipInputValue.value) {
    urlParams.append('value', newSpaceshipInputValue.value);
}
if (selectedSpaceShip.status !== newSpaceshipInputStatus.value) {
    urlParams.append('status', newSpaceshipInputStatus.value);
}

urlParams.append('armaments', JSON.stringify(armaments));

url = url + '?' + urlParams;

});