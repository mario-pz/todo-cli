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
        <Tasks {userID} />
    {/if}
</main>
