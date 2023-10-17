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

                if (data && Array.isArray(data.addSpacecraft)) {

                spacecraftsSelect.innerHTML = '';

                data.addSpacecraft.forEach(element => {



                const option = document.createElement('option');
                option.setAttribute('value', element.id);
                option.value = element.id;
                option.textContent = element.name;
                spacecraftsSelect.appendChild(option);
                option.addEventListener('click', e => {

                    currentSpaceShipId = element.id;
                    editSpacecraft.style.display = 'block';
                    editSpacecraft.value = `Edit ${element.name}`;

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
            body: formData,
        })
            .then(resp => resp.json())
            .then(data => console.log(data))
            .catch(err => console.log(err))

    } else if (addSpacecraft.value === 'Update') {
        addSpacecraft.value = 'Submit';

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

    fetch(url,
        {
            method: 'PUT',
            headers:
            {
                'Accept': 'application/json'
            }
        })
        .then(resp => resp.json())
        .then(data => console.log(data))
        .catch(err => console.log(err))

}

newSpaceshipInputName.value = '';
newSpaceshipInputClass.value = '';
newSpaceshipInputCrew.value = '';
newSpaceshipInputImage.value = '';
newSpaceshipInputValue.value = '';
newSpaceshipInputStatus.value = '';
newSpaceshipInputArmamentTitle.value = '';
newSpaceshipInputArmamentQty.value = '';
armamentSelected.innerHTML = '';

});


addArmamentButton.addEventListener('click', e => {
    const title = newSpaceshipInputArmamentTitle.value;
    const qty = newSpaceshipInputArmamentQty.value;

    if (title == '' || qty == '') {
        alert('error');
        return;
    }

    const armament = document.createElement('div');
    armament.textContent = `${title}: ${qty}`;
    armament.classList.add('armament-item')
    armamentSelected.appendChild(armament)

    armaments.push({ title, qty })

//DELETE
deleteSpacecraft.addEventListener('click', e => {

    fetch(`/spaceship/${currentSpaceShipId}`, {
        method: 'DELETE',
        headers: {
            'Accept': 'application/json',
        }
    })
        .then(resp => resp.json())
        .then(data => console.log(data))
        .catch(err => console.log(err))
});

showAllSpacecraft.addEventListener('click', e => {

    const _name = checkboxName.checked
    const _class = checkboxClass.checked;
    const _status = checkboxStatus.checked;

    let url = `/spaceship`
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
                throw new Error(resp);
            }
            spacecraftsInfo.textContent = JSON.stringify(resp)
        })
        .catch(err => console.log(err));


//Editing

editSpacecraft.addEventListener('click', e => {

    armamentSelected.textContent = '';
    armaments = [];

    newSpaceshipInputName.value = selectedSpaceShip.name;
    newSpaceshipInputClass.value = selectedSpaceShip.class;
    newSpaceshipInputCrew.value = selectedSpaceShip.crew;
    newSpaceshipInputImage.value = selectedSpaceShip.image;
    newSpaceshipInputValue.value = selectedSpaceShip.value;
    newSpaceshipInputStatus.value = selectedSpaceShip.status;

    if (selectedSpaceShip.armament && selectedSpaceShip.armament.length > 0) {

        selectedSpaceShip.armament.forEach(armament => {
            const el = createArmamentElement(armament.title, armament.qty);
            armamentSelected.appendChild(el);
            armaments.push(armament)
          
        });
    }
    

    //Update

    addSpacecraft.value = "Update";

});

const createArmamentElement = (title, qty) => {
    const armament = document.createElement('div');
    armament.textContent = `${title}: ${qty}`;
    armament.classList.add('armament-item');
    armament.addEventListener('click', e => {
        const content = e.target.textContent;
        const arr = content.split(":");
        const title = arr[0].trim();
        const qty = arr[1].trim();
    });
    return armament;


}