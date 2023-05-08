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
    }
</script>

<h1>Your tasks:</h1>
{#if tasks.length > 0}
    <ul>
        {#each tasks as task}
            <li class="task">
                <button
                    style={task.completed
                        ? "text-decoration: line-through;"
                        : ""}
                    on:click={() => toggleTask(task)}
                >
                    {task.description}
                </button>
            </li>
        {/each}
    </ul>
{:else}
    <p>No tasks found.</p>
{/if}

<style>
    .task {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin: 8px 0;
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
</style>
