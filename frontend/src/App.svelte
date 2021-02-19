<script>
    import { onMount } from 'svelte';

    let tasks = [];
    let name = '';
                        
    onMount(async () => {
        const res = await fetch(`http://localhost:8080/alltasks`, {method: `GET`});
        tasks = await res.json();
    });

    const addTask = async () => {
        const res = await fetch (
            `http://localhost:8080/addtask`,
            {
                method: `POST`,
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({name: name, done: false})
            }
        );

        const newTask = await res.json();
        tasks = [...tasks, newTask];
        name = '';
    };

    const remove = async (task) => {
        const res = await fetch(`http://localhost:8080/deletetask/${task.id}`, {method: `DELETE`});
        if (res.ok) {
            tasks = tasks.filter(i => i !== task);
        }
    };

    const toggle = async (task) => {
        const res = await fetch(
            `http://localhost:8080/api/v1/tasks/${task.id}`,
            {
                method: `PATCH`,
                header: {
                    'Content-Type': 'application/json',
                },
                body:JSON.stringify({name: task.name, done: !task.done})
            }
        );

        const modifiedTask = await res.json();
        task = modifiedTask;
        tasks = tasks;
    };
</script>

<main>
    <div>
        <h1>Todo List</h1>

        <form>
            <label for="name">Add an Task</label>
            <input id="name" type="text" bind:value="{name}">
            <button on:click={addTask}>Send it...</button>
        </form>
        <hr>
        <ul>
            {#each tasks as task}
            <li class:done="{task.done}">

                <input type="checkbox" bind:checked={task.done} on:click="{() => toggle(task)}">
                <span>{task.name}</span>
                <button on:click={() => remove(task)}>&times;</button>
            </li>
            {/each}
        </ul>
    </div>
</main>