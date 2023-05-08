<script>
    import Form from "./Form.svelte";
    import Tasks from "./Tasks.svelte";

    let loggedIn = false;

    function updateLoggedIn() {
        loggedIn = true;
    }
    let userID = JSON.parse(localStorage.getItem("user"));
    if (userID) {
        loggedIn = true;
        userID = userID.user_id;
    }
</script>

<main>
    {#if !loggedIn}
        <Form {updateLoggedIn} />
    {:else}
        <nav class="navbar navbar-expand-lg navbar-light bg-success">
            <div class="container-fluid">
                <a class="navbar-brand mx-auto" href="/">Todo Guru</a>
                <a class="navbar-brand mx-auto navbar-right" href="/"
                    >Welcome back</a
                >
            </div>
        </nav>
        <div class="progress">
            <div
                class="progress-bar progress-bar-warning progress-bar-striped"
                role="progressbar"
                aria-valuemin="0"
                aria-valuemax="100"
            />
        </div>

        <Tasks {userID} />
    {/if}
</main>
