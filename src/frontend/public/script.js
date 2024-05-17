document.getElementById('registrationForm').addEventListener('submit', async function(event) {
    log("Start");
    event.preventDefault();

    const formData = {
        id: uuid.v4(),
        login: document.getElementById('login').value,
        password: document.getElementById('password').value,
        fullname: document.getElementById('fullname').value,
        email: document.getElementById('email').value,
        phone: document.getElementById('phone').value,
        birthdate: document.getElementById('birthdate').value
    };

    try {
        const response = await fetch('http://localhost:8080/api/v1/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        const result = await response.json();
        console.log(result);
        alert('Registration successful!');
    } catch (error) {
        console.error('Error:', error);
        alert('Registration failed!');
    }
});
