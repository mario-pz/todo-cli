<script>
    let email = "";
    let password = "";
    let errorMessage = "";

    export let updateLoggedIn;

    async function handleSubmit(event) {
        event.preventDefault();

        try {
            const response = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username: email, password }),
            });

            const user_id = await response.json();

            if (response.ok && Number.isInteger(user_id.user_id)) {
                errorMessage = "";
                localStorage.setItem("user", JSON.stringify(user_id));
                localStorage.setItem(
                    "todo-guru-user",
                    JSON.stringify({ username: email, password })
                );
                updateLoggedIn();
            } else {
                errorMessage = "Login failed. Please try again.";
            }
        } catch (error) {
            errorMessage = error.message;
        }
    }
</script>

<form on:submit={handleSubmit}>
    <label>
        Email:
        <input type="email" bind:value={email} />
    </label>

    <label>
        Password:
        <input type="password" bind:value={password} />
    </label>

    <button type="submit">Submit</button>

    {#if errorMessage}
        <p>{errorMessage}</p>
    {/if}
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 1rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        background-color: #f5f5f5;
    }

    label {
        margin-bottom: 1rem;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        font-size: 1rem;
    }

    input {
        padding: 0.5rem;
        margin-bottom: 1rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 1rem;
        width: 100%;
    }

    button[type="submit"] {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        background-color: #007bff;
        color: #fff;
        font-size: 1rem;
        cursor: pointer;
    }

    button[type="submit"]:hover {
        background-color: #0069d9;
    }
</style>
