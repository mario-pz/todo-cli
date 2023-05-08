<script>
    export let userID;

    let tasks = [];

    async function fetchTasks() {
        try {
            const response = await fetch(
                `http://localhost:8080/api/gather_tasks?user_id=${userID}`
            );
            const data = await response.json();
            tasks = data;
        } catch (error) {
            console.error(error);
        }
    }

    $: fetchTasks();

    function toggleTask(task) {
        task.completed = !task.completed;
        tasks = [...tasks];
        updateProgressBar();
    }

    function updateProgressBar() {
        let completedTasks = tasks.filter((task) => task.completed);
        let percentCompleted = (completedTasks.length / tasks.length) * 100;
        let progressBar = document.querySelector(".progress-bar");
        progressBar.style.width = `${percentCompleted}%`;
        progressBar.setAttribute("aria-valuenow", percentCompleted);

        return percentCompleted;
    }
</script>

<div class="bg-danger d-flex align-items-center justify-content-center vh-100">
    <ul class="list-unstyled text-center">
        <h1>Your tasks:</h1>
        {#if tasks.length > 0}
            {#each tasks as task}
                <li class="add-it-here mb-3">
                    <button
                        class="btn btn-danger"
                        style={task.completed
                            ? "text-decoration: line-through;"
                            : ""}
                        on:click={() => toggleTask(task)}
                    >
                        {task.description}
                    </button>
                </li>
            {/each}
        {:else}
            <p>No tasks found.</p>
        {/if}
    </ul>
</div>
