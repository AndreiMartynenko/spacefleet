const spacecraftsSelect = document.querySelector('#spacecrafts');

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
        })

    });
})
.catch(err => console.log(err));